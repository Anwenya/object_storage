package main

import (
	"log"
	"net/http"
	"os"

	"com.wlq/objects_storage/chapter4/data_server/heartbeat"
	"com.wlq/objects_storage/chapter4/data_server/locate"
	"com.wlq/objects_storage/chapter4/data_server/objects"
	"com.wlq/objects_storage/chapter4/data_server/temp"
	"com.wlq/objects_storage/config"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc(config.RouterDataObjects, objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv(config.EnvKeyListenAddress), nil))
}
