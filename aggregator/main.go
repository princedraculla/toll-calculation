package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github/princedraculla/toll-calculation/types"
	"net"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	HTTPlistenAddr := flag.String("http", ":5000", "http server")
	GRPClistenAddr := flag.String("grpc", ":5001", "http server")
	flag.Parse()
	store := NewMemoryStore()

	svc := NewInvoiceAggregator(store)

	svc = NewLogMiddleWareAggregator(svc)

	go makeGRPCTransport(*GRPClistenAddr, svc)
	makeHttpTransport(*HTTPlistenAddr, svc)
}

func makeGRPCTransport(listedAddr string, svc Aggregator) error {
	fmt.Println("GRPC server is running at ", listedAddr)

	ln, err := net.Listen("TCP", listedAddr)
	if err != nil {
		return err
	}

	defer ln.Close()

	server := grpc.NewServer([]grpc.ServerOption{}...)
	types.RegisterAggregatorServer(server, NewGRPCAggregatorServer(svc))
	return server.Serve(ln)
}

func makeHttpTransport(listenAdrr string, svc Aggregator) {
	fmt.Println("http server is running at ", listenAdrr)
	http.HandleFunc("/aggregate", HandlerAggregate(svc))
	http.HandleFunc("/invoice", handlerGetInvoice(svc))
	if err := http.ListenAndServe(listenAdrr, nil); err != nil {
		panic(err)
	}

}

func HandlerAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error(), "where": "in aggregate distance call"})
			return
		}
	}
}

func handlerGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.URL.Query()["obu"]
		if !ok {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "obu ID not provided"})
			return
		}
		obuId, err := strconv.Atoi(r.URL.Query()["obu"][0])

		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		invoice, err := svc.CalculateInvoice(obuId)
		fmt.Println(invoice)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error in finding invoice": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, invoice)

	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
