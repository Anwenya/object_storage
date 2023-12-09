package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"com.wlq/objects_storage/config"
)

func put(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create(os.Getenv(config.EnvKeyStorageRoot) + config.DirPath + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()
	io.Copy(file, r.Body)
}
