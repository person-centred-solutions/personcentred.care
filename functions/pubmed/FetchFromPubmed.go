package pubmed

import (
	"functions/entrez"

	"github.com/imroc/req"
)

const ncbiEntrezAPIURL = "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"
const pubmedDatabaseName = "pubmed"

// FetchFromPubmed performs article search in pubmed database
func FetchFromPubmed(searchTerms string) (entrez.EntrezResponse, error) {
	params := req.Param{
		"db":   pubmedDatabaseName,
		"term": searchTerms,
	}
	response, err := req.Get(ncbiEntrezAPIURL, params)

	var pubmedResponse entrez.EntrezResponse
	response.ToXML(&pubmedResponse)
	if err != nil {
		return entrez.EntrezResponse{}, err
	}
	return pubmedResponse, nil
}
