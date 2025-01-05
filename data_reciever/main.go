package main

import (
	"encoding/json"
	"fmt"
	"github/princedraculla/toll-calculation/types"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
)

var kafkaTopic = "obudata"

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	prod  *kafka.Producer
}

func NewWsHandler() *DataReceiver {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  p,
	}
}

func main() {

	// Produce messages to topic (asynchronously)

	// Wait for message deliveries before shutting down

	wsHandler := NewWsHandler()
	http.HandleFunc("/ws", wsHandler.WsHandler)
	http.ListenAndServe(":50000", nil)

}

func (dr *DataReceiver) ProduceData(data *types.OBUData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = dr.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)
	return err
}

func (dr *DataReceiver) WsHandler(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error while creating connection", err)
	}
	dr.conn = conn
	go dr.WsLoop()
}

func (dr *DataReceiver) WsLoop() {
	fmt.Println("new ObU client connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("error while receiving Data : ", err)
		}
		if err := dr.ProduceData(&data); err != nil {
			fmt.Println("error while producing data in kafka: ", err)
		}
		fmt.Printf("data recieving, contains OBU ID [%d], :: <lat %.2f, long %.2f> its recived \n", data.ObuID, data.Lat, data.Long)
		//dr.msgch <- data
	}
}
