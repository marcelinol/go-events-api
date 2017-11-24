#!/bin/bash

date +"%T"
echo "{\"email\":\"value1\"}" > post1.txt
echo "{\"email\":\"value2\"}" > post2.txt
echo "{\"email\":\"value3\"}" > post3.txt
echo "{\"email\":\"value4\"}" > post4.txt
echo "{\"email\":\"value5\"}" > post5.txt
echo "{\"email\":\"value6\"}" > post6.txt
echo "{\"email\":\"value7\"}" > post7.txt
echo "{\"email\":\"value8\"}" > post8.txt
echo "{\"email\":\"value9\"}" > post9.txt
echo "{\"email\":\"value10\"}" > post10.txt
ab -k -c 3 -n 100000 -T application/json -p post1.txt http://localhost:8080/event &
ab -k -c 3 -n 100000 -T application/json -p post2.txt http://localhost:8081/event &
ab -k -c 3 -n 100000 -T application/json -p post3.txt http://localhost:8090/event &
ab -k -c 3 -n 100000 -T application/json -p post4.txt http://localhost:8083/event &
ab -k -c 3 -n 100000 -T application/json -p post5.txt http://localhost:8084/event &
ab -k -c 3 -n 100000 -T application/json -p post6.txt http://localhost:8085/event &
ab -k -c 3 -n 100000 -T application/json -p post7.txt http://localhost:8086/event &
ab -k -c 3 -n 100000 -T application/json -p post8.txt http://localhost:8087/event &
ab -k -c 3 -n 100000 -T application/json -p post9.txt http://localhost:8088/event &
ab -k -c 3 -n 100000 -T application/json -p post10.txt http://localhost:8089/event &
sleep 60
rm post1.txt
rm post2.txt
rm post3.txt
rm post4.txt
rm post5.txt
rm post6.txt
rm post7.txt
rm post8.txt
rm post9.txt
rm post10.txt
