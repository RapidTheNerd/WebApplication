package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){
	http.HandleFunc("/", responseWriter)
	http.ListenAndServe(":8080", nil)
}

func responseWriter(w http.ResponseWriter, req * http.Request){
	w.Write([]byte("General test"))
}

type Page struct {
	title string
	body []byte
}

func(p* Page) save() error {
	f := p.title + ".txt"
	return ioutil.WriteFile(f, p.body, 0600)
}

func load(title string) (*Page, error){
	f:= title+ ".txt"
	body, err := ioutil.ReadFile(f)
	if err != nil{
		return nil, err
	}
	return &Page{title, body}, nil

}
func view(w http.ResponseWriter, r * http.Request){
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>&s</div>", p.title, p.body)
}
