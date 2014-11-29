package main

import (
	"net/http"
	"fmt"
	//"encoding/json"
)

func taxonomy(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Hello taxonomy!")
}

func providers(response http.ResponseWriter, r *http.Request) {
	//Only GET methods
	if r.Method != "GET" {
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	var params = r.URL.Query()
	
	var gender = params.Get("gender")
	var lastname1 = params.Get("lastname1")
	var lastname2 = params.Get("lastname2")
	var lastname3 = params.Get("lastname3")
	var specialty = params.Get("specialty")
	var distance = params.Get("distance")
	var zipcode = params.Get("zipcode")

 	//check params
 	if len(zipcode)==0 && len(lastname1)==0 && len(specialty)==0 {
		//response.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(response,"not enough parameters")
		return
 	}

	//building the query string
	
	fmt.Fprintf(response, "Hello providers! #%s#%s#%s#%s#%s#%s#%s#", gender, lastname1,lastname2,lastname3,specialty,distance,zipcode)
}

func shortlist(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Hello shortlist!")
}

func transaction(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Hello transaction!")
}
