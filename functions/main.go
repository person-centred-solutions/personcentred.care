package main

import (
	"fmt"
	"log"

	"./entrez"
)

func hello() string {
	return "Hello, world."
}

func main() {
	xmlResp := "<eSearchResult><Count>6</Count><RetMax>6</RetMax><RetStart>0</RetStart><IdList><Id>19008416</Id><Id>18927361</Id></IdList></eSearchResult>"
	jsonTransp, err := entrez.UnmarshallEntrezResponse(xmlResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(jsonTransp)
}
