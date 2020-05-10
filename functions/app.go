package main

import (
	"fmt"
	"functions/carecounts"
	"log"
)

func main() {
	publicationCounts, err := carecounts.FetchCareCounts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", publicationCounts)
}
