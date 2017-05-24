package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"strconv"
	"encoding/json"
	"github.com/gocarina/gocsv"
	"github.com/levigross/grequests"
)

type Location struct {
	Zip string `csv:"Postal Code"`
	Town string `csv:"Place Name"`
	State string `csv:"State Abbreviation"`
	NotUsed string `csv:"-"`
}

var matched []Location

func main() {
	csvPtr := flag.String("csv", "", "CSV location")
	zipPtr := flag.String("zip", "", "Zipcode to check")
	statePtr := flag.String("state", "", "State to check")
	flag.Parse()

	zipsFile, err := os.OpenFile(*csvPtr, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer zipsFile.Close()

	locations := []*Location{}

	if err := gocsv.UnmarshalFile(zipsFile, &locations); err != nil { // Load clients from file
		panic(err)
	}

	for _, location := range locations {
		if location.Zip != "" {
			if *zipPtr != "" && location.Zip == *zipPtr {
				// fmt.Println("Zip:", location.Zip)
				checkFIOS(location)
			} else if *statePtr  != "" && location.State == *statePtr {
				// fmt.Println("State:", location.Zip)
				checkFIOS(location)
			}
		}
	}

	if len(matched) > 0 {
		fmt.Println("Found Service in these locations:")

		for _, location := range matched {
			fmt.Println(location.Zip, ":", location.Town, location.State)
		}
	} else {
		fmt.Println("No service found.")
	}
}

type Hits struct {
	Hits struct {
		Total int `json:"total"`
		Hits []interface{} `json:hits`
	} `json:"hits"`
}

func checkFIOS(location *Location) {
	url := "https://www.verizon.com/foryourhome/ordering/Services/getCity"

	values := map[string]string{"zipCode": location.Zip}
	headers := map[string]string{"Content-Type": "application/json"}

	ro := &grequests.RequestOptions{
		JSON:    values,
		Headers: headers,
	}
	res, err := grequests.Post(url, ro)

	if err != nil || res.Error != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Status != 200 - ", res.StatusCode)
	}

	//"{\"hits\":{\"total\":1,\"hits\":[{\"fields\":{\"ZIP\":[\"01864\"],\"ALTCITY\":[\"NORTH READING\"],\"STATE\":[\"MA\"],\"CITY\":[\"NORTH READING, MA\"]}}]}}" 
	str := res.String()
	unstr, _ := strconv.Unquote(str)
	var content Hits
	err = json.Unmarshal([]byte(unstr), &content)

	if err != nil {
		log.Fatal("Err:", err)
	}

	// fmt.Println(content.Hits.Total)
	if content.Hits.Total > 0 {
		matched = append(matched, *location)
	}
}
