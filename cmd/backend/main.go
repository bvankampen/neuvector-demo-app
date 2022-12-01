package main

import (
	"html/template"
	"net/http"
	"os"
)

type Data struct {
	Text string
}

func indexHandler(w http.ResponseWriter, req *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/backend.gohtml",
		"templates/base.gohtml",
	)

	req.ParseForm()
	data := Data{
		Text: req.Form.Get("text"),
	}

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
	http.ListenAndServe(os.Getenv("APP_BACKEND_IP")+":"+os.Getenv("APP_BACKEND_PORT"), nil)
}
