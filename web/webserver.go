package web

import (
	"net/http"
	"fmt"
	"encoding/json"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"os"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

type Route struct {
	route      string
	handler    func(http.ResponseWriter, *http.Request)
	httpMethod string
}

var routes = []Route{
	{
		route:      "/books",
		handler:    booksHandler,
		httpMethod: "GET",
	},
	Route{
		httpMethod:       "GET",
		route:      "/authors",
		handler:  authorsHandler,
	},
	Route{
		httpMethod:       "GET",
		route:      "/authors/{uuid}",
		handler:  getAuthorByUuid,
	},
	Route{
		httpMethod:       "DELETE",
		route:      "/authors/{uuid}",
		handler:  deleteAuthor,
	},
	Route{
		httpMethod:       "POST",
		route:      "/authors",
		handler:  createAuthor,
	},
	Route{
		httpMethod:       "PUT",
		route:      "/authors/{uuid}",
		handler:  updateAuthor,
	},
}


func StartServer() {
	mux := mux.NewRouter()
	for _, route := range routes {
		mux.HandleFunc(route.route, route.handler).Methods(route.httpMethod)
	}

	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}

func serializeData(data interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if data, err := json.Marshal(data); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
		return err
	} else {
		fmt.Fprintln(w, string(data))
		return nil
	}
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading body!\"}")
		return
	}
	var author model.AuthorDto
	if err := json.Unmarshal(body, &author); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error unmarshling the body!\"}")
		return
	}
	model.Authors = append(model.Authors, author)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	uuid := mux.Vars(r)["uuid"]
	w.Header().Set("Content-Type", "application/json")

	var updatedSlice []model.AuthorDto
	for _, author := range model.Authors {
		if author.UUID != uuid {
			updatedSlice = append(updatedSlice, author)
		}
	}
	model.Authors = updatedSlice
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	uuid := mux.Vars(r)["uuid"]
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading body!\"}")
		return
	}
	var newAuthor model.AuthorDto
	if err := json.Unmarshal(body, &newAuthor); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error unmarshling the body!\"}")
		return
	}

	for idx, author := range model.Authors {
		if author.UUID == uuid {
			model.Authors[idx] = newAuthor
			return
		}
	}
	fmt.Fprintln(w, "{\"message\":\"The author does not exist!\"}")
	w.WriteHeader(http.StatusNotFound)
}

func getAuthorByUuid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	uuid := mux.Vars(r)["uuid"]
	w.Header().Set("Content-Type", "application/json")
	for _, author := range model.Authors {
		if author.UUID == uuid {
			if data, err := json.Marshal(author); err != nil {
				fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
				return
			} else {
				fmt.Fprintln(w, string(data))
				return
			}
		}
	}

	fmt.Fprintln(w, "{\"message\":\"The book does not exist!\"}")
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
