package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/mahmoud-shabban/simpleWebServer/data"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Wrold!")
}
func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
	mux.HandleFunc("/", sayHello)

	mux.HandleFunc("/login/", login)

	data.CheckError(server.ListenAndServe())
}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t, err := t.ParseFiles("static/index.html")
	data.CheckError(err)
	t.Execute(w, nil)
}
