package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Subdomain string `mapstructure:"subdomain"`
	Token     string `mapstructure:"api_token"`
	Default   struct {
		TaskId int `mapstructure:"task_id"`
	}
}

var cfg Config
var cfgFile string
var debug bool

const subdomainKey = "subdomain"
const apiTokenKey = "api_token"
const defaultTaskIdKey = "default.task_id"

func InitConfig() (*Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}

		// lookup order first look in the current dir then in the home dir
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)

		viper.SetConfigName(".hakuna")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("HAKUNA_CLI")
	viper.BindEnv(subdomainKey)
	viper.BindEnv(apiTokenKey)
	viper.BindEnv(defaultTaskIdKey, "HAKUNA_CLI_DEFAULT_TASK_ID")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && debug {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func EnableDebug() {
	debug = true
}

func BindFlag(name string, flag *pflag.Flag) {
	viper.BindPFlag(name, flag)
}

func BindSubdomainFlag(flag *pflag.Flag) {
	viper.BindPFlag(subdomainKey, flag)
}

func BindTokenFlag(flag *pflag.Flag) {
	viper.BindPFlag(apiTokenKey, flag)
}

func SetConfigFile(path string) {
	cfgFile = path
}
