package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func main() {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			enc := json.NewEncoder(w)
			w.Header().
				Set("Content-Type",
					"application/json; charset=utf-8")
			resp := Resp{
				Language:    "English",
				Translation: "Hello",
			}
			if err := enc.Encode(resp); err != nil {
				panic("unable to encode response: " + err.Error())
			}
		})

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
