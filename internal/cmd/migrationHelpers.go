package cmd

import (
	"database/sql"
	"log"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database"
	"github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate/source"
	migSrcDriver "github.com/mattes/migrate/source/file"

	"github.com/pkg/errors"
)

func mustHandleMigrations(db *sql.DB, shouldMigrate bool) {
	fd := &migSrcDriver.File{}
	srcDriver, err := fd.Open("file://internal/data/config/migrations/")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open migration source driver"))
	}
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to initialize migration driver"))
	}
	if shouldMigrate {
		version, err := migrateSchema(srcDriver, dbDriver)
		if err != nil {
			log.Fatalln(errors.Wrap(err, "failed to migrate database"))
		}
		log.Println("database now on the latest version ", version)
	} else {
		version, err := checkSchemaCompatibility(srcDriver, dbDriver)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "database version (%d) is incorrect; migration pending", version))
		}
		log.Println("database is on the latest version ", version)
	}
}

func checkSchemaCompatibility(srcDriver source.Driver, dbDriver database.Driver) (int, error) {
	version, _, _ := dbDriver.Version()
	if version == -1 {
		return version, errors.Errorf("no migrations have been run before")
	}
	next, err := srcDriver.Next(uint(version))
	if err == nil {
		return version, errors.Errorf("a later migration (%d) was found", next)
	}
	return version, nil
}

func migrateSchema(srcDriver source.Driver, dbDriver database.Driver) (int, error) {

	version, dirty, err := dbDriver.Version()
	if dirty {
		log.Println("previous attempt to execute migration ", version, " failed; will clear dirty flag and try again")
		prev, err := srcDriver.Prev(uint(version))
		migTo := int(prev)
		if err != nil {
			migTo = -1
		}
		if err := dbDriver.SetVersion(migTo, false); err != nil {
			return 0, errors.Wrapf(err, "failed to force migration table to most recent successful version (%d)", version)
		}
		version = int(prev)
	}

	m, err := migrate.NewWithInstance("file", srcDriver, "postgres", dbDriver)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create migration")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return 0, errors.Wrap(err, "failed to run migration")
	}
	version, _, err = dbDriver.Version()
	if err != nil {
		return 0, errors.Wrap(err, "failed to verify the database version after an apparently successful migration")
	}
	return version, nil
}
