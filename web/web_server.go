package web

import (
	"fmt"
	"net/http"

	gotov1 "Mygoto.com/goto_v1"
)

const AddForm = `
	<form method="POST" action="/add">
	URL: <input type="text" name="url">
	<input type="submit" value="Add">
	</form>`

var store = gotov1.NewURLStore()

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}