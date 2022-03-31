package main

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
)

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

type blog struct {
	Id      string
	Topic   string
	Content string
	State   bool
	Edit    string
	Delete  string
}

var DataStructure []blog

var v blog

func main() {
	r := chi.NewRouter()
	register(r)

	err := http.ListenAndServe(":8080", r)
	log.Println("running app on port :8080")
	handleErr(err)

}

func register(router *chi.Mux) {
	router.Get("/", index)
	router.Get("/newpost", getContent)
	router.Get("/delete/{Id}", delete)
	router.Post("/newpost", postContent)
	router.Get("/update/{Id}", update)
	router.Post("/update/{Id}", postUpdate)
}

func index(w http.ResponseWriter, r *http.Request) {
	s, err := template.ParseFiles("template/index.html")
	handleErr(err)

	err = s.Execute(w, DataStructure)
	handleErr(err)

}

func getContent(w http.ResponseWriter, r *http.Request) {
	s, err := template.ParseFiles("template/newpost.html")
	handleErr(err)

	err = s.ExecuteTemplate(w, "newpost.html", nil)
	handleErr(err)
}

func postContent(w http.ResponseWriter, r *http.Request) {
	p := blog{}

	err := r.ParseForm()
	handleErr(err)

	topic := r.PostForm.Get("topic")
	content := r.PostForm.Get("content")

	p.Id = uuid.NewString()
	p.Topic = topic
	p.Content = content
	p.State = true

	DataStructure = append(DataStructure, p)

	http.Redirect(w, r, "/", 302)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")

	for i, _ := range DataStructure {
		if id == DataStructure[i].Id {

			//Deletes the item with such index
			DataStructure = append(DataStructure[:i], DataStructure[i+1:]...)
		}
	}

	http.Redirect(w, r, "/", 302)
}

func update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")

	for i, _ := range DataStructure {
		if id == DataStructure[i].Id {

			v = DataStructure[i]

			s, err := template.ParseFiles("template/editpost.html")
			handleErr(err)

			err = s.Execute(w, DataStructure[i])
			handleErr(err)

		}
	}
}

func postUpdate(w http.ResponseWriter, r *http.Request) {
	//This gets/populates the content of the form
	err := r.ParseForm()
	handleErr(err)

	for i, _ := range DataStructure {
		if DataStructure[i].Id == v.Id {
			DataStructure = append(DataStructure[:i], DataStructure[i+1:]...)
		}
	}

	topic := r.PostForm.Get("topic")
	content := r.PostForm.Get("content")

	v.Topic = topic
	v.Content = content

	DataStructure = append(DataStructure, v)

	http.Redirect(w, r, "/", 302)
}
