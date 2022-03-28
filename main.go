package main

import (
	"fmt"
	"net/http"
)

func welcomeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "<h1>"+r.URL.Path+"</h1>")
}

func main() {

	http.HandleFunc("/", welcomeFunc)
	http.ListenAndServe(":3000", nil)
}
