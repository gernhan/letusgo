package schema

import "encoding/xml"

func (n XMLChoices) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return MarshalXMLChoicesOnly(e, start, n.XMLChoices)
}

func MarshalXMLChoicesOnly(e *xml.Encoder, start xml.StartElement, elements []interface{}) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	err := MarshalChoicesInOrder(e, elements)
	if err != nil {
		return err
	}

	if err = e.EncodeToken(start.End()); err != nil {
		return err
	}
	return nil
}

func MarshalChoicesInOrder(e *xml.Encoder, elements []interface{}) error {
	//log.Println("Custom marshaller")
	for _, elem := range elements {
		switch v := elem.(type) {
		case string:
			if err := e.EncodeToken(xml.CharData(v)); err != nil {
				return err
			}
		case xml.StartElement:
			if len(v.Name.Local) > 0 {
				err := e.EncodeToken(v)
				if err != nil {
					return err
				}
				if len(v.Attr) == 0 {
					err = e.EncodeToken(xml.EndElement{Name: v.Name})
					if err != nil {
						return err
					}
				}
			}
		default:
			err := e.Encode(v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
