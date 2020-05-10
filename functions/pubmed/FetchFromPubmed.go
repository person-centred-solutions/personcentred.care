package pubmed

import (
	"github.com/imroc/req"
	"functions/entrez"
)

const ncbiEntrezApiUrl = "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"
const pubmedDatabaseName = "pubmed"

func FetchFromPubmed(searchTerms string) string, err {
	param := req.Param{
		"db": pubmedDatabaseName,
		"term":  searchTerms,
	}
	var pubmedResponse EntrezResponse
	_, err := req.Get(ncbiEntrezApiUrl, params, req.BodyXML(&pubmedResponse))
	if err != nil {
		return "", err
	}
	return pubmedResponse
}
