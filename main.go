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
    tmpl, err := template.ParseFiles("src/templates/index.html", "src/templates/sidebar.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    err = tmpl.Execute(w, struct{ IsExpanded bool }{IsExpanded: false})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	 tmpl, _ := template.ParseFiles("src/templates/login.html", "src/templates/sidebar.html")
	 tmpl.Execute(w, nil)
}

func storiesPageHandler(w http.ResponseWriter, r *http.Request) {
	 tmpl, _ := template.ParseFiles("src/templates/stories.html", "src/templates/sidebar.html")
	 tmpl.Execute(w, nil)
}

func editorPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ :=template.ParseFiles("src/templates/editor.html", "src/templates/sidebar.html" )
    tmpl.Execute(w, nil)
}

func handleToggleSidebar(w http.ResponseWriter, r *http.Request) {
    isExpanded := r.URL.Query().Get("expanded") == "true"
    tmpl, err := template.ParseFiles("src/templates/sidebar.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = tmpl.ExecuteTemplate(w, "sidebar", struct{ IsExpanded bool }{IsExpanded: !isExpanded})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

/***********************************
          Serves the App 
***********************************/
func main() {
	 fs := http.FileServer(http.Dir("src"))
	 http.Handle("/src/", http.StripPrefix("/src/", fs))

	 http.HandleFunc("/", mainPageHandler)
	 http.HandleFunc("/login", loginPageHandler)
	 http.HandleFunc("/stories", storiesPageHandler)
     http.HandleFunc("/editor", editorPageHandler)
     http.HandleFunc("/toggle-sidebar", handleToggleSidebar)
	 log.Fatal(http.ListenAndServe(":8080", nil))
}




