package main

import (
	"log"
	"net/http"
	"os"

	"com.wlq/objects_storage/chapter1/objects"
	"com.wlq/objects_storage/config"
)

func main() {
	http.HandleFunc(config.RouterDataObjects, objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv(config.EnvKeyListenAddress), nil))
}
