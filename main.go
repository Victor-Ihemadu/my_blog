package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type blog struct {
	Id      string
	Topic   string
	Content string
}
type DataStructure struct {
	DataStructure []blog
}

var Data1 = DataStructure{DataStructure: []blog{}}

var content blog

func main() {
	r := chi.NewRouter()
	register(r)

	err := http.ListenAndServe(":8082", r)
	log.Println("running app on port :8082")
	if err != nil {
		fmt.Println(err)
	}

}

func register(r *chi.Mux) {
	r.Get("/", homePage)
	r.Get("/create", newPost)
	r.Get("/delete/{Id}", deletePost)
	r.Get("/edit", editPost)
}

//func ArticleCtx(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		articleID := chi.URLParam(r, "articleID")
//		article, err := dbGetArticle(articleID)
//		if err != nil {
//			http.Error(w, http.StatusText(404), 404)
//			return
//		}
//		ctx := context.WithValue(r.Context(), "article", article)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
