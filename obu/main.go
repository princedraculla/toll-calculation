package main

import (
	"fmt"
	"github/princedraculla/toll-calculation/types"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const recieverEndpoint = "ws://localhost:50000/ws"

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
	obuIds := gentOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial(recieverEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIds); i++ {
			lat, long := gentLocation()
			data := types.OBUData{
				ObuID: obuIds[i],
				Lat:   lat,
				Long:  long,
			}

			if err := conn.WriteJSON(data); err != nil {
				fmt.Printf("error while generating obu data : %s", err)
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
