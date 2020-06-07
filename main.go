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

	t, err := template.ParseFiles(tmpl)

	if err != nil {
		log.Print("tempalte parsing error : ", err)
	}

	err = t.Execute(w, pageVars)

	if err != nil {
		log.Print("tempalte executiong error : ", err)
	}

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/submitShow", SubmitShow)
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
