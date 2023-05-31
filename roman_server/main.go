package main

import (
	"fmt"
	"html"
	"net/http"
	roman_server "roman_server/roman_numeral"
	"strconv"
	"strings"
	"time"
)

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%q", html.EscapeString(roman_server.Numerals[number]))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
	}

}

func main() {
	http.HandleFunc("/", handleGetRequest)

	// create and run server
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	fmt.Print("Server running on port ", s.Addr)
}
