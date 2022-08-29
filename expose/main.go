package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Hello guys")
}