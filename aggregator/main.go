package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github/princedraculla/toll-calculation/types"
	"net/http"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5000", "http server")
	flag.Parse()
	store := NewMemoryStore()
	var svc = NewInvoiceAggregator(store)

	makeHttpTransport(*listenAddr, svc)
}

func makeHttpTransport(listenAdrr string, svc Aggregator) {
	fmt.Println("server is running at ", listenAdrr)
	http.HandleFunc("/aggregate", HandlerAggregate(svc))
	if err := http.ListenAndServe(listenAdrr, HandlerAggregate(svc)); err != nil {
		panic(err)
	}

}

func HandlerAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			panic(err)
		}
	}
}
