package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Data struct {
	Backend string
	Error   string
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

	if action != "" {

		url := strings.Trim(os.Getenv("APP_BACKEND_URL"), " ") + "/backend?action=" + action
		fmt.Printf("Getting: %s\n", url)
		resp, err := http.Get(url)

		//defer func(Body io.ReadCloser) {
		//	err := Body.Close()
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//}(resp.Body)

		if err != nil {
			data.Error = err.Error()
			fmt.Printf("Error: %s", err.Error())
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				r := new(Response)
				err := json.NewDecoder(resp.Body).Decode(r)
				if err != nil {
					fmt.Println(err)
				}
				data.Backend = r.Text
			} else {
				data.Backend = "Error " + string(rune(resp.StatusCode)) + " " + resp.Status
			}

		}

		fmt.Printf("Result: %s\n", data.Backend)

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
	ip := strings.Trim(os.Getenv("APP_FRONTEND_IP"), " ")
	port := strings.Trim(os.Getenv("APP_FRONTEND_PORT"), " ")
	err := http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}

}
