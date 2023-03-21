# Kafka-Go-
Kafka-Go project to parse into ptorobuf and json message


start the project by following command: sudo docker-compose up


now to create a producer: 
kafkacat -P -b localhost:9092 -t test-topic -H "content-type:application/json" -l test.json


now create consumer:

sudo docker exec -it my-sevice sh
./main

