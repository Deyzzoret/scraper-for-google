package constants

// Specify the domains we will search on
 var GoogleDomains = map[string]string {

 }


// Structure speficying the expected result from scraper operation
 type SearchResult struct {
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
 } 

//  Save list of user agents used to launch request
 var UserAgents = []string {
	"tst",
 }