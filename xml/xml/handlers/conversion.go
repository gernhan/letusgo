package xml_handlers

import (
	encoding "encoding/xml"
	"fmt"
)

func convertToXML(data interface{}) (string, error) {
	xmlData, err := encoding.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return "", err
	}
	s := string(xmlData)

	return doStandardReplacements(s), nil
}
