package objects

import (
	"fmt"

	"com.wlq/objects_storage/chapter4/api_server/heartbeat"
	"com.wlq/objects_storage/lib/objectstream"
)

func putStream(hash string, size int64) (*objectstream.TempPutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewTempPutStream(server, hash, size)
}
