package temp

import (
	"compress/gzip"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"

	"com.wlq/objects_storage/chapter8/data_server/locate"
	"com.wlq/objects_storage/lib/utils"
)

func (t *tempInfo) hash() string {
	s := strings.Split(t.Name, ".")
	return s[0]
}

func (t *tempInfo) id() int {
	s := strings.Split(t.Name, ".")
	id, _ := strconv.Atoi(s[1])
	return id
}

func commitTempObject(datFile string, tempinfo *tempInfo) {
	f, _ := os.Open(datFile)
	defer f.Close()

	d := url.PathEscape(utils.CalculateHash(f))
	f.Seek(0, io.SeekStart)

	w, _ := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + tempinfo.Name + "." + d)
	w2 := gzip.NewWriter(w)
	io.Copy(w2, f)
	w2.Close()
	os.Remove(datFile)
	locate.Add(tempinfo.hash(), tempinfo.id())
}
