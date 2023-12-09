package locate

import (
	"os"
	"strconv"

	"com.wlq/objects_storage/config"
	"com.wlq/objects_storage/lib/rabbitmq"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv(config.EnvKeyRabbitMQAddress))
	defer q.Close()
	q.Bind(config.RabbitExchangeKeyDataServer)
	c := q.Consume()
	// 判断该消息携带的文件名是否在本地存在
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		// 存在此文件则告知接口服务本数据服务的地址
		if Locate(os.Getenv(config.EnvKeyStorageRoot) + config.DirPath + object) {
			q.Send(msg.ReplyTo, os.Getenv(config.EnvKeyListenAddress))
		}
	}

}
