package main

import (
	"api/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ConfigEnvironmentVariables()

	fmt.Printf("Run API port :%d \n", config.Port)

	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
