// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	var nextPage string	
	var page int
	var minify bool
	fmt.Println("Do you want to enter the new page (y/n)?")
	fmt.Scanln(&nextPage)		
	for nextPage == "y"{		
		fmt.Println("Enter the page number)?")
		fmt.Scanln(&page)
		fmt.Println("Do you want to truncate the larger fields or not (True/False)?")		
		fmt.Scanln(&minify)
		fmt.Printf(
			"Query Credits: %d\nScan Credits:  %d\n\n",
			info.QueryCredits,
			info.ScanCredits)

		hostSearch, err := s.HostSearch(os.Args[1],page,minify)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("Host Data Dump\n")
		for _, host := range hostSearch.Matches {
			fmt.Println("==== start ",host.IPString,"====")
			h,_ := json.Marshal(host)
			fmt.Println(string(h))
			fmt.Println("==== end ",host.IPString,"====")
			//fmt.Println("Press the Enter Key to continue.")
			//fmt.Scanln()
		}


		fmt.Printf("IP, Port, City\n")

		for _, host := range hostSearch.Matches {
			fmt.Printf("%s, %d, %s, %f, %f\n", host.IPString, host.Port, host.Location.City, host.Location.Latitude, host.Location.Longitude)
		}		
		fmt.Println("Do you want to enter the new page (y/n)?")
		fmt.Scanln(&nextPage)		
	}	
}