package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)
type Credits struct {
    DateFrom string `json:date_from`
    DateTo string `json:date_to`
    CreditLimits CreditLimits `json:credit_limits`
}

type CreditLimits struct {
    M1 uint `json:"mm"`
    M2 uint `json:"mm"`
}


func postCredits(w http.ResponseWriter, r *http.Request) {
    var credits Credits

    err := json.NewDecoder(r.Body).Decode(&credits)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
    fmt.Fprintf(w, "Credits: %+v", credits)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/statistics", postCredits)

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}