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
	logger.Println("Rendering Home.html")
	render(w, "home.html", "home", HomePageVars)
}

//SubmitShow for submittted details
func SubmitShow(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //r is url.Values which is a map[string][]string
	const (
		layoutISO = "2006-01-02 15:04:05"
		layoutUS  = "January 2, 2006"
	)

	var svalues []string

	eName := r.FormValue("name")
	d, _ := time.Parse(layoutISO, r.FormValue("date")+" "+r.FormValue("time"))
	year, mon, day := d.Date()
	p(year)
	p(int(mon))
	p(day)
	p(d.Hour())
	CountPagaeVars := PageVars{Year: year,
		Month: int(mon) - 1,
		Day:   day,
		Hour:  d.Hour(),
		Min:   d.Minute(),
		Secs:  d.Second(),
		Name:  eName}
	fmt.Println(CountPagaeVars)
	for k, v := range r.Form {
		logger.Println(fmt.Sprintf("%s == %s", k, v))
		svalues = append(svalues, v...)
	}
	logger.Println(svalues)
	render(w, "timer.html", "timer", CountPagaeVars)

}
