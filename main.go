package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "The Dark Knight", Director: "Christopher Nolan"},
				{Title: "The Goodfellas", Director: "Martin Scorsese"},
			},
		}

		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		htmlstr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'> %s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlstr)
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
