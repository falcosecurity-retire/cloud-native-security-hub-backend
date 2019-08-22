package main

import (
	"github.com/falcosecurity/cloud-native-security-hub/web"
	"log"
	"net/http"
)

func main() {
	router := web.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
