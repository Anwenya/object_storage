package objects

import (
	"fmt"

	"com.wlq/objects_storage/chapter7/api_server/heartbeat"
	"com.wlq/objects_storage/chapter7/api_server/locate"
	"com.wlq/objects_storage/lib/rs"
)

func GetStream(hash string, size int64) (*rs.RSGetStream, error) {
	locateInfo := locate.Locate(hash)
	// 定位到的数据块小于最小可用数据块视为定位失败
	if len(locateInfo) < rs.DATA_SHARDS {
		return nil, fmt.Errorf("object %s locate fail, result %v", hash, locateInfo)
	}
	dataServers := make([]string, 0)
	// 数据块没有达到最大数据块时进行填补
	if len(locateInfo) != rs.ALL_SHARDS {
		dataServers = heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS-len(locateInfo), locateInfo)
	}
	return rs.NewRSGetStream(locateInfo, dataServers, hash, size)
}
