package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

//PageVars contains variables to display on page
type PageVars struct {
	Date  string
	Time  string
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
	Secs  int
	Name  string
}

func render(w http.ResponseWriter, tmpl string, def string, pageVars PageVars) {

	tmpl = fmt.Sprintf(filepath.FromSlash("templates/%s"), tmpl)

	t :=
		template.Must(template.ParseFiles(tmpl))

	err := t.ExecuteTemplate(w, def, pageVars)

	if err != nil {
		log.Print("tempalte executiong error : ", err)
	}

}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(filepath.FromSlash("/public")))
	mux.Handle(filepath.FromSlash("/static/"), http.StripPrefix(filepath.FromSlash("/static/"), files))
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/submitShow", SubmitShow)
	server.ListenAndServe()
}

//Home function
func Home(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	HomePageVars := PageVars{
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04:05"),
	}
	render(w, "home.html", "home", HomePageVars)
}
