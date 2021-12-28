/*
Copyright © 2021 Henning Dahlheim hactar@cyberkraft.ch

*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hdahlheim/hakuna-go/pkg/hakuna"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type configData struct {
	Subdomain     string
	APIToken      string
	DefaultTaskId int
}

var config configData

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hakuna",
	Short: "A cli to interact with the hakuna time tracking tool.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.hakuna-cli.yaml)")
	rootCmd.PersistentFlags().String("subdomain", "", "--subdomain=[hakuan-api-key]")
	rootCmd.PersistentFlags().String("token", "", "--token=[hakuan-api-key]")
	viper.BindPFlag("subdomain", rootCmd.Flags().Lookup("subdomain"))
	viper.BindPFlag("api_token", rootCmd.Flags().Lookup("token"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".hakuna-cli")
	}

	viper.SetEnvPrefix("HAKUNA_CLI")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Fprintf(os.Stderr, "unable to decode into struct, %v", err)
	}
}

func getHakunaClient() *hakuna.Hakuna {
	client := http.Client{Timeout: 1 * time.Second}
	subdomain := viper.GetString("subdomain")
	token := viper.GetString("api_token")

	h, err := hakuna.New(subdomain, token, client)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return h
}
