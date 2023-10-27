#!/bin/bash

curl -v 192.168.2.135:7777/objects/test2 -XPUT -d"this is object test2"

curl -v 192.168.2.136:7777/locate/test3
echo
curl -v 192.168.2.136:7777/objects/test2
echo
