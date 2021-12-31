/*
Copyright © 2021 Henning Dahlheim hactar@cyberkraft.ch

*/
package cmd

import (
	"net/http"
	"os"
	"time"

	"github.com/hdahlheim/hakuna-go/internal/config"
	"github.com/hdahlheim/hakuna-go/pkg/hakuna"
	"github.com/spf13/cobra"
)

var debug bool
var cfgFile string
var cliConfig *config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hakuna",
	Short: "A cli to interact with the hakuna time tracking tool.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.SilenceUsage = true
	},
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

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (if not provided the lockup order is ./.hakuna.yaml, $HOME/.hakuna.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "print debug information")
	rootCmd.PersistentFlags().String("subdomain", "", "--subdomain=[hakuna-api-key]")
	rootCmd.PersistentFlags().String("token", "", "--token=[hakuna-api-key]")

	config.BindSubdomainFlag(rootCmd.PersistentFlags().Lookup("subdomain"))
	config.BindTokenFlag(rootCmd.PersistentFlags().Lookup("token"))
}

func initConfig() {
	if cfgFile != "" {
		config.SetConfigFile(cfgFile)
	}

	if debug {
		config.EnableDebug()
	}

	cfg, err := config.InitConfig()
	cobra.CheckErr(err)
	cliConfig = cfg
}

func initHakunaClient() *hakuna.Hakuna {
	client := http.Client{Timeout: 1 * time.Second}
	subdomain := cliConfig.Subdomain
	token := cliConfig.Token

	h, err := hakuna.New(subdomain, token, client)
	cobra.CheckErr(err)

	return h
}
