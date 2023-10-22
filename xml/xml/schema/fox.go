package schema

import "encoding/xml"

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
	ALIGNMENTAttr   string      `xml:"ALIGNMENT,attr,omitempty"`
	HORZRULEAttr    string      `xml:"HORZRULE,attr,omitempty"`
	MINFONTSIZEAttr interface{} `xml:"MINFONTSIZE,attr,omitempty"`
	HORZCLIPAttr    bool        `xml:"HORZCLIP,attr,omitempty"`
	LINESPACEAttr   interface{} `xml:"LINESPACE,attr,omitempty"`
}

// P ...
type P struct {
	*CharStyles
	*ParStyles
	*FOXType
	InnerText string `xml:",chardata"`
}

// SPAN ...
type SPAN struct {
	*CharStyles
	*ParStyles
	*FOXType
	InnerText string `xml:",chardata"`
}

// BR ...
type BR struct{}

// B is Paragraph in bold face (Fetter Textabschnitt)
type B struct {
	*CharStyles
	*ParStyles
	*FOXType
	InnerText string `xml:",chardata"`
}

// I is Paragraph in italics (Kursiver Textabschnitt)
type I struct {
	*CharStyles
	*ParStyles
	*FOXType
	InnerText string `xml:",chardata"`
}

// SUB is Subscript (tiefgestellter Textabschnitt)
type SUB struct {
	*CharStyles
	*ParStyles
	*FOXType
	InnerText string `xml:",chardata"`
}

// SUP is Superscript (hochgestellter Textabschnitt)
type SUP struct {
	*CharStyles
	*ParStyles
	*FOXType
	InnerText string `xml:",chardata"`
}

// FOXType is Content type "FOX Format (Inhalts-Typ "FOX-Format")
type FOXType struct {
	SPAN      []*SPAN `xml:"SPAN"`
	P         []*P    `xml:"P"`
	BR        []*BR   `xml:"BR"`
	B         []*B    `xml:"B"`
	I         []*I    `xml:"I"`
	SUB       []*SUB  `xml:"SUB"`
	SUP       []*SUP  `xml:"SUP"`
	InnerText string  `xml:",chardata"`
	Choices   []XMLChoice
	*XMLChoices
}

type XMLChoice struct {
	XMLName xml.Name
	Value   interface{} `xml:",innerxml"`
}
