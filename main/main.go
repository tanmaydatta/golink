package main

import (
	"fmt"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("url: %v\n", r.URL)
	http.Redirect(w, r, fmt.Sprintf("http://go.grab.com%s", r.URL), 301)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
