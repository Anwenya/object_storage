package objects

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"com.wlq/objects_storage/lib/es"
	"com.wlq/objects_storage/lib/utils"
)

func put(w http.ResponseWriter, r *http.Request) {
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, e := storeObject(r.Body, url.PathEscape(hash))
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}
	// 存储之后在es中进行元数据的记录
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size := utils.GetSizeFromHeader(r.Header)
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
