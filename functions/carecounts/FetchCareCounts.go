package carecounts

import (
	"functions/pubmed"
	"strings"
)

// FetchCareCounts fetches counts of publications on person and patient centred cares
func FetchCareCounts() (CareCounts, error) {
	personCentredCount, err := fetchPersonCentredCount()
	if err != nil {
		return CareCounts{}, err
	}

	patientCentredCount, err := fetchPatientCentredCount()
	if err != nil {
		return CareCounts{}, err
	}

	var output CareCounts
	output.PatientCentredCarePublications = patientCentredCount
	output.PersonCentredCarePublications = personCentredCount

	return output, nil
}

func fetchPersonCentredCount() (uint32, error) {
	personCentredTerms := []string{"person-centred", "person-centered", "person centred", "person centered"}
	return fetchPubmedCount(constructSearchTerm(personCentredTerms))
}

func fetchPatientCentredCount() (uint32, error) {
	patientCentredTerms := []string{"patient-centred", "patient-centered", "patient centred", "patient centered"}
	return fetchPubmedCount(constructSearchTerm(patientCentredTerms))
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
