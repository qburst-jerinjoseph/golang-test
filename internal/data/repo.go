package data

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// repo implements Repo.
var _ Repo = (*repo)(nil)

type repo struct {
	*sql.DB
}

//NewRepo returns a new repo
func NewRepo(db *sql.DB) *repo {
	return &repo{db}
}

const getSampleQuery = `
SELECT text FROM sample
`

func (r *repo) GetSample(c context.Context) ([]string, error) {
	var data []string
	rows, err := r.QueryContext(c, getSampleQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		if err != nil {
			return nil, err
		}
		data = append(data, s)
	}
	return data, nil
}
