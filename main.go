package main

import (
	"fmt"
	"net/http"
)

func welcomeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome")
}

func main() {

	http.HandleFunc("/", welcomeFunc)
	http.ListenAndServe(":8080", nil)
}
