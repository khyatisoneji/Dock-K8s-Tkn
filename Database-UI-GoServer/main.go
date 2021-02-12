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

var db *sql.DB
var err error

func webServer(w http.ResponseWriter, r *http.Request){
   results, err := db.Query("SELECT * FROM test")
    if err != nil {
        panic(err.Error())
    }
    var printStr []string
    for results.Next(){
        var id int
        var str string
        err = results.Scan(&id, &str)
        if err != nil {
            panic(err.Error())
        }
        printStr = append(printStr, str)
        log.Println(str)
    }
    json.NewEncoder(w).Encode(printStr)
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
    _,err = db.Exec("CREATE TABLE IF NOT EXISTS test ( id integer, str varchar(500) )")
    if err != nil {
       panic(err)
    }
    printString := []string{
        "Hello World!!",
        "Hope you have a good day!",
    }
    fmt.Println(printString)
    for id, str := range printString {
        insert,err := db.Prepare("INSERT INTO test(id, str) VALUES(?, ?)")
        if err != nil {
            panic(err.Error())
        }
        insert.Exec(id,str)
        log.Println(str)
    }
    handleRequests()
}
