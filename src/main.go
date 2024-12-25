package main

import (
	"html/template"
	"log"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	
	tmpl, _ := template.ParseFiles("./templates/main.html")
	tmpl.Execute(w, nil)

}


func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}




