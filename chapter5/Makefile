
create_es_index:
	curl localhost:9200/metadata -XPUT -H "content-type:application/json" -d'{"mappings": {"properties": {"name":{"type": "keyword","index":"false"},"version": {"type":"integer"}, "size":{"type":"integer"}, "hash":{"type":"keyword"}}}}'

server_chapter_5:
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.129:7777 STORAGE_ROOT=/home/wlq/tmp/1 go run data_server/main.go &
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.130:7777 STORAGE_ROOT=/home/wlq/tmp/2 go run data_server/main.go &
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.131:7777 STORAGE_ROOT=/home/wlq/tmp/3 go run data_server/main.go &
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.132:7777 STORAGE_ROOT=/home/wlq/tmp/4 go run data_server/main.go &
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.133:7777 STORAGE_ROOT=/home/wlq/tmp/5 go run data_server/main.go &
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.134:7777 STORAGE_ROOT=/home/wlq/tmp/6 go run data_server/main.go &
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.135:7777 go run api_server/main.go &
	ES_SERVER=localhost:9200 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 LISTEN_ADDRESS=192.168.133.136:7777 go run api_server/main.go &

bind_ip:
	ifconfig ens33:1 192.168.133.129/24
	ifconfig ens33:2 192.168.133.130/24
	ifconfig ens33:3 192.168.133.131/24
	ifconfig ens33:4 192.168.133.132/24
	ifconfig ens33:5 192.168.133.133/24
	ifconfig ens33:6 192.168.133.134/24
	ifconfig ens33:7 192.168.133.135/24
	ifconfig ens33:8 192.168.133.136/24


# 测试 
# 内容 "this object will be separate to 4+2 shards"
# hash MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=

# 上传一个文件
# curl -v 192.168.133.136:7777/objects/test5 -XPUT -d "this object will be separate to 4+2 shards" -H "Digest: SHA-256=MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8="

# 查看上传结果
# ls /tmp/?/objects

# curl -v 192.168.133.136:7777/locate/MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=
# 									MBMxWHrPMsuOBaVYHkwScZQRyTRMQyiKp2oelpLZza8=

# ps -ef | grep 'go-build' | awk '{print $2}' | xargs kill -9