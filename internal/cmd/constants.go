package cmd

const (
	rootConfigFileName = "config"

	// flags
	flagEnv   = "env"
	flagLevel = "level"

	// app env vars
	lazyGoEnv       = "LAZY-GO_ENV"
	lazyGoLogstream = "LAZY-GO_LOG_STREAM"
	lazyGoLogLevel  = "LAZY-GO_LOG_LEVEL"
	dbHost          = "DB_HOST"
	dbUser          = "DB_USER"
	dbPassword      = "DB_PASSWORD"
	dbPort          = "DB_PORT"
	dbName          = "DB_NAME"
	awsRegion       = "AWS_REGION"
	s3Bucket        = "S3_BUCKET"
	lsAwsHost       = "LOCALSTACK_AWS_HOST"
	lsAwsPort       = "LOCALSTACK_AWS_PORT"
	lazyGoPort      = "LAZY-GO_PORT"
	appName         = "APP_NAME"

	// defaults
	defaultDbName     = "lazygo_dev"
	defaultDbHost     = "127.0.0.1"
	defaultDbPort     = "5432"
	defaultDbUser     = "lazygo"
	defaultDbPassword = "lazygo"
	defaultAwsRegion  = "ap-northeast-1"
	defaultS3Bucket   = "lazy_go_s3_bucket"
	defaultLsAwsHost  = "localstack"
	defaultLsAwsPort  = "4572"
	defaultPort       = "3001"
	defaultAppName    = "lazygo"
)
