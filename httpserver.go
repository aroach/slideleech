package main

import (
	"log"
	"net/http"
)

func serveSlides(slideDirectory string) {
	fs := http.FileServer(http.Dir(slideDirectory))
	http.Handle("/", fs)

	log.Println("Listening...")
	log.Println("Open your browser to http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
