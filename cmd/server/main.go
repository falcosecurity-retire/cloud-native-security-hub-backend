package main

import (
	"github.com/falcosecurity/cloud-native-security-hub/web"
	"log"
	"net/http"
	"os"
)

func main() {
	router := web.NewRouterWithLogger(log.New(os.Stderr, "", log.Ltime|log.Ldate|log.LUTC))

	log.Fatal(http.ListenAndServe(":8080", router))
}
