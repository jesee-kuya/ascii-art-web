package webart

import (
	"html/template"
	"net/http"

	webart "webart/ascii"
)

type Result struct {
	Result string
}

type Error_template struct {
	Title   string
	Message string
	Status  int
}

// ProcessForm handles the data passed to the form
func ProcessForm(w http.ResponseWriter, r *http.Request, filename string) {
	var art string
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		ShowError(w, "server error", " Page Not Found", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		ShowError(w, "server error", "Internal Server Error", http.StatusInternalServerError)
		return
	}
	text := r.Form.Get("input")
	banner := r.Form.Get("banner")
	Bannercontent, err := webart.FileReader(banner)
	if err != nil {
		ShowError(w, "server error", "Resource not found or damaged", http.StatusNotFound)
		return
	}
	if !webart.CheckAscii(text) {
		ShowError(w, "Bad request", "Bad request", http.StatusBadRequest)
		return
	}
	art = webart.Ascii(Bannercontent, text)
	data := Result{Result: art}
	w.WriteHeader(http.StatusOK)
	t.Execute(w, data)
}

// ShowError handles all errors encountered
func ShowError(writer http.ResponseWriter, title string, message string, code int) {
	tmpl, _ := template.ParseFiles("template/error.html")
	Error_case := Error_template{Title: title, Message: message, Status: code}
	writer.WriteHeader(code)
	tmpl.Execute(writer, Error_case)
}
