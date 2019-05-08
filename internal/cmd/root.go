package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the `nmc` cmd (executable)
var rootCmd = &cobra.Command{
	Use:   "lazygo",
	Short: "lazygo is an boiler plate code in for api development",
	Long: `
	lazygo is an boiler plate code in for api development
	`,
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read in config file")
	}
	rootCmd.PersistentFlags().StringP(flagEnv, "e", "dev", "The environment to run in (eg. dev, test, staging, prod)")
	viper.BindPFlag(flagEnv, rootCmd.PersistentFlags().Lookup(flagEnv))
	viper.BindEnv(flagEnv, lazyGoEnv)

	rootCmd.PersistentFlags().StringP(flagLevel, "l", "INFO", "The minimum level of logs which are written")
	viper.BindPFlag(flagLevel, rootCmd.PersistentFlags().Lookup(flagLevel))
	viper.BindEnv(flagLevel, lazyGoLogLevel)
}

// Execute is the entry into the CLI, executing the root CMD.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
