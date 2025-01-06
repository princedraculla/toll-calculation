package main

import (
	"fmt"
	"github/princedraculla/toll-calculation/types"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	prod  DataProducer
}

func NewWsHandler() *DataReceiver {
	var (
		p   DataProducer
		err error
	)
	p, err = NewKafkaProducer()
	if err != nil {
		fmt.Printf("intialing error for kafka Producer: %v\n", err)
	}
	p = NewLogMiddleWare(p)
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

func (dr *DataReceiver) Receiver(data *types.OBUData) error {
	return dr.prod.ProduceData(data)
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
		if err := dr.Receiver(&data); err != nil {
			fmt.Println("error while producing data in kafka: ", err)
		}

		//dr.msgch <- data
	}
}
