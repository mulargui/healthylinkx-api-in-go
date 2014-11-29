package main

import (
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/", http.NotFound)
	http.HandleFunc("/taxonomy", taxonomy)
	http.HandleFunc("/providers", providers)
	http.HandleFunc("/shortlist", shortlist)
	http.HandleFunc("/transaction", transaction)
		
	log.Fatal(http.ListenAndServe(":8081", nil))
}
