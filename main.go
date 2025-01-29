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
	servr := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/", sayHello)
	data.CheckError(servr.ListenAndServe())

	mux.HandleFunc("/login/", login)

}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t, err := t.ParseFiles("./static/index.html")
	data.CheckError(err)
	t.Execute(w, nil)
}
