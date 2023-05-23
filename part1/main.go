package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World") //Writing "Hello World" to ResponseWriter
}

func main() {
	// routes
	// This is associating paths with the different web pages (functions)
	http.HandleFunc("/", helloWorldPage) //path, webPageFunction
	//This code is setting up a ServerMux with associated web pages
	//The basic job of a ServerMux is to route different requests based on its path to different resources in the server

	// Create & Start the server
	// This is creating an instance of the Server and calling the ListenAndServe method
	http.ListenAndServe("", nil) //host:port address, ServeMux
	// By default host is localhost
	// port is 80 (Example port, ":8080")
	// ServeMux is DefaultServeMux

	// When we are associating all the default paths with the different web pages(functions) uisng HandleFunc function,
	// we are essentially sending them to DefaultServeMux
	// and when we are kicking off the server with ListenAndServe function,
	// we are kick-starting the DefaultServerMux

	// So, HandleFunc is setting up the DefaultServeMux
	// And, ListenAndServe is kisck starting the DefaultServeMux

	// The actual kick start is inside: http.ListenAndServe -> server.ListenAndServe -> srv.Serve
}
