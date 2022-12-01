package main

import (
	"html/template"
	"net/http"
	"os"
)

type Data struct {
	Backend string
}

func indexHandler(w http.ResponseWriter, req *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/frontend.gohtml",
		"templates/base.gohtml",
	)

	data := Data{Backend: os.Getenv("APP_BACKEND_URL")}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	static := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", static))
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(os.Getenv("APP_FRONTEND_IP")+":"+os.Getenv("APP_FRONTEND_PORT"), nil)

}
