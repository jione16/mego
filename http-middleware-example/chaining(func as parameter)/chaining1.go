package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	route := MyTypeOfGirl{name: "hard to spread", age: 21}
	http.Handle("/ohno", isFemale("male", isCute(true, &route)))
	http.Handle("/notcute", isFemale("female", isCute(false, &route)))
	http.Handle("/pass", isFemale("female", isCute(true, &route)))
	log.Fatal(http.ListenAndServe(":1000", nil))

}

//declaring a struct to implement ServeHTTP method
type MyTypeOfGirl struct {
	name string
	age  int
}

//so myTypeOfGirl now is implemented Handler interface
func (h *MyTypeOfGirl) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.Write([]byte("Her name is " + h.name + " and she's " + strconv.Itoa(h.age) + " years old."))
}

//gender validation middleware
func isFemale(gender string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if gender == "female" {
			next.ServeHTTP(writer, request)
		} else {
			writer.Write([]byte("No no..~"))
		}
	})
}

//is cute validation middleware
func isCute(isCute bool, next http.Handler) http.Handler {
	/*handleFunc is an interface implemented ServeHTTP method*/
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if isCute {
			next.ServeHTTP(writer, request)
		} else {
			writer.Write([]byte("Hmm do you have heart?"))
		}
	})
}
