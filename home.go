package main

import (
	"net/http"
	"time"
)

//HomePage for home page contents
func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	HomePageVars := PageVars{
		Date: now.Format("2020-01-01"),
		Time: now.Format("15:04:05"),
	}
	render(w, "home.html", HomePageVars)
}
