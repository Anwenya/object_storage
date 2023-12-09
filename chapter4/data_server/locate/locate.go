package locate

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"com.wlq/objects_storage/config"
	"com.wlq/objects_storage/lib/rabbitmq"
)

var objects = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) bool {
	mutex.Lock()
	_, ok := objects[hash]
	mutex.Unlock()
	return ok
}

func Add(hash string) {
	mutex.Lock()
	objects[hash] = 1
	mutex.Unlock()
}

func Del(hash string) {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv(config.EnvKeyRabbitMQAddress))
	defer q.Close()
	q.Bind(config.RabbitExchangeKeyDataServer)
	c := q.Consume()
	// 判断该消息携带的文件名是否在本地存在
	for msg := range c {
		hash, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		// 存在此文件则告知接口服务本数据服务的地址
		exist := Locate(hash)

		if exist {
			q.Send(msg.ReplyTo, os.Getenv(config.EnvKeyListenAddress))
		}
	}

}

func CollectObjects() {
	files, _ := filepath.Glob(os.Getenv(config.EnvKeyStorageRoot) + config.DirPath + "*")
	for i := range files {
		hash := filepath.Base(files[i])
		objects[hash] = 1
	}
}
