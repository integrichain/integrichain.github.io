package main

import (
	"log"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Rendering index.html")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/mail/", http.StripPrefix("/mail/", http.FileServer(http.Dir("mail"))))
	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}