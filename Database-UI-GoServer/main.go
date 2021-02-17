package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
)

type printStr struct {
    Text string
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
		fmt.Println(r.Body)
                decoder := json.NewDecoder(r.Body)
   		var str printStr
   		err := decoder.Decode(&str)
                fmt.Println(str)
   		if err != nil {
    			panic(err.Error())
   		}
	   	insert,err := db.Prepare("INSERT INTO demo(val) VALUES(?)")
	    	if err != nil {
			panic(err.Error())
	    	}
	    	insert.Exec(str.Text)
	    	fmt.Println(str.Text)
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
