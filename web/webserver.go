package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"io/ioutil"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

type Route struct {
	route string
	handler func(http.ResponseWriter, *http.Request)
	httpMethod string
}

var routes = []Route {
	{
		route: "/books",
		handler: booksHandler,
		httpMethod: "GET",
	},
}

func StartServer() {
	mux := mux.NewRouter()

	for _, route := range routes{
		mux.HandleFunc(route.route, route.handler).Methods(route.httpMethod)
	}
	mux.HandleFunc("/books", booksHandler).Methods("GET")
	mux.HandleFunc("/books/{uuid}", getBookByUuid).Methods("GET")
	mux.HandleFunc("/authors", createAuthor).Methods("POST")
	mux.HandleFunc("/authors", authorsHandler).Methods("GET")

	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error unmar the body!\"}")
		return
	}
	var author model.AuthorDto
	if err := json.Unmarshal(&body); err != nil{

	}
}

func getBookByUuid(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	uuid := mux.Vars(r)["uuid"]
	for _, book := range model.Books {
		if book.UUID == uuid {
			add, err := json.Marshal(book)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(w, string(add))
			return
		}
	}
	fmt.Fprintln(w, "no book added")
	w.WriteHeader(http.StatusNotFound)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serializedContent, err := json.Marshal(model.Books)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(serializedContent))
}

func authorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serializedContent, err := json.Marshal(model.Authors)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(serializedContent))
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}