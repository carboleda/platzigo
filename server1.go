//Web server (Reto)

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler);
	http.HandleFunc("/google", redirect);
	log.Fatal(http.ListenAndServe("localhost:8000", nil));
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path);
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://www.google.com", 301);
}