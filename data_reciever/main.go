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
}

func NewWsHandler() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func main() {
	wsHandler := NewWsHandler()
	http.HandleFunc("/ws", wsHandler.WsHandler)
	http.ListenAndServe(":50000", nil)

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
		fmt.Printf("data recieving, contains OBU ID [%d], :: <lat %.2f, long %.2f> its recived \n", data.ObuID, data.Lat, data.Long)
		dr.msgch <- data
	}
}
