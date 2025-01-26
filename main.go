package main

import (
	"data"
	"fmt"
	"net/http"
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
