package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"mirea.com/web-prog/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/get", handlerGet).Methods(http.MethodGet)
	router.HandleFunc("/delete/{id}", handleDelete).Methods(http.MethodDelete)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	corsObj := handlers.AllowedOrigins([]string{"*"})


	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, corsObj)(router)))
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	var (
		raw []byte
		err error
	)

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "error: %v", err.Error())
		} else {
			_, _ = fmt.Fprintf(w, "%v", string(raw))
		}
	}()

	raw, err = model.GetCollection()
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "error: %v", err.Error())
		}
	}()

	rawID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return
	}

	model.Remove(id)
}