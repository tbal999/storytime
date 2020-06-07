package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var story []string

func getPort() string {
	p := os.Getenv("PORT")
	fmt.Println(p)
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func add(w http.ResponseWriter, r *http.Request) {
	story = append(story, mux.Vars(r)["string"])
	for index := range story {
		fmt.Fprintf(w, story[index]+"\n")
	}
}

func front(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The neverending story...")
	for index := range story {
		fmt.Fprintf(w, story[index]+"\n")
	}
}

func main() {
	port := getPort()
	fmt.Println("API has started.")
	fmt.Println("Running on port... " + port)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", front)
	router.HandleFunc("/add/{string}", add).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
