package main

import (
	"net/http"
	"bytes"
	"fmt"
	//"encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func taxonomy(response http.ResponseWriter, r *http.Request) {
	//Only GET methods
	if r.Method != "GET" {
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	var user="root"
	var password="awsawsdb"
	var database="healthylinkx"
	con, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err != nil { /*error*/}

	defer con.Close()

	rows, err := con.Query("SELECT * FROM taxonomy")
	if err != nil { /*error*/}
	
	var buffer bytes.Buffer
	var taxonomy string
	for rows.Next() {
		err = rows.Scan(&taxonomy)
		if err == nil {
				buffer.WriteString(taxonomy)
		}
		fmt.Fprintf(response, buffer. String())
	}
		
	fmt.Fprintf(response, "Hello taxonomy!")

/*	
	$sql = mysql_query($query, $this->db);

	if(mysql_num_rows($sql) <= 0)
		$this->response('no taxonomy records',204); //If no records "No Content" status
		
	$result = array();
	while($rlt = mysql_fetch_array($sql,MYSQL_ASSOC))
		$result[] = $rlt;

	// If success everythig is good send header as "OK" and return list of specialities in JSON format
	$this->response($this->json($result), 200);
*/
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
