package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"com.wlq/objects_storage/config"
)

func get(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(os.Getenv(config.EnvKeyStorageRoot) + config.DirPath + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusFound)
		return
	}

	defer file.Close()
	io.Copy(w, file)
}
