package locate

import (
	"encoding/json"
	"os"
	"time"

	"com.wlq/objects_storage/config"
	"com.wlq/objects_storage/lib/rabbitmq"
	"com.wlq/objects_storage/lib/rs"
	"com.wlq/objects_storage/lib/types"
)

func Locate(name string) (locateInfo map[int]string) {
	q := rabbitmq.New(os.Getenv(config.EnvKeyRabbitMQAddress))
	q.Publish(config.RabbitExchangeKeyDataServer, name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	locateInfo = make(map[int]string)
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}
	return
}

func Exist(name string) bool {
	return len(Locate(name)) >= rs.DATA_SHARDS
}
