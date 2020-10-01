package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Data struct{ a, b int }

func (d *Data) middleware(handler http.Handler) http.Handler {

	middleFunc := func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Executing middleware before request phase!", d.a)

		// Pass control back to the handler
		handler.ServeHTTP(w, r)

		fmt.Println("Executing middleware after response phase!", d.b)
	}

	return http.HandlerFunc(middleFunc)
}

func (d *Data) mainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainHandler...", d.a, d.b)
	w.Write([]byte("OK" + strconv.Itoa(d.a)))
}

func main() {

	d := &Data{a: 1, b: 2}
	// HandlerFunc returns a HTTP Handler
	mainLogicHandler := http.HandlerFunc(d.mainLogic)

	http.Handle("/", d.middleware(mainLogicHandler))

	http.ListenAndServe(":8000", nil)

}
