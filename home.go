package main

import (
	"fmt"
	"net/http"
	"time"
)

//HomePage for home page contents
func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	HomePageVars := PageVars{
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04:05"),
	}
	render(w, "home.html", HomePageVars)
}

//SubmitShow for submittted details
func SubmitShow(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //r is url.Values which is a map[string][]string

	var svalues []string
	for k, v := range r.Form {
		fmt.Printf("%s == %s\n", k, v)
		svalues = append(svalues, v...)
	}

	fmt.Println(svalues)

}
