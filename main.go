package main

import "net/http"

func main(){
	http.HandleFunc("/", responseWriter)
	http.ListenAndServe(":8080", nil)
}

func responseWriter(w http.ResponseWriter, req * http.Request){
	w.Write([]byte("General test"))
}