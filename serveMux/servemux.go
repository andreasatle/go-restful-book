package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct{}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random value is: %f\n", rand.Float64())
}

func randFloat64(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, rand.Float64())
}

func randInt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, rand.Intn(100))
}

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", randFloat64)
	newMux.HandleFunc("/randomInt", randInt)
	http.ListenAndServe(":8000", newMux)
}
