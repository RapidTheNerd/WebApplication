package main

import (
	"net/http"
	"io/ioutil"
	"html/template"
)

func main(){
	p := &Page{title: "test", body: []byte("a test page...idk")}
	p.save()
	http.HandleFunc("/test/", view)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/save/", save)
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
	f := title+ ".txt"
	body, err := ioutil.ReadFile(f)
	if err != nil{
		return nil, err
	}
	return &Page{title, body}, nil

}

func edit(w http.ResponseWriter, r* http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, _ := load(title)
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func save(w http.ResponseWriter, r* http.Request){
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{title, []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func view(w http.ResponseWriter, r * http.Request){
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	t, _ := template.ParseFiles("testpage.html")
	t.Execute(w, p)
}
