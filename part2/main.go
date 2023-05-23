package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/deep":
		fmt.Fprint(w, "Hello Deep")
	default:
		fmt.Fprint(w, "Error!")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
	// By default the HTTP method is GET
}

func htmlVsdPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")

	//Here the Content-type is implicitly set to text/html as we are using html tags.
	//We can change it explicitly and set it to plain
	w.Header().Set("Content-Type", "text/plain") //Header is just a map that takes key-value pairs of string type
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did *not* timeout") //This will never happen as the WriteTimeout configured for the custom server is 1s
	// And we are making it sleep for 2 seconds.
	// So, timeout will happen for writing before this is shown.
}

func helloWorldDeepMode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloWorldDeepMode")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello World</h1>")
}

func main() {
	// http.HandleFunc("/", htmlVsdPlain)
	// Anything matching the path "/" will get the same response, "Hello World"
	// Including the path "/deep". So, even if the path partially matches,
	// we will get the same response
	// To prevent that we can use switch-case for different routes inside the helloWorldPage function
	// http.HandleFunc("/timeout", timeout)

	// http.ListenAndServe("", nil)

	// Configure our own custom Server
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000, //1s
		WriteTimeout: 1000,
	}

	// Creating our own custom ServeMux (Multiplexer)
	var muxDeepMode http.ServeMux
	server.Handler = &muxDeepMode                   //Setting our custom server's handler to our Mux
	muxDeepMode.HandleFunc("/", helloWorldDeepMode) //setting up router with our Mux

	server.ListenAndServe()
}

/*
Multiplexer
Router
Handler
*/
