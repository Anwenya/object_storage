package temp

import (
	"net/http"
	"os"
	"strings"

	"com.wlq/objects_storage/config"
)

func del(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Split(r.URL.EscapedPath(), "/")[2]
	infoFile := os.Getenv(config.EnvKeyStorageRoot) + "/temp/" + uuid
	datFile := infoFile + ".dat"
	os.Remove(infoFile)
	os.Remove(datFile)
}
