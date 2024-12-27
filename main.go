package main

import (
	"html/template"
	"log"
	"net/http"
)



/***********************************
            Page Routing 
***********************************/
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	 tmpl, _ := template.ParseFiles("src/templates/index.html")
	 tmpl.Execute(w, nil)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	 tmpl, _ := template.ParseFiles( "src/templates/login.html")
	 tmpl.Execute(w, nil)
}

func storiesHandler(w http.ResponseWriter, r *http.Request) {
	 tmpl, _ := template.ParseFiles("src/templates/stories.html")
	 tmpl.Execute(w, nil)
}


/***********************************
          Serves the App 
***********************************/
func main() {
	 fs := http.FileServer(http.Dir("src"))
	 http.Handle("/src/", http.StripPrefix("/src/", fs))

	 http.HandleFunc("/main", mainPageHandler)
	 http.HandleFunc("/login", loginPageHandler)
	 http.HandleFunc("/stories", storiesHandler)

	 log.Fatal(http.ListenAndServe(":8080", nil))
}




