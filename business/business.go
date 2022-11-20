package business

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DisplayAllSport() []Sport {
	response, err := http.Get("https://api.the-odds-api.com/v4/sports/?apiKey=7b3b6841d51a06e87f48d6b49ebf1bea&all=true")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sport []Sport
	json.Unmarshal(responseData, &sport)
	return sport
}
