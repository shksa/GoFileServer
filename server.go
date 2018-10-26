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

// ServerConfig defines the config. of the server
type ServerConfig struct {
	ServerPort string
	Apps       []AppConfig
}

var serverConfig ServerConfig

func readConfig() {
	viper.SetDefault("ENV", "dev")
	viper.SetConfigName("ServerConfig")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // ensures that Viper will read from environment variables as well.

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("\n\nUsing config: %s\n", viper.ConfigFileUsed())

	serverEnvironment := viper.GetString("ENV") // cannot use "env" as the environment variable name in the exec command.

	fmt.Printf("\nUsing environment: %q\n", serverEnvironment)

	switch serverEnvironment {
	case "dev", "DEV", "development", "DEVELOPMENT":
		if err := viper.UnmarshalKey("Development", &serverConfig); err != nil {
			log.Fatal(err)
		}

	case "prod", "PROD", "production", "PRODUCTION":
		if err := viper.UnmarshalKey("Production", &serverConfig); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Println("This case should not be matched")
	}
}

func init() {
	readConfig()
	spew.Dump(serverConfig)
}

func main() {

	http.HandleFunc("/greet", greet)

	for _, appConfig := range serverConfig.Apps {
		appFilesHandler := http.FileServer(http.Dir(appConfig.AppBuildDir))
		appFilesHandler = http.StripPrefix(appConfig.AppHomePage, appFilesHandler)
		http.Handle(appConfig.AppHomePage, appFilesHandler)
	}

	TCPNetworkAddress := fmt.Sprintf("localhost:%s", serverConfig.ServerPort)
	log.Println("Listening...")
	err := http.ListenAndServe(TCPNetworkAddress, logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
