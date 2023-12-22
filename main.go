package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func init() {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("EDS")
	viper.AutomaticEnv() // read in environment variables that match

	// Find home directory.
	home, _ := homedir.Dir()

	// Search config in home directory with name ".colly-eds" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".colly-eds")
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}

func main() {
	if err := login(); err != nil {
		log.Fatal(err)
	}
	log.Println("登录成功")
}
