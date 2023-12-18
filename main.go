package main

/*
https://www.sohamkamani.com/golang/how-to-build-a-web-application/
https://freshman.tech/web-development-with-go/
http://127.0.0.1:8080/resources/html/tables-basic.html
*/

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func templating(w http.ResponseWriter, r *http.Request) {
	fmt.Println("templating...")
	tpl := template.Must(template.ParseFiles("./resources/app/page.html"))
	tpl.Execute(w, nil)
}

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/app/pages/main", templating).Methods("GET")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./resources/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/app/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/app/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}