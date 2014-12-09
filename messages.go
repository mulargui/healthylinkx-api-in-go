package main

import (
	"net/http"
	//"bytes"
	"fmt"
	"encoding/json"
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
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	defer con.Close()

	rows, err := con.Query("SELECT * FROM taxonomy")
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	type Speciality struct {
		Classification string
	}
	var MyList [] Speciality
	var taxonomy Speciality

	for rows.Next() {
		err = rows.Scan(&taxonomy.Classification)
		if err == nil {
			MyList = append(MyList,taxonomy)
		}
	}
		
	b, err := json.Marshal(MyList)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(b)
}

func shortlist(response http.ResponseWriter, r *http.Request) {
	//Only GET methods
	if r.Method != "GET" {
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	var params = r.URL.Query()
	var npi1 = params.Get("NPI1")
	var npi2 = params.Get("NPI2")
	var npi3 = params.Get("NPI3")

 	//check params
 	if len(npi1)==0 {
		response.WriteHeader(http.StatusNoContent)
		return
 	}

	var user="root"
	var password="awsawsdb"
	var database="healthylinkx"
	con, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer con.Close()

	transaction, err := con.Exec("INSERT INTO transactions VALUES (DEFAULT,DEFAULT,?,?,?)", npi1,npi2,npi3)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	transactionid,err := transaction.LastInsertId()
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}	
	
	type provider struct {
		NPI string
		Provider_Full_Name string
		Provider_Full_Street string
		Provider_Full_City string
		Provider_Business_Practice_Location_Address_Telephone_Number string
	}
	type list struct {
		Transaction int64
		Providers [] provider
	}
	var MyList list
	MyList.Transaction=transactionid

	//return detailed data of the selected providers
	var querystring = "SELECT NPI,Provider_Full_Name,Provider_Full_Street, Provider_Full_City, Provider_Business_Practice_Location_Address_Telephone_Number FROM npidata2 WHERE ((NPI = '"+ npi1 +"')"
 	if len(npi2)!=0 {
		querystring += "OR (NPI = '" + npi2 +"')"
	}
 	if len(npi3)!=0 {
		querystring += "OR (NPI = '" + npi3 +"')"
	}
	querystring += ")";
	

	rows, err := con.Query(querystring)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	var item provider
	for rows.Next() {
		err = rows.Scan(&(item.NPI), 
			&(item.Provider_Full_Name), 
			&(item.Provider_Full_Street), 
			&(item.Provider_Full_City), 
			&(item.Provider_Business_Practice_Location_Address_Telephone_Number))
		if err == nil {
			MyList.Providers = append(MyList.Providers,item)
		}
	}
	
	b, err := json.Marshal(MyList)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	//response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(b)
}

func transaction(response http.ResponseWriter, r *http.Request) {
	//Only GET methods
	if r.Method != "GET" {
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	var params = r.URL.Query()
	var id = params.Get("id")

 	//check params
 	if len(id)==0 {
		response.WriteHeader(http.StatusNoContent)
		return
 	}

	var user="root"
	var password="awsawsdb"
	var database="healthylinkx"
	con, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer con.Close()

	//return detailed data of the selected providers
	var querystring = "SELECT * FROM transactions WHERE (id = '"+ id +"')"

	rows, err := con.Query(querystring)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return

/*
	$rlt = mysql_fetch_array($sql,MYSQL_ASSOC);
	
	//get the providers
	$NPI1 = $rlt["NPI1"];
	$NPI2 = $rlt["NPI1"];
	$NPI3 = $rlt["NPI2"];
	
	//get the details of the providers
	$query = "SELECT NPI,Provider_Full_Name,Provider_Full_Street, Provider_Full_City,
		Provider_Business_Practice_Location_Address_Telephone_Number 
 		FROM npidata2 WHERE ((NPI = '$NPI1')";
	if(!empty($NPI2))
		$query .= "OR (NPI = '$NPI2')";
	if(!empty($NPI3))
		$query .= "OR (NPI = '$NPI3')";
	$query .= ")";
	
	$sql = mysql_query($query, $this->db);

	if(mysql_num_rows($sql) <= 0)
		$this->response('no NPI record',204); // If no records "No Content" status
		
	$result = array();
	while($rlt = mysql_fetch_array($sql,MYSQL_ASSOC))
		$result[] = $rlt;

	// If success everythig is good send header as "OK" and return list of providers in JSON format
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
		response.WriteHeader(http.StatusNoContent)
		return
 	}

	//building the query string
	fmt.Fprintf(response, "Hello providers! #%s#%s#%s#%s#%s#%s#%s#", gender, lastname1,lastname2,lastname3,specialty,distance,zipcode)
}
