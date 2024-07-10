package main

import (
	"html/template"
	"net/http"

	webart "webart/server"
)

func Asciiweb(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case http.MethodGet:
			tmpl, err := template.ParseFiles("template/index.html")
			if err != nil {
				webart.ShowError(w, "server error", "500 | Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, nil)
		case http.MethodPost:
			webart.ProcessForm(w, r, "template/index.html")
		default:
			http.Error(w, "405 | method not allowed", http.StatusMethodNotAllowed)
			return
		}
	} else {
		webart.ShowError(w, "page not found", "404 | Page Not Found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", Asciiweb)
	http.ListenAndServe(":7050", nil)
}
