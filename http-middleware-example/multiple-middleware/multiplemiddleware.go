package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.
		Handle("/", middleware(http.
			HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Hello"))
			}),
			firstMiddleware,
			secondMiddleware))
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func middleware(handler http.Handler, middleware ...func(header http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		handler = mw(handler)
	}
	return handler
}

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//logic
		fmt.Println("fist middleware checking validation")
		pass := true
		if pass {
			fmt.Println("fist middleware passed")
			next.ServeHTTP(writer, request)
		} else {
			writer.Write([]byte("fist middleware not pass"))
			return
		}
	})
}

func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//logic
		fmt.Println("second middleware checking validation")
		pass := false
		if pass {
			fmt.Println("second middleware passed")
			next.ServeHTTP(writer, request)
		} else {
			writer.Write([]byte("second middleware not pass"))
			return
		}
	})
}
