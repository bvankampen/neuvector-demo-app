package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type Data struct {
	Backend string
}

type Response struct {
	Text string `json:"text"`
}

func indexHandler(w http.ResponseWriter, req *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/frontend.gohtml",
		"templates/base.gohtml",
	)

	action := req.URL.Query().Get("action")
	data := new(Data)

	resp, err := http.Get(os.Getenv("APP_BACKEND_URL") + "/backend?action=" + action)
	defer resp.Body.Close()

	if err != nil {
		data.Backend = err.Error()
	} else {
		if resp.StatusCode == 200 {
			r := new(Response)
			json.NewDecoder(resp.Body).Decode(r)
			data.Backend = r.Text
		} else {
			data.Backend = "Error " + string(resp.StatusCode) + " " + resp.Status
		}

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
	http.ListenAndServe(os.Getenv("APP_FRONTEND_IP")+":"+os.Getenv("APP_FRONTEND_PORT"), nil)

}
