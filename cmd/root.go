/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cdeez/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cdeez",
	Short: "CDEEZ (nuts)\nA super duper blazzzzingly fast gopher powered continuous deployment tool",
}

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "Serves the cdeez api (to be used on prod server)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			utils.Setup(args[0])
		} else {
			utils.Setup("")
		}
	},
}


var appsCmd = &cobra.Command{
	Use: "apps",
	Short: "Display all currently running app names",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Apps())
	},
}


var deployCmd = &cobra.Command{
	Use: "deploy",
	Short: "Deploy a new app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Deploy(args[0], args[1]))
	},
}


var statsCmd = &cobra.Command{
	Use: "stats",
	Short: "Display currently used resources for a speciffic app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Stats(args[0]))
	},
}


var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete an app",
	Run: func(cmd *cobra.Command, args []string) {
		Delete(args[0])
	},
}

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Setup a new app for deployment",
	Run: func(cmd *cobra.Command, args []string) {
		Init(args[0])
	},
}


var restartCmd = &cobra.Command{
	Use: "restart",
	Short: "Update an app",
	Run: func(cmd *cobra.Command, args []string) {
		Restart(args[0])
	},
}

func Check(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func Execute() {
	rootCmd.AddCommand(appsCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(restartCmd)
	err := rootCmd.Execute()
	Check(err)
}

func init() {
	cobra.OnInitialize(initConfig)


	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cdeez.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cdeez" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cdeez")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	}
}
