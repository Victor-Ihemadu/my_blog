package main

import (
	"fmt"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	//template, _ := w.
	files, err := template.ParseFiles("template/index.html")
	handleErr(err)
	err2 := files.Execute(w, Data1)
	fmt.Println("execution successful")
	handleErr(err2)
}

func newPost(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("template/newpost.html")
	handleErr(err)

	topic := r.FormValue("title")
	content := r.FormValue("content")
	id := uuid.NewString()
	newBlog := blog{
		Id:      id,
		Topic:   topic,
		Content: content,
	}

	Data1.DataStructure = append(Data1.DataStructure, newBlog)
	err2 := files.Execute(w, nil)
	handleErr(err2)
	http.Redirect(w, r, "/", http.StatusFound)

}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")
	log.Println("container 1 ", Data1.DataStructure)
	for i, value := range Data1.DataStructure {
		//log.Println("i and value", i, value)
		if id == value.Id {
			//log.Println("i and value", id, value.Id)
			Data1.DataStructure = append(Data1.DataStructure[:i], Data1.DataStructure[i+1:]...)
			log.Println("container 2 ", Data1.DataStructure)
		}
	}
	log.Println(Data1.DataStructure)
	http.Redirect(w, r, "/", 302)
}

func editPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")

	for _, value := range Data1.DataStructure {
		if id == value.Id {

			content = value

			files, err := template.ParseFiles("template/editpost.html")
			handleErr(err)

			err = files.Execute(w, value)
			handleErr(err)

		}
	}

}
