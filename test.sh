#!/bin/bash
for j in {1..20}
do
  for i in {1..500}
  do
    curl -d "{\"email\":\"lu$i@example.com\"}" -H "Content-Type: application/json" -X POST http://localhost:8080/event
  done
done
