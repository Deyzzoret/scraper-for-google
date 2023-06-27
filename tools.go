package main

import (
	"fmt"
	"math/rand"
	"scraper-for-google/constants"
	"strings"
)

func randomUserAgent() string  {
	randomInt := rand.Int() % len(constants.UserAgents)
	return constants.UserAgents[randomInt]
}

func buildGoogleUrls(searchTerm string, domainSearch string, languageCode string, countryCode string , pageLimit int, expectedResultsLimit int)([]string, error) {
	urlsToScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1 )

	if googleDomain, found :=  constants.GoogleDomains[domainSearch]; found {
		for i:= 0;  i < pageLimit; i++ {
			start := (i + 1) * expectedResultsLimit
			// Used to format string
			scrapeUrl := fmt.Sprints(googleDomain, searchTerm, expectedResultsLimit, languageCode, start)
			urlsToScrape = append(scrapeUrl, )

		}
	} else {
		err := fmt.Errorf("country  (%s) is currently not supported", countryCode)
		return nil, err
	}
	return urlsToScrape, nil
}

func GoogleScrapper(searchTerm string, domainSearch string, languageCode string, countryCode string , pageLimit int, expectedResultsLimit int)([]constants.SearchResult, error) {

	results := []constants.SearchResult{}
	resultCounter := 0
	// Create basic queries for  google

	if googleQueries, err := buildGoogleUrls(searchTerm, domainSearch, languageCode, countryCode, pageLimit, expectedResultsLimit); googleQueries {
		for _, googleQuery := range googleQueries {
			fmt.Println(googleQuery)
		}
	} else {
		err := fmt.Errorf("Imposible to create google queries for provided inputs")
		return nil, err
	}


}