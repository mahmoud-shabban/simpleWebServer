package main

import (
	"fmt"
	"net/http"

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

}
