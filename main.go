package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

//PageVars contains varailbes to display on page
type PageVars struct {
	Date string
	Time string
}

func render(w http.ResponseWriter, tmpl string, pageVars PageVars) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)

	t :=
		template.Must(template.ParseFiles(tmpl))

	err := t.ExecuteTemplate(w, "home", pageVars)

	if err != nil {
		log.Print("tempalte executiong error : ", err)
	}

}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
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
	render(w, "home.html", HomePageVars)
}
