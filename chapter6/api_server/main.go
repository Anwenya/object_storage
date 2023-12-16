package main

import (
	"log"
	"net/http"
	"os"

	"com.wlq/objects_storage/chapter6/api_server/heartbeat"
	"com.wlq/objects_storage/chapter6/api_server/locate"
	"com.wlq/objects_storage/chapter6/api_server/objects"
	"com.wlq/objects_storage/chapter6/api_server/temp"
	"com.wlq/objects_storage/chapter6/api_server/versions"
	"com.wlq/objects_storage/config"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc(config.RouterApiObjects, objects.Handler)
	http.HandleFunc(config.RouterApiLocate, locate.Handler)
	http.HandleFunc(config.RouterApiVersions, versions.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv(config.EnvKeyListenAddress), nil))
}
