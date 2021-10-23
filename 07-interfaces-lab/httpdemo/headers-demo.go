package main

import (
	"io"
	"net/http"
)

func main() {
	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/sendstrailers", func(w http.ResponseWriter, req *http.Request) {
	//		// Before any call to WriteHeader or Write, declare
	//		// the trailers you will set during the HTTP
	//		// response. These three headers are actually sent in
	//		// the trailer.
	//		w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	//		w.Header().Add("Trailer", "AtEnd3")
	//		w.Header().Set("AtEnd1", "value 1")
	//		w.Header().Set("AtEnd2", "value 2")
	//		w.Header().Set("AtEnd3", "value 3")
	//		w.Header().Set("Link", `http://example.com/TheBook/chapter2; rel="previous"; title="previous chapter"`)
	//
	//		w.Header().Set("Content-Type", "application/json") // normal header
	//		w.WriteHeader(http.StatusOK)
	//
	//		io.WriteString(w, `
	//{
	//	"username": "john",
	//	"password": "mypass123"
	//}
	//`)
	//
	//	})
	http.HandleFunc("/sendstrailers", func(w http.ResponseWriter, req *http.Request) {
		// Before any call to WriteHeader or Write, declare
		// the trailers you will set during the HTTP
		// response. These three headers are actually sent in
		// the trailer.
		w.Header().Set("Trailer", "AtEnd1, AtEnd2")
		w.Header().Add("Trailer", "AtEnd3")
		w.Header().Set("AtEnd1", "value 1")
		w.Header().Set("AtEnd2", "value 2")
		w.Header().Set("AtEnd3", "value 3")
		w.Header().Set("Link", `http://example.com/TheBook/chapter2; rel="previous"; title="previous chapter"`)

		w.Header().Set("Content-Type", "application/json") // normal header
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, `
{
	"username": "john", 
	"password": "mypass123"
}
`)

	})
	//http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":8080", nil)
}
