package temp

import (
	"os"

	"com.wlq/objects_storage/chapter4/data_server/locate"
	"com.wlq/objects_storage/config"
)

func commitTempObject(datFile string, tempinfo *tempInfo) {
	os.Rename(datFile, os.Getenv(config.EnvKeyStorageRoot)+"/objects/"+tempinfo.Name)
	locate.Add(tempinfo.Name)
}
