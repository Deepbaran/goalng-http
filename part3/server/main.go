package main

import (
	"io/ioutil"
	"net/http"
)

//Server

func url(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	path := r.URL.Path
	r.ParseForm()
	w.Write([]byte("http://" + host + path + "?name=" + r.Form.Get("name")))
}

func body(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body, _ := ioutil.ReadAll(r.Body)
	w.Write([]byte(string(body)))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
		// http://localhost/url?name=Deep
	case "/body":
		body(w, r)
		// http://localhost/body
		// with body in json: {"name": "Deep"}
	default:
		w.Write([]byte("Hello World"))
		// http://localhost
	}
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe("", nil)
}
