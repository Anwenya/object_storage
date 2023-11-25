package objects

import (
	"log"
	"net/url"
	"os"

	"com.wlq/objects_storage/config"
	"com.wlq/objects_storage/lib/utils"

	"com.wlq/objects_storage/chapter4/data_server/locate"
)

func getFile(hash string) string {
	file := os.Getenv(config.EnvKeyStorageRaoot) + config.DirPath + hash
	f, _ := os.Open(file)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	return file
}
