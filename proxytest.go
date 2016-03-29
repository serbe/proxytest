package main

import (
	"fmt"
	"log"
	"net/http"
)

var headers = []string{
	"HTTP_VIA",
	"HTTP_X_FORWARDED_FOR",
	"HTTP_FORWARDED_FOR",
	"HTTP_X_FORWARDED",
	"HTTP_FORWARDED",
	"HTTP_CLIENT_IP",
	"HTTP_FORWARDED_FOR_IP",
	"VIA",
	"X_FORWARDED_FOR",
	"FORWARDED_FOR",
	"X_FORWARDED",
	"FORWARDED",
	"CLIENT_IP",
	"FORWARDED_FOR_IP",
	"HTTP_PROXY_CONNECTION",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<p>RemoteAddr: %s</p>", r.RemoteAddr)
        for _, header := range headers {
            str := r.Header.Get(header)
            if str != "" {
                fmt.Fprintf(w, "<p>%s: %s</p>", header, str)
            }
        }
    })
	log.Fatal(http.ListenAndServe(":8080", nil))
}
