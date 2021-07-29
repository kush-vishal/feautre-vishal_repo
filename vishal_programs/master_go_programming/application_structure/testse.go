package main

import (
	"log"
	"net/http"
)

func ServeHTTP2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi vishal"))

}
func main() {
	http.HandleFunc("/", ServeHTTP2)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
