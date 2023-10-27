package main

import (
	"log"
	"net/http"
	"os"

	"com.wlq/objects_storage/chapter2/api_server/heartbeat"
	"com.wlq/objects_storage/chapter2/api_server/locate"
	"com.wlq/objects_storage/chapter2/api_server/objects"
	"com.wlq/objects_storage/config"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc(config.RouterApiObjects, objects.Handler)
	http.HandleFunc(config.RouterApiLocate, locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv(config.EnvKeyListenAddress), nil))
}
