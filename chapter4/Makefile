
create_es_index:
	curl localhost:9200/metadata -XPUT -H "content-type:application/json" -d'{"mappings": {"properties": {"name":{"type": "keyword","index":"false"},"version": {"type":"integer"}, "size":{"type":"integer"}, "hash":{"type":"keyword"}}}}'

server_chapter_4:
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
# 上传多个文件  同内容 不同名
# curl -v 192.168.133.135:7777/objects/test4_1 -XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
# curl -v 192.168.133.135:7777/objects/test4_2 -XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
# curl -v 192.168.133.135:7777/objects/test4_3 -XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
# curl -v 192.168.133.135:7777/objects/test4_4 -XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFG2SOwUQIXY="
# 上传一个hash错误的情况
# curl -v 192.168.133.135:7777/objects/test4_4 -XPUT -d "this object will have only 1 instance" -H "Digest: SHA-256=aWKQ2BipX94sb+h3xdTbWYAu1yzjn5vyFF2SOwUQIXY="
