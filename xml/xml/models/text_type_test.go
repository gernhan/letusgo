package xml_test

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type TextType struct {
	Content string `xml:",innerxml"`
	ID      string `xml:"ID,attr,omitempty"`
	Pos     string `xml:"Pos,attr,omitempty"`
	Ref     string `xml:"Ref,attr,omitempty"`
}

func TestTextTypeXML(t *testing.T) {
	// Sample TextType data
	text := TextType{
		Content: `This is the first line.
							And this is the second line.`,
		ID:  "text1",
		Pos: "top",
		Ref: "reference1",
	}

	// Marshal TextType to XML
	xmlData, err := xml.MarshalIndent(text, "", "    ")
	if err != nil {
		t.Errorf("Error marshaling TextType to XML: %v", err)
	}

	fmt.Println(string(xmlData))
}
