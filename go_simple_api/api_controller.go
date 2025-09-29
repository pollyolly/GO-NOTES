//Go File in package main are linked together 
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:Title`
	Desc string `json:Desc`
	Content string `json:Content`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles {
		Article{Title:"Test title", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", allArticles)
	//log.Fatal(http.ListenAndServe(":8081", nil))
}
