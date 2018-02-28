package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"projects/goworkshop/model"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/authors", authorsHandler)
	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
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