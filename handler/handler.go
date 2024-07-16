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
			fmt.Fprint(w, err)
		}
		data.Text = r.FormValue("input")
		data.Banner = r.FormValue("bannerFiles")
		fileArr, _ := web.FileReader(data.Banner)
		result.Art = web.Ascii(data.Text, fileArr)
		t.Execute(w, result)
	}
}
