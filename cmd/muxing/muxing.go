package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	r := mux.NewRouter()

	r.HandleFunc("/name/{param}", nameHandler).Methods(http.MethodGet)
	r.HandleFunc("/bad", badHandler).Methods(http.MethodGet)
	r.HandleFunc("/data", dataHandler).Methods(http.MethodPost)
	r.HandleFunc("/headers", headersHandler).Methods(http.MethodPost)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello, %s", mux.Vars(r)["param"])))
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	w.Write([]byte("I got message:\n"))
	w.Write([]byte(string(body)))
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	c := strconv.Itoa(a + b)

	w.Header().Set("a+b", c)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
