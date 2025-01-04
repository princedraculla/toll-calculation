package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const recieverEndpoint = "ws://localhost:50000"

type OBUData struct {
	ObuID int     `json:"obuId"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func genCordinate() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func gentLocation() (float64, float64) {
	return genCordinate(), genCordinate()
}

func gentOBUIDs(num int) []int {
	ids := make([]int, num)
	for i := 0; i < num; i++ {
		ids[i] = rand.Intn(99999)
	}
	return ids
}

func main() {
	lat, long := gentLocation()
	obuIds := gentOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial(recieverEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIds); i++ {
			data := &OBUData{
				ObuID: obuIds[i],
				Lat:   lat,
				Long:  long,
			}
			fmt.Println(data)
			conn.WriteJSON(data)
		}
		time.Sleep(time.Second)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
