package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


func goServer(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, `Hello world`)
}

func handleRequests() {
    router := mux.NewRouter()
    router.HandleFunc("/welcome",goServer)
    log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
    handleRequests()
}
