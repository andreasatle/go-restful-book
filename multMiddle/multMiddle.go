package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

// City contains the name and area of a city
type City struct {
	Name string
	Area uint64
}

func main() {

	// main business-logic
	mainLogicHandler := http.HandlerFunc(mainLogicPOST)

	// Filter with chain of functions
	chain := alice.New(filterContentType, setServerTimeCookie, filterPOST)

	// Route /city
	http.Handle("/city", chain.Then(mainLogicHandler))

	http.ListenAndServe(":8000", nil)

}

func mainLogicPOST(w http.ResponseWriter, r *http.Request) {
	var city City
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&city)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	log.Printf("Got %s city with area of %d sq miles!\n", city.Name, city.Area)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Created!\n"))
}

func filterPOST(handler http.Handler) http.Handler {
	local := func(w http.ResponseWriter, r *http.Request) {

		log.Println("Currently in the check POST middleware")

		// Filtering requests by MIME type
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method not allowed!\n"))
			return
		}

		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(local)
}

func filterContentType(handler http.Handler) http.Handler {
	local := func(w http.ResponseWriter, r *http.Request) {

		log.Println("Currently in the check content type middleware")

		// Filtering requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON\n"))
			return
		}

		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(local)
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	local := func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// Setting cookie to each and every response
		cookie := http.Cookie{
			Name:  "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	}
	return http.HandlerFunc(local)
}
