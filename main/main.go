package main

import (
	"log"
	"net/http"

	"Mygoto.com/web"
)

func main() {
	http.HandleFunc("/", web.Redirect)
	http.HandleFunc("/add", web.Add)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

