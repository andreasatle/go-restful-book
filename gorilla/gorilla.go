package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

// URL: http://localhost:8000/articles?id=1&category=books
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", query["category"][0])
	fmt.Fprintf(w, "ID is: %v\n", query["id"][0])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")
	//url, err := r.Get("articleRoute").URL("category", "books", "id", "123")
	//fmt.Println(url, err)
	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")

	s := &http.Server{
		Handler:      r,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
