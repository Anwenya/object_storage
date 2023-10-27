

# 使用别名接口在同一个接口上绑定多个地址
bind_ip:
	ifconfig ens33:1 192.168.2.129/24
	ifconfig ens33:2 192.168.2.130/24
	ifconfig ens33:3 192.168.2.131/24
	ifconfig ens33:4 192.168.2.132/24
	ifconfig ens33:5 192.168.2.133/24
	ifconfig ens33:6 192.168.2.134/24
	ifconfig ens33:7 192.168.2.135/24
	ifconfig ens33:8 192.168.2.136/24

storage_dir:
	for i in `seq 1 6`; do mkdir -p /home/wlq/tmp/$i/objects; done

server_chapter_1:
	LISTEN_ADDRESS=:7777 STORAGE_ROOT=/home/wlq/tmp go run chapter1/main.go

server_chapter_2:
	LISTEN_ADDRESS=192.168.2.129:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 STORAGE_ROOT=/home/wlq/tmp/1 go run chapter2/data_server/main.go &
	LISTEN_ADDRESS=192.168.2.130:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 STORAGE_ROOT=/home/wlq/tmp/2 go run chapter2/data_server/main.go &
	LISTEN_ADDRESS=192.168.2.131:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 STORAGE_ROOT=/home/wlq/tmp/3 go run chapter2/data_server/main.go &
	LISTEN_ADDRESS=192.168.2.132:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 STORAGE_ROOT=/home/wlq/tmp/4 go run chapter2/data_server/main.go &
	LISTEN_ADDRESS=192.168.2.133:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 STORAGE_ROOT=/home/wlq/tmp/5 go run chapter2/data_server/main.go &
	LISTEN_ADDRESS=192.168.2.134:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 STORAGE_ROOT=/home/wlq/tmp/6 go run chapter2/data_server/main.go &
	LISTEN_ADDRESS=192.168.2.135:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 go run chapter2/api_server/main.go &
	LISTEN_ADDRESS=192.168.2.136:7777 RABBITMQ_SERVER=amqp://wlq:wlq@localhost:5672 go run chapter2/api_server/main.go &