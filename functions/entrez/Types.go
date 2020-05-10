package entrez

import "encoding/xml"

type EntrezResponse struct {
	XMLName  xml.Name `xml:"eSearchResult" json:"-"`
	Count    uint32   `xml:"Count" json:"Count"`
	RetMax   uint32   `xml:"RetMax" json:"RetMax"`
	RetStart uint32   `xml:"RetStart" json:"RetStart"`
	IdList   IdList   `xml:"IdList" json:"ids"`
}

type IdList struct {
	Items []uint32 `xml:"Id" json:"list"`
}
