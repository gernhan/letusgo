package schema

import (
	"encoding/xml"
)

// FOX is Root element (Wurzel-Element bei Stand-Alone-Übergabe z.B. in der DOR)
type FOX *FOXType

// TEXTFRAME ...
type TEXTFRAME struct {
	IDAttr         string `xml:"ID,attr,omitempty"`
	XAttr          int    `xml:"X,attr"`
	YAttr          int    `xml:"Y,attr"`
	WIDTHAttr      int    `xml:"WIDTH,attr"`
	HEIGHTAttr     int    `xml:"HEIGHT,attr,omitempty"`
	HORZALIGNAttr  string `xml:"HORZALIGN,attr,omitempty"`
	VERTALIGNAttr  string `xml:"VERTALIGN,attr,omitempty"`
	POSTUPDATEAttr bool   `xml:"POSTUPDATE,attr,omitempty"`
	*FOXType
}

// CharStyles is If specified the SPAN content is generated in run time in this frame. Existing content is thus deleted.  (bei Angabe wird der SPAN-Inhalt zur Laufzeit aus diesem Feldnamen generiert. ein vorgegebener Inhalt wird damit gelöscht).
type CharStyles struct {
	XMLName        xml.Name    `xml:"charStyles"`
	FONTNAMEAttr   string      `xml:"FONTNAME,attr,omitempty"`
	FONTSIZEAttr   interface{} `xml:"FONTSIZE,attr,omitempty"`
	WEIGHTAttr     string      `xml:"WEIGHT,attr,omitempty"`
	ITALICAttr     bool        `xml:"ITALIC,attr,omitempty"`
	FIELDNAMEAttr  string      `xml:"FIELDNAME,attr,omitempty"`
	VERTOFFSETAttr interface{} `xml:"VERTOFFSET,attr,omitempty"`
	SPACINGAttr    interface{} `xml:"SPACING,attr,omitempty"`
}

// ParStyles is paragraph related style attribute (absatzbezogene Style-Attribute)
type ParStyles struct {
	XMLName         xml.Name    `xml:"parStyles"`
	ALIGNMENTAttr   string      `xml:"ALIGNMENT,attr,omitempty"`
	HORZRULEAttr    string      `xml:"HORZRULE,attr,omitempty"`
	MINFONTSIZEAttr interface{} `xml:"MINFONTSIZE,attr,omitempty"`
	HORZCLIPAttr    bool        `xml:"HORZCLIP,attr,omitempty"`
	LINESPACEAttr   interface{} `xml:"LINESPACE,attr,omitempty"`
}

// P ...
type P struct {
	CharStyles *CharStyles
	ParStyles  *ParStyles
	*FOXType
}

// SPAN ...
type SPAN struct {
	CharStyles *CharStyles
	ParStyles  *ParStyles
	*FOXType
}

// BR ...
type BR struct {
}

// B is Paragraph in bold face (Fetter Textabschnitt)
type B *FOXType

// I is Paragraph in italics (Kursiver Textabschnitt)
type I *FOXType

// SUB is Subscript (tiefgestellter Textabschnitt)
type SUB *FOXType

// SUP is Superscript (hochgestellter Textabschnitt)
type SUP *FOXType

// FOXType is Content type "FOX Format (Inhalts-Typ "FOX-Format")
type FOXType struct {
	SPAN []*SPAN    `xml:"SPAN"`
	P    []*P       `xml:"P"`
	BR   []*BR      `xml:"BR"`
	B    []*FOXType `xml:"B"`
	I    []*FOXType `xml:"I"`
	SUB  []*FOXType `xml:"SUB"`
	SUP  []*FOXType `xml:"SUP"`
}

