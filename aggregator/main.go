package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github/princedraculla/toll-calculation/types"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5000", "http server")
	flag.Parse()
	store := NewMemoryStore()

	svc := NewInvoiceAggregator(store)

	svc = NewLogMiddleWareAggregator(svc)

	makeHttpTransport(*listenAddr, svc)
}

func makeHttpTransport(listenAdrr string, svc Aggregator) {
	fmt.Println("server is running at ", listenAdrr)
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
		obuId, err := strconv.Atoi(r.URL.Query()["obu"][0])
		writeJSON(w, http.StatusOK, map[string]any{" value recieved from query": r.URL.Query()["obu"][0]})

		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err := writeJSON(w, http.StatusOK, map[string]int{"obu id in params": obuId}); err != nil {
			logrus.Errorf("error while sending id to front: %v", err)
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
