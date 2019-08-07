package main

import (
	"cloud-native-visibility-hub/web"
	"log"
	"net/http"
)

func main() {
	router := web.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
