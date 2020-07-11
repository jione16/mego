package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	subRouter := router.PathPrefix("/auth").Subrouter()
	subRouter.Use(authMiddleware)
	subRouter.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/user"))
	})
	subRouter.HandleFunc("/admin", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/admin"))
	})
	log.Fatal(http.ListenAndServe(":1000", router))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Auth middleware validation")
		next.ServeHTTP(writer, request)
	})
}
