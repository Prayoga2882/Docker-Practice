package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	port := os.Getenv("APP_PORT")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+ port, nil)
}

func HelloServer(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Hello guys")
}