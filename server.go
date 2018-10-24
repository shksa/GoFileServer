package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// AppConfig is config for an app
type AppConfig struct {
	AppHomePage string
	AppBuildDir string
}

// ServerConfig is the config for this file server
type ServerConfig struct {
	ServerPort             string
	ServerHomePageBuildDir string
	Apps                   []AppConfig
}

var serverConfig ServerConfig

func readConfig() {
	viper.SetConfigName("ServerConfig")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // ensures that Viper will read from environment variables as well.

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&serverConfig); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

func init() {
	readConfig()
	spew.Dump(serverConfig)
}

func main() {

	http.HandleFunc("/greet", greet)

	ServerHomePageFSHandler := http.FileServer(http.Dir(serverConfig.ServerHomePageBuildDir))
	http.Handle("/", ServerHomePageFSHandler)

	for _, AppConfig := range serverConfig.Apps {
		AppBuildFSHandler := http.FileServer(http.Dir(AppConfig.AppBuildDir))
		http.Handle(AppConfig.AppHomePage, http.StripPrefix(AppConfig.AppHomePage, AppBuildFSHandler))
	}

	log.Println("Listening...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
