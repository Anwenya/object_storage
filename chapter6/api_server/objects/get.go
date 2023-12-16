package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"com.wlq/objects_storage/lib/es"
	"com.wlq/objects_storage/lib/utils"
)

func get(w http.ResponseWriter, r *http.Request) {
	// 解析出文件名和版本号
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	versionId := r.URL.Query()["version"]
	version := 0
	var e error
	// 没有指定版本号使用0，默认是最新版本
	if len(versionId) != 0 {
		version, e = strconv.Atoi(versionId[0])
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	// 获得对应版本的元数据
	meta, e := es.GetMetadata(name, version)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 如果hash为空视为不存在
	if meta.Hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 存在则以hash为文件名去数据服务获得该文件
	stream, e := GetStream(meta.Hash, meta.Size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 是否是有range参数
	offset := utils.GetOffsetFromHeader(r.Header)
	if offset != 0 {
		stream.Seek(offset, io.SeekCurrent)
		w.Header().Set("content-range", fmt.Sprintf("bytes %d-%d/%d", offset, meta.Size-1, meta.Size))
		w.WriteHeader(http.StatusPartialContent)
	}

	io.Copy(w, stream)
	stream.Close()
}
