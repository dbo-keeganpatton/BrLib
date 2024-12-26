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
	 tmpl, _ := template.ParseFiles("./web/templates/index.html")
	 tmpl.Execute(w, nil)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	 tmpl, _ := template.ParseFiles( "./web/templates/login.html")
	 tmpl.Execute(w, nil)
}

func storiesHandler(w http.ResponseWriter, r *http.Request) {
	 tmpl, _ := template.ParseFiles("./web/templates/stories.html")
	 tmpl.Execute(w, nil)
}


/***********************************
          Serves the App 
***********************************/
func main() {

	 fs := http.FileServer(http.Dir("./web/static"))
	 http.Handle("/static/", http.StripPrefix("/static/", fs))


	 http.HandleFunc("/", mainPageHandler)
	 http.HandleFunc("/login", loginPageHandler)
	 http.HandleFunc("/stories", storiesHandler)

	 log.Fatal(http.ListenAndServe(":8080", nil))

}




