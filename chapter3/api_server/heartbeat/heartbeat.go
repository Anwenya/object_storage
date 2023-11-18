package heartbeat

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"com.wlq/objects_storage/config"
	"com.wlq/objects_storage/lib/rabbitmq"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv(config.EnvKeyRabbitMQAddress))
	defer q.Close()
	q.Bind(config.RabbitExchangeKeyApiServer)
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

// 每5秒清理一次超过10秒没有发送心跳的数据访问地址
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

// 获得当前数据服务列表
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}

// 随机选择一个数据服务
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
