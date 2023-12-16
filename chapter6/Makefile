
create_es_index:
	curl localhost:9200/metadata -XPUT -H "content-type:application/json" -d'{"mappings": {"properties": {"name":{"type": "keyword","index":"false"},"version": {"type":"integer"}, "size":{"type":"integer"}, "hash":{"type":"keyword"}}}}'

server_chapter_6:
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

# 计算生成的文件的hash
# openssl dgst -sha256 -binary /tmp/file | base64
# gMYs2HTexMuy8Xe4+u/3NmEveLZmm7gb4aT2oeVqVQ4=
# gMYs2HTexMuy8Xe4+u%2F3NmEveLZmm7gb4aT2oeVqVQ4=
# gMYs2HTexMuy8Xe4+u%2F3NmEveLZmm7gb4aT2oeVqVQ4=
# gMYs2HTexMuy8Xe4+u%2F3NmEveLZmm7gb4aT2oeVqVQ4=

# 将文件分段上传
# 先生成一个token
# curl -v 192.168.133.135:7777/objects/test6 -XPOST -H "Digest: SHA-256=gMYs2HTexMuy8Xe4+u/3NmEveLZmm7gb4aT2oeVqVQ4=" -H "Size: 100000"

# 根据token查看已上传的大小
# curl -I 192.168.133.135:7777/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJnTVlzMkhUZXhNdXk4WGU0K3UlMkYzTm1FdmVMWm1tN2diNGFUMm9lVnFWUTQ9IiwiU2VydmVycyI6WyIxOTIuMTY4LjEzMy4xMzA6Nzc3NyIsIjE5Mi4xNjguMTMzLjEzMjo3Nzc3IiwiMTkyLjE2OC4xMzMuMTMzOjc3NzciLCIxOTIuMTY4LjEzMy4xMzE6Nzc3NyIsIjE5Mi4xNjguMTMzLjEzNDo3Nzc3IiwiMTkyLjE2OC4xMzMuMTI5Ojc3NzciXSwiVXVpZHMiOlsiMTM1ZmEzZGYtNjNjZS00YjU5LWI1OGUtZmZiZTVjNTJiMTA0IiwiNjZkMmI5YjQtNDhhOC00ZjJiLTg0YTAtYzhiYzQyOTRiZmViIiwiNjdiODI0M2YtNWY4MS00NmYzLWIxMzUtZDFjNTVmOTk0YjAwIiwiMDU2ZWY3YzctYTg5Mi00ZWVmLTg5MjYtMmQ3ZmVjMDQxYzhhIiwiZTFhYjdiOTItNmY0Ny00NmI0LWE1MDQtNDQ0NzA2Y2Y1ZmI2IiwiZTJhYmNkZDAtNTQxOC00NzU4LTk4MzYtNmMxODExYTFjY2EyIl19

# 先上传前50000个字节
# 因为我们接口服务会以32000大小为单位写入数据
# 所以实际写入的只有前32000个字节,剩下的会被丢弃
# curl -v -XPUT --data-binary @/tmp/first 192.168.133.135:7777/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJnTVlzMkhUZXhNdXk4WGU0K3UlMkYzTm1FdmVMWm1tN2diNGFUMm9lVnFWUTQ9IiwiU2VydmVycyI6WyIxOTIuMTY4LjEzMy4xMzA6Nzc3NyIsIjE5Mi4xNjguMTMzLjEzMjo3Nzc3IiwiMTkyLjE2OC4xMzMuMTMzOjc3NzciLCIxOTIuMTY4LjEzMy4xMzE6Nzc3NyIsIjE5Mi4xNjguMTMzLjEzNDo3Nzc3IiwiMTkyLjE2OC4xMzMuMTI5Ojc3NzciXSwiVXVpZHMiOlsiMTM1ZmEzZGYtNjNjZS00YjU5LWI1OGUtZmZiZTVjNTJiMTA0IiwiNjZkMmI5YjQtNDhhOC00ZjJiLTg0YTAtYzhiYzQyOTRiZmViIiwiNjdiODI0M2YtNWY4MS00NmYzLWIxMzUtZDFjNTVmOTk0YjAwIiwiMDU2ZWY3YzctYTg5Mi00ZWVmLTg5MjYtMmQ3ZmVjMDQxYzhhIiwiZTFhYjdiOTItNmY0Ny00NmI0LWE1MDQtNDQ0NzA2Y2Y1ZmI2IiwiZTJhYmNkZDAtNTQxOC00NzU4LTk4MzYtNmMxODExYTFjY2EyIl19

# 查看已上传的大小 应该是32000

# 再从32000开始上传后续的数据
# curl -v -XPUT --data-binary @/tmp/second -H "range:bytes=32000-" 192.168.133.135:7777/temp/eyJOYW1lIjoidGVzdDYiLCJTaXplIjoxMDAwMDAsIkhhc2giOiJnTVlzMkhUZXhNdXk4WGU0K3UlMkYzTm1FdmVMWm1tN2diNGFUMm9lVnFWUTQ9IiwiU2VydmVycyI6WyIxOTIuMTY4LjEzMy4xMzA6Nzc3NyIsIjE5Mi4xNjguMTMzLjEzMjo3Nzc3IiwiMTkyLjE2OC4xMzMuMTMzOjc3NzciLCIxOTIuMTY4LjEzMy4xMzE6Nzc3NyIsIjE5Mi4xNjguMTMzLjEzNDo3Nzc3IiwiMTkyLjE2OC4xMzMuMTI5Ojc3NzciXSwiVXVpZHMiOlsiMTM1ZmEzZGYtNjNjZS00YjU5LWI1OGUtZmZiZTVjNTJiMTA0IiwiNjZkMmI5YjQtNDhhOC00ZjJiLTg0YTAtYzhiYzQyOTRiZmViIiwiNjdiODI0M2YtNWY4MS00NmYzLWIxMzUtZDFjNTVmOTk0YjAwIiwiMDU2ZWY3YzctYTg5Mi00ZWVmLTg5MjYtMmQ3ZmVjMDQxYzhhIiwiZTFhYjdiOTItNmY0Ny00NmI0LWE1MDQtNDQ0NzA2Y2Y1ZmI2IiwiZTJhYmNkZDAtNTQxOC00NzU4LTk4MzYtNmMxODExYTFjY2EyIl19

# GET这个对象对比一下数据
# curl -v 192.168.133.135:7777/objects/test6 > /tmp/output
# diff -s /tmp/output /tmp/file

# 用range头部指定下载test6对象的后68KB数据
# curl 192.168.133.135:7777/objects/test6 -H "range: bytes=32000-" > /tmp/output2

# diff -s /tmp/output2 /tmp/second
