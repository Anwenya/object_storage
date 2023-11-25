package locate

import (
	"os"
	"strconv"
	"time"

	"com.wlq/objects_storage/config"
	"com.wlq/objects_storage/lib/rabbitmq"
)

func Locate(name string) string {
	q := rabbitmq.New(os.Getenv(config.EnvKeyRabbitMQAddress))
	q.Publish(config.RabbitExchangeKeyDataServer, name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
