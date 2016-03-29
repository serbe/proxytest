package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Hello, %q</p>", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "<p>X-Forwarded-For: %s</p>", r.Header.Get("X-Forwarded-For"))
		fmt.Fprintf(w, "<p>X-Forwarded-Host: %s</p>", r.Header.Get("X-Forwarded-Host"))
		fmt.Fprintf(w, "<p>Forwarded: %s</p>", r.Header.Get("Forwarded"))
		fmt.Fprintf(w, "<p>X-Forwarded-For: %s</p>", r.Header.Get("Via"))
		fmt.Fprintf(w, "<p>RemoteAddr: %s</p>", r.RemoteAddr)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
