package carecounts

import (
	"functions/pubmed"
	"log"
	"strings"
)

// FetchCareCounts fetches counts of publications on person and patient centred cares
func FetchCareCounts() (CareCounts, error) {
	personCentredTerms := []string{"person-centred", "person-centered", "person centred", "person centered"}

	// patientCentredTerms := []string{"patient-centred", "patient-centered", "patient centred", "patient centered"}

	var output CareCounts

	log.Printf("%v", constructSearchTerm(personCentredTerms))
	return output, nil
}

func fetchPubmedCount(terms string) (uint32, error) {
	pubmedResponse, err := pubmed.FetchFromPubmed(terms)
	if err != nil {
		return 0, err
	}
	return pubmedResponse.Count, nil
}

func constructSearchTerm(alternativeTerms []string) string {
	var termsTransformed []string

	for _, term := range alternativeTerms {
		termTransformed := "\"" + term + "\"[All Terms]"
		termsTransformed = append(termsTransformed, termTransformed)
	}

	return strings.Join(termsTransformed, " OR ")
}
