package main

import (
	"fmt"
	"io/ioutil"
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
func Gethandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PARAM := vars["PARAM"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", PARAM)))
}
func Badhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
func Post1handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", body)))
}
func Post2handler(w http.ResponseWriter, r *http.Request) {
	//header, err := ioutil.ReadAll(r.Header)
	ua := r.Header.Get("a")
	ua2 := r.Header.Get("b")

	result, err := strconv.Atoi(ua + ua2)
	if err != nil {
		log.Fatal(err)
	}
	result2 := strconv.Itoa(result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("a+b\n%s", result2)))

}

func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", Gethandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", Badhandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", Badhandler).Methods(http.MethodGet)
	router.HandleFunc("/data", Post1handler).Methods(http.MethodPost)
	router.HandleFunc("/headers", Post2handler).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
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
