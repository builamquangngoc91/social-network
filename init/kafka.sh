#! /bin/bash

docker exec -it kafka bash -c "
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 16 --topic event &&
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 16 --topic feed && 
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 16 --topic comment";