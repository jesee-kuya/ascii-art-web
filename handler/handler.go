package handler

import (
	"fmt"
	"html/template"
	"net/http"

	web "web/functions"
)

type Result struct {
	Art string
}

type NewError struct {
	Err string
}
type Data struct {
	Text   string
	Banner string
}

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	var result Result
	var data Data
	t, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Fprint(w, err)
	}

	if r.Method == http.MethodPost {
		if err = r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "500 | Internal Server Error", http.StatusNotFound)
			return
		}
		data.Text = r.FormValue("input")
		if err = web.NonAscii(data.Text); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "400 | Bad Request - Non ascii character detected", http.StatusBadRequest)
			return
		}
		data.Banner = r.FormValue("bannerFiles")
		fileArr, err := web.FileReader(data.Banner)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "404 |  Resource Not Found", http.StatusNotFound)
			return
		}
		result.Art = web.Ascii(data.Text, fileArr)
	}
	if r.URL.Path != "/" && r.URL.Path != "/ascii-web" {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "404 | Page Not Found", http.StatusNotFound)
		return
	}
	t.ExecuteTemplate(w, "index.html", result)
}
