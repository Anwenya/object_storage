package temp

import (
	"net/url"
	"os"
	"strconv"
	"strings"

	"com.wlq/objects_storage/chapter5/data_server/locate"
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
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	os.Rename(datFile, os.Getenv("STORAGE_ROOT")+"/objects/"+tempinfo.Name+"."+d)
	locate.Add(tempinfo.hash(), tempinfo.id())
}
