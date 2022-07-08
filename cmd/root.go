package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gitter",
	Short: "A cli used to manage git",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .gitter.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		path, err := getConfigPath()

		cobra.CheckErr(err)

		viper.AddConfigPath(path)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gitter")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func getConfigPath() (string, error) {
	if cfgFile != "" {
		return cfgFile, nil
	}

	localPath, err := os.Getwd()

	if err != nil {
		return "", err
	}

	// Locally
	if _, err := os.Stat(fmt.Sprintf("%s/.gitter.yaml", localPath)); err == nil {
		return localPath, nil
	}

	homePath, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	// Globally
	if _, err := os.Stat(fmt.Sprintf("%s/.gitter.yaml", homePath)); err == nil {
		return homePath, nil
	}

	return "", errors.New("unable to find config file")
}
