package main

import (
	"github.com/amerine/fake"
	"github.com/amerine/ti3/server"
	"io"
	"log"
	"net/http"
)

func FakeRoot(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
	}
	io.WriteString(w, fake.FullName())
}

func main() {
	mux := server.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
