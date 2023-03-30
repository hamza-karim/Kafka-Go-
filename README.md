# Kafka-Go-
Kafka-Go project to parse into ptorobuf and json message


start the project by following command: sudo docker-compose up


now to create a producer: 
kafkacat -P -b localhost:9092 -t test-topic -H "content-type:application/json" -l test.json


now create consumer:

sudo docker exec -it my-sevice sh
./main

echo '{"SrcAddr":"192.168.1.1","DstAddr":"192.168.1.2","SrcPort":1234,"DstPort":5678,"InIf":1,"OutIf":2,"SrcMac":"01:23:45:67:89:ab","DstMac":"cd:ef:01:23:45:67"}' | \
sudo docker-compose exec -T kafka kafka-console-producer --broker-list kafka:9092 --topic input-topic
