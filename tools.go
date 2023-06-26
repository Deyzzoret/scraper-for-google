package main

import (
	"math/rand"
	"scraper-for-google/constants"
)

func randomUserAgent() string  {
	randomInt := rand.Int() % len(constants.UserAgents)
	return constants.UserAgents[randomInt]
}