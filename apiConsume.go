package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// A Response struct to map the Entire Response
type Response struct {
	Summary     GlobalSummary    `json:"Global"`
	CountryData []CountrySummary `json:"Countries"`
}

// A Country struct to map every country count
type CountrySummary struct {
	Name      string `json:"Country"`
	Code      string `json:"CountryCode"`
	Confirmed int    `json:"TotalConfirmed"`
	Death     int    `json:"TotalDeaths"`
	Recovered int    `json:"TotalRecovered"`
}

// A struct to map global count for cases
type GlobalSummary struct {
	Confirmed int `json:"TotalConfirmed"`
	Death     int `json:"TotalDeaths"`
	Recovered int `json:"TotalRecovered"`
}

func main() {
	country_name := flag.String("country", "India", "a string")

	flag.Parse()

	response, err := http.Get("https://api.covid19api.com/summary")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	//fmt.Println(responseObject)
	fmt.Println("Total Cases summary of world ")
	fmt.Printf("%+v\n", responseObject.Summary)

	countries := responseObject.CountryData

	fmt.Println("Summary of", *country_name)

	for i := range countries {
		if strings.EqualFold(countries[i].Name, *country_name) {
			fmt.Printf("%+v\n", countries[i])
		}
	}

}
