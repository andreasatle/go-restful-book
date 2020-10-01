package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andreasatle/go-restful-book/romanNumerals/romanNumerals"
)

func root(w http.ResponseWriter, r *http.Request) {
	urlSplit := strings.Split(r.URL.Path, "/")
	if urlSplit[1] == "roman_number" {
		number, err := strconv.Atoi(strings.TrimSpace(urlSplit[2]))
		if err != nil || number < 1 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Page not found"))
			return
		}
		fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
	}
}

func main() {
	http.HandleFunc("/", root)
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
