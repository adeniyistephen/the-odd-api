package main

import (
	"fmt"
	"net/http"

	"github.com/adeniyistephen/the-odd-api/business"
)

func main() {
	http.HandleFunc("/", DisplayAllSport)
	http.ListenAndServe(":8080", nil)
}

func DisplayAllSport(w http.ResponseWriter, r *http.Request) {
	sport := business.DisplayAllSport()
	fmt.Fprintln(w, "List of Current Betting Sports from the-odd-api.com")
	for i := 1; i < len(sport); i++{
		fmt.Fprintln(w, sport[i])
	}
}
