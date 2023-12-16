package objects

import (
	"fmt"

	"com.wlq/objects_storage/chapter6/api_server/heartbeat"
	"com.wlq/objects_storage/lib/rs"
)

func putStream(hash string, size int64) (*rs.RSPutStream, error) {
	// 找到最大数量的可用的数据服务
	servers := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(servers) != rs.ALL_SHARDS {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return rs.NewRSPutStream(servers, hash, size)
}
