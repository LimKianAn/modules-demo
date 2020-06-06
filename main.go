package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"rsc.io/quote"
	q2 "rsc.io/quote/v2"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		wr.Write([]byte(quote.Hello() + " (v1) " + q2.HelloV2() + " (v2)"))
	})

	rtr.HandleFunc("/{topic}", func(wr http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		wr.Write([]byte("This route is " + vars["topic"] + "."))
	})

	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
