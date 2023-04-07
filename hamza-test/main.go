package main

import (
        "context"
        "encoding/json"
  //      "fmt"
    //    "reflect"
        "log"
  //      "net"
        "flag"
//        "github.com/netsampler/gopacket/protos/gopacket"
//        "github.com/golang/protobuf/proto"
//      "google.golang.org/protobuf" 
      "github.com/segmentio/kafka-go"
)

var (


     kafkaBroker        = flag.String("kafka.brokers", "kafka:9092", "Kafka brokers list separated by commas")
     kafkaInTopic       = flag.String("input-topic", "input-topic", "Kafka topic to consume from")
     kafkaOutTopic      = flag.String("output-topic", "output-topic", "Kafka topic to produce to")
)

type InputMessage struct {
    SrcAddr       string `json:"SrcAddr"`
    DstAddr       string `json:"DstAddr"`
    SrcMac        string `json:"SrcMac"`
    DstMac        string `json:"DstMac"`
    If            int    `json:"If"`
    OutIf         int    `json:"OutIf"`
    SrcPort       int    `json:"SrcPort"`
    DstPort       int    `json:"DstPort"`
    SequenceNum   int    `json:"SequenceNum"`
    TimeFlowStart int64  `json:"TimeFlowStart"`
    TimeFlowEnd   int64  `json:"TimeFlowEnd"`
}



func main() {
          var inputMsg InputMessage


        r := kafka.NewReader(kafka.ReaderConfig{
                Brokers: []string{*kafkaBroker},
                Topic:   *kafkaInTopic,
                GroupID: "group-id",
        })
        r.SetOffset(0)

        w := kafka.NewWriter(kafka.WriterConfig{
                Brokers: []string{*kafkaBroker},
                Topic:   *kafkaOutTopic,
        })

        defer r.Close()
        defer w.Close()

 for {
                msg, err := r.ReadMessage(context.Background())
                if err != nil {
                        log.Printf("Error reading message: %v", err)
                        log.Print(msg.Value)
                        continue
                }

               
                if err := json.Unmarshal(msg.Value, &inputMsg); err != nil {
                        log.Printf("Error unmarshalling message: %v", err)
                      //log.Print(inputMsg)
                      // log.Print(msg.Value)
                        continue
                }
            //   log.Printf("Unmarshaled message: %#v", inputMsg)
                 log.Printf("ScrAddr =   " , inputMsg.SrcAddr)

   }
}
