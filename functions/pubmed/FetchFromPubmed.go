package pubmed

import (
	"functions/entrez"

	"github.com/imroc/req"
)

const ncbiEntrezAPIURL = "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"
const pubmedDatabaseName = "pubmed"

func FetchFromPubmed(searchTerms string) (entrez.EntrezResponse, error) {
	params := req.Param{
		"db":   pubmedDatabaseName,
		"term": searchTerms,
	}
	var pubmedResponse entrez.EntrezResponse
	_, err := req.Get(ncbiEntrezAPIURL, params, req.BodyXML(&pubmedResponse))
	if err != nil {
		return "", err
	}
	return pubmedResponse, nil
}
