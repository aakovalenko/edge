package main

import (
    "log"
    "net/http"
    "encoding/json"
    "fmt"
	"github.com/gorilla/mux"
)


// type Credits struct {
//     DateFrom string `json:date_from`
//     DateTo string `json:date_to`
//     CreditLimits CreditLimits `json:credit_limits`
// }

// type CreditLimits struct {
//     M1 uint `json:"mm"`
//     M2 uint `json:"mm"`
// }

type Credits struct {
    DateFrom string
    DateTo string
}

// func get(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)
//     w.Write([]byte(`{"message": "200"}`))
// }

func post(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "application/json")
    // w.WriteHeader(http.StatusCreated)
    // w.Write([]byte(`{"message": "post called"}`))

    var c Credits

    fmt.Println(r.Body)

    err := json.NewDecoder(r.Body).Decode(&c)
    fmt.Println(err)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "Credits: %+v", c)

}

func main() {
    r := mux.NewRouter()
    // r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/statistics", post).Methods(http.MethodPost)

    log.Fatal(http.ListenAndServe(":8080", r))
}