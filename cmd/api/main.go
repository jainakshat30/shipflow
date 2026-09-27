package main

import (
	"fmt"
	"log"
	"net/http"
)

// newRouter builds the app's routes. Returning http.Handler (not *ServeMux)
// lets middleware wrap the mux later without changing callers.
func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	return mux
}

func main() {
	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

// log.Fatal(http.ListenAndServe(":8080", nil)) <-- this is wrong way of doing it because it never compiles, cause only declaration can live at the top level of the package also nil would make it use the default mux which is not what we want, we want to use our own mux
