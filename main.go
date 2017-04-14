package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){
	p := &Page{title: "test", body: []byte("a test page...idk")}
	p.save()
	http.HandleFunc("/test/", view)
	http.ListenAndServe(":8080", nil)
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
}
