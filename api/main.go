package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Run API port :5002")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":5002", r))
}