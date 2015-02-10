package main

import (
	"net"
	"net/http"
	"encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	_ "log"
)

/*
type Database struct {
	con *DB
}

func (d Database) Open() error { 
	var user="root"
	var password="awsawsdb"
	var database="healthylinkx"

	con, err := sql.Open("mysql", user + ":" + password + "@/" + database)

	return err
}
*/

func taxonomy(response http.ResponseWriter, r *http.Request) {
	//Only GET methods
	if r.Method != "GET" {
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	//allow cross domain requests
	response.Header().Set("Access-Control-Allow-Origin", "*")

	var user="root"
	var password="awsawsdb"
	var database="healthylinkx"
	
	//supporting docker containers
	var host="MySQLDB"
	hostaddr, err := net.LookupHost(host)
	if err != nil { 
		hostaddr[0]="127.0.0.1"
	}
		
	con, err := sql.Open("mysql", user + ":" + password + "@tcp(" + hostaddr[0] + ":3306)/" + database)
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

	//allow cross domain requests
	response.Header().Set("Access-Control-Allow-Origin", "*")
	
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
	
	//supporting docker containers
	var host="MySQLDB"
	hostaddr, err := net.LookupHost(host)
	if err != nil { 
		hostaddr[0]="127.0.0.1"
	}
		
	con, err := sql.Open("mysql", user + ":" + password + "@tcp(" + hostaddr[0] + ":3306)/" + database)
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
	querystring += ")"
	
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
	
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(b)
}

func transaction(response http.ResponseWriter, r *http.Request) {
	//Only GET methods
	if r.Method != "GET" {
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//allow cross domain requests
	response.Header().Set("Access-Control-Allow-Origin", "*")
	
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
	
	//supporting docker containers
	var host="MySQLDB"
	hostaddr, err := net.LookupHost(host)
	if err != nil { 
		hostaddr[0]="127.0.0.1"
	}
		
	con, err := sql.Open("mysql", user + ":" + password + "@tcp(" + hostaddr[0] + ":3306)/" + database)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	defer con.Close()

	//get list of the selected providers
	var querystring = "SELECT * FROM transactions WHERE (id = '"+ id +"')"

	rows, err := con.Query(querystring)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	if !rows.Next(){
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}

	var empty, empty2, npi1, npi2, npi3 string
	err = rows.Scan(&empty,
			&empty2,
			&npi1, 
			&npi2, 
			&npi3)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	rows.Close()
	
	//get the details of the providers
	querystring = "SELECT NPI,Provider_Full_Name,Provider_Full_Street, Provider_Full_City, Provider_Business_Practice_Location_Address_Telephone_Number FROM npidata2 WHERE ((NPI = '" + npi1 + "')"
 	if len(npi2)!=0 {
		querystring += "OR (NPI = '" + npi2 +"')"
	}
 	if len(npi3)!=0 {
		querystring += "OR (NPI = '" + npi3 +"')"
	}
	querystring += ")"
	
	rows2, err := con.Query(querystring)
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
	var MyList [] provider	
	var item provider
	
	for rows2.Next() {
		err = rows2.Scan(&(item.NPI), 
			&(item.Provider_Full_Name), 
			&(item.Provider_Full_Street), 
			&(item.Provider_Full_City), 
			&(item.Provider_Business_Practice_Location_Address_Telephone_Number))
		if err == nil {
			MyList = append(MyList,item)
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

func providers(response http.ResponseWriter, r *http.Request) {
	//Only GET methods
	if r.Method != "GET" {
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//allow cross domain requests
	response.Header().Set("Access-Control-Allow-Origin", "*")
	
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
	
	var user="root"
	var password="awsawsdb"
	var database="healthylinkx"
	
	//supporting docker containers
	var host="MySQLDB"
	hostaddr, err := net.LookupHost(host)
	if err != nil { 
		hostaddr[0]="127.0.0.1"
	}
		
	con, err := sql.Open("mysql", user + ":" + password + "@tcp(" + hostaddr[0] + ":3306)/" + database)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	
	defer con.Close()	
	
	//building the query string
 	var query = "SELECT NPI,Provider_Full_Name,Provider_Full_Street,Provider_Full_City FROM npidata2 WHERE ("
 	if len(lastname1)!=0{
 		query += "((Provider_Last_Name_Legal_Name = '" + lastname1 + "')"
 	}
	if len(lastname2)!=0{
 		query += " OR (Provider_Last_Name_Legal_Name = '" + lastname2 + "')"
 	}
	if len(lastname3) !=0{
 		query += " OR (Provider_Last_Name_Legal_Name = '" + lastname3 + "')"
 	}
	if len(lastname1) !=0 {
 		query += ")"
	}
 	if len(gender)!=0{
 		if len(lastname1)!=0{
 			query += " AND (Provider_Gender_Code = '" + gender + "')"
 		}else{
 			query += "(Provider_Gender_Code = '" + gender + "')"
		}
	}
 	if len(specialty)!=0{
 		if len(lastname1)!=0 || len(gender)!=0{
 			query += " AND (Classification = '" + specialty + "')"
 		}else{
 			query += "(Classification = '" + specialty + "')"
		}
	}
		
 	//case 1: no need to calculate zip codes at a distance
 	if len(distance)==0 || len(zipcode)==0{
 		if len(zipcode)!=0{
 			if len(lastname1)!=0 || len(gender)!=0 || len(specialty)!=0{
 				query += " AND (Provider_Short_Postal_Code = '"+ zipcode + "')"
 			}else{
 				query += "(Provider_Short_Postal_Code = '" + zipcode + "')"
			}
		}
		query += ") limit 50"
	}else{
		//we need to find zipcodes at a distance
		//lets get a few zipcodes! 	
		var queryapi = "http://www.zipcodeapi.com/rest/GFfN8AXLrdjnQN08Q073p9RK9BSBGcmnRBaZb8KCl40cR1kI1rMrBEbKg4mWgJk7/radius.json/" + zipcode + "/" + distance + "/mile"

		responseapi, err := http.Get(queryapi)
		if err != nil { 
			response.WriteHeader(http.StatusNotAcceptable)
			return
		}
        defer responseapi.Body.Close()
		
        contents, err := ioutil.ReadAll(responseapi.Body)
		if err != nil { 
			response.WriteHeader(http.StatusNotAcceptable)
			return
		}
		
		//unmarshall json data
		type Message struct {
			Zip_code string
			Distance float32
			City string
			State string
		}
		type ListMessage struct{
			Zip_codes [] Message
		}
		var MyMessages ListMessage	

		err = json.Unmarshal(contents, &MyMessages)
		if err != nil { 
			response.WriteHeader(http.StatusNotAcceptable)
			return
		}
	
		//complete the query
 		if len(lastname1)!=0 || len(gender)!=0 || len(specialty)!=0{
 			query += " AND ((Provider_Short_Postal_Code = '" + MyMessages.Zip_codes[0].Zip_code + "')"
 		}else{
 			query += "((Provider_Short_Postal_Code = '" + MyMessages.Zip_codes[0].Zip_code + "')"
		}
		
		var i int		
		for i=1; i<len(MyMessages.Zip_codes);i++ {
 			query += " OR (Provider_Short_Postal_Code = '" + MyMessages.Zip_codes[i].Zip_code + "')"
		}

  		query += ")) limit 50"
	}
	
	//lets make the query and retrieve the data to send back
	type provider struct {
		NPI string
		Provider_Full_Name string
		Provider_Full_Street string
		Provider_Full_City string
	}
	var MyList [] provider

	rows, err := con.Query(query)
	if err != nil { 
		response.WriteHeader(http.StatusNotAcceptable)
		return
	}
	var item provider
	
	for rows.Next() {
		err = rows.Scan(&(item.NPI), 
			&(item.Provider_Full_Name), 
			&(item.Provider_Full_Street), 
			&(item.Provider_Full_City))
		if err == nil {
			MyList = append(MyList,item)
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
