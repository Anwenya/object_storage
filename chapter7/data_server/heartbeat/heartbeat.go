package heartbeat

import (
	"os"
	"time"

	"com.wlq/objects_storage/config"
	"com.wlq/objects_storage/lib/rabbitmq"
)

// 将自己的监听地址发送到接口服务的交换服务以声明本数据服务正常
func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv(config.EnvKeyRabbitMQAddress))
	defer q.Close()
	for {
		q.Publish(config.RabbitExchangeKeyApiServer, os.Getenv(config.EnvKeyListenAddress))
		time.Sleep(5 * time.Second)
	}
}
