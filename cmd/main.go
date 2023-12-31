package main

import (
	"log"
	"mama-money-ussd/internal"
	"net/http"
)

func main() {
	http.HandleFunc("/ussd", internal.HandleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
