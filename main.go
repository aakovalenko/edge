package main

import (
    "log"
    "net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "200"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/statistics", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodGet)

    log.Fatal(http.ListenAndServe(":8080", r))
}