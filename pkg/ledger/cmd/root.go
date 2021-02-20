// Package cmd various CLI commands related to the login service
package cmd

import (
	"fmt"
	"github.com/google/logger"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var log *logger.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ledger",
	Short: "Serve the login services needed and offer CLI support for account operations",
	Long:  `The purpose of the login service is to handle packets related to user account login and server selection.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	log = logger.Init("InitLogger", true, false, ioutil.Discard)
	log.Info("root init()")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/ledger.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("ledger")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
		log.Infof("using config file: %v", viper.ConfigFileUsed())
	}

	requiredParams := []string{
		"database.postgres.host",
		"database.postgres.port",
		"database.postgres.db_user",
		"database.postgres.db_password",
	}
	for _, rp := range requiredParams {
		if !viper.IsSet(rp) {
			log.Fatalf("missing required parameter %v", rp)
		}
	}
}
