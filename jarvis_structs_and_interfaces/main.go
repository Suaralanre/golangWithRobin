package main

import (
	"flag"
	"fmt"
	"jarvishttp/config"
	"jarvishttp/handler"
	"net/http"
)

var configFile = flag.String("config", "", "the full path to the alternative config file")

func main() {
	flag.Parse()

	// read config from file or env vars
	jsonConfig, err := config.ReadConfig(*configFile)
	if err != nil {
		fmt.Println("failed to read config file, using default values. Error: ", err.Error())
		jsonConfig = config.DefaultConfig()
	}

	fmt.Printf("%+v\n", jsonConfig)

	// set up logger

	// establish database connection

	handler := handler.HTTPHandler{
		DB:     nil,
		Logger: nil,
		Config: jsonConfig,
	}

	router := http.NewServeMux()

	router.HandleFunc("/abc", handler.ABC)
}
