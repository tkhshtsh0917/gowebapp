package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once          sync.Once
	filename      string
	template_data *template.Template
}

func (t *templateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	html := filepath.Join("html", t.filename)
	t.once.Do(func() {
		t.template_data = template.Must(template.ParseFiles(html))
	})

	if err := t.template_data.Execute(res, nil); err != nil {
		log.Fatal("Execute: ", err)
	}
}

func main() {
	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on: ", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
