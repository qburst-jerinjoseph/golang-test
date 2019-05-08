package cmd

import (
	"lazy-go/internal/app"
	"lazy-go/internal/data"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the lazy go server",
	Long: `
lazy go 'serve' starts the lazy-go server.
Ex.
	Parameter         	Default             Flag                Env var
	---------------  	--------------      ------------------  ---------------
	 Environment      	 dev                 --env (-e)          LAZY-GO_ENV
	 Lazy-go log level 	 INFO                --level (-l)        LAZY-GO_LOG_LEVEL
	 Concurrency      	 1                   --concurrency (-c)  LAZY-GO_CONCURRENCY
	 DB host    		 db                              		 DB_HOST
	 DB port    		 5432                            		 DB_PORT
	 DB name    		 lazygo_dev                      		 DB_NAME
	 DB user    		 lazygo                          		 DB_USER
	 DB pwd     		 lazygo                          		 DB_PASSWORD
	 3S Bucket        	 lazy_go_s3_bucket                		 S3_BUCKET
	 AWS region       	 ap-northeast-1                          AWS_REGION
	 Timezone         	 UTC                                     TZ
	 LocalStack Host  	 localstack							 	 LOCALSTACK_AWS_HOST
	 LocalStack Port  	 4572									 LOCALSTACK_AWS_PORT

	 Environments:    dev       test   staging   production
	 Log Levels:	  DEBUG     INFO   WARN      ERROR         FATAL    PANIC
> lazy-go serve
`,
	Run: serveRun,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveRun(*cobra.Command, []string) {
	viper.SetDefault(dbHost, defaultDbHost)
	viper.BindEnv(dbHost)
	viper.SetDefault(dbPort, defaultDbPort)
	viper.BindEnv(dbPort)
	viper.SetDefault(dbName, defaultDbName)
	viper.BindEnv(dbName)
	viper.SetDefault(dbUser, defaultDbUser)
	viper.BindEnv(dbUser)
	viper.SetDefault(dbPassword, defaultDbPassword)
	viper.BindEnv(dbPassword)
	viper.SetDefault(dbHost, defaultDbHost)
	viper.BindEnv(dbHost)
	viper.SetDefault(awsRegion, defaultAwsRegion)
	viper.BindEnv(awsRegion)
	viper.SetDefault(s3Bucket, defaultS3Bucket)
	viper.BindEnv(s3Bucket)
	viper.SetDefault(lsAwsHost, defaultLsAwsHost)
	viper.BindEnv(lsAwsHost)
	viper.SetDefault(lsAwsPort, defaultLsAwsPort)
	viper.BindEnv(lsAwsPort)
	viper.SetDefault(appName, defaultAppName)
	viper.BindEnv(appName)
	viper.SetDefault(lazyGoPort, defaultPort)
	viper.BindEnv(lazyGoPort)
	l := logrus.New()

	dbUser := viper.GetString(dbUser)
	dbPwd := viper.GetString(dbPassword)
	dbHost := viper.GetString(dbHost)
	dbPort := viper.GetString(dbPort)
	dbName := viper.GetString(dbName)
	db := dbInit(dbUser, dbPwd, dbHost, dbPort, dbName)
	s := app.Server{
		Name:   viper.GetString(appName),
		Repo:   data.NewRepo(db),
		Logger: l,
	}
	app.Serve(s, viper.GetString(lazyGoPort), s.Router())

}
