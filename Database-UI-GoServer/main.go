package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    "io/ioutil"
)

type printStr struct {
    text string
}

var db *sql.DB
var err error

func webServer(w http.ResponseWriter, r *http.Request){
	if(r.Method == http.MethodGet) {
		results, err := db.Query("SELECT * FROM demo")
		if err != nil {
			panic(err.Error())
	    	}
	    	var strings []string
	    	for results.Next(){
			var id int
			var str string
			err = results.Scan(&id, &str)
			if err != nil {
		    		panic(err.Error())
			}
			strings = append(strings, str)
			log.Println(str)
    		}
    		json.NewEncoder(w).Encode(strings)
	} else if(r.Method == http.MethodPost) {
                data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
	    	}
                var str printStr
        	err = json.Unmarshal(data, &str)
	   	if err != nil {
	    		panic(err.Error())
	   	}
	   	insert,err := db.Prepare("INSERT INTO demo(val) VALUES(?)")
	    	if err != nil {
			panic(err.Error())
	    	}
	    	insert.Exec(str.text)
	    	fmt.Println(str.text)
        } else {
	  	fmt.Println("Not supported")
        }
}

func handleRequests() {
    router := mux.NewRouter()
    router.HandleFunc("/welcome", webServer)
    log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
    db,err= sql.Open("mysql","khyati:khyati@tcp(mysqldb:3306)/goServer")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    _,err = db.Exec("CREATE TABLE IF NOT EXISTS demo ( id integer NOT NULL AUTO_INCREMENT, val varchar(500),  PRIMARY KEY (id) )")
    if err != nil {
       panic(err)
    }
    handleRequests()
}
