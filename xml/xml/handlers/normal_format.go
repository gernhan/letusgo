package xml_handlers

import (
	"github.com/gernhan/xml/models"
	"github.com/gernhan/xml/xml/schema"
)

func setupPSIFormatSiBelContDatBelTrav(id int, text string) *schema.TRAV {
	trav := schema.TRAV{
		IDAttr: id,
		Value:  text,
	}
	return &trav
}

func setupPSIFormatSiBelContDatBel() *schema.CONTDATBEL {
	contDatBel := schema.CONTDATBEL{
		TRAV: []*schema.TRAV{
			setupPSIFormatSiBelContDatBelTrav(2, "Produkt/Pack"),
			setupPSIFormatSiBelContDatBelTrav(3, "Vertragsbeginn"),
			setupPSIFormatSiBelContDatBelTrav(4, "Ende Mindestvertragslaufzeit"),
			setupPSIFormatSiBelContDatBelTrav(5, "Kündigungsfrist"),
			setupPSIFormatSiBelContDatBelTrav(6, "spätester Kündigungseingang"),
		},
	}
	return &contDatBel
}

func setupPSIFormatSiBel() *schema.SIBEL {
	sibel := schema.SIBEL{
		GESSIBEL: []*schema.FOXType{
			{
				SPAN: []*schema.SPAN{
					{
						InnerText: "Rufnummern / Anschlussbezogene Beträge",
					},
				},
			},
		},
		CONTDATBEL: []*schema.CONTDATBEL{
			setupPSIFormatSiBelContDatBel(),
		},
	}

	return &sibel
}

func setupPSIFormatRecZusAktuellGes() *schema.GES {
	return &schema.GES{XMLChoices: &schema.XMLChoices{
		XMLChoices: []interface{}{
			"Rechnungsbetrag (in EUR exkl. MwSt.) vom ",
			schema.SPAN{
				CharStyles: &schema.CharStyles{
					FIELDNAMEAttr: "R_VON",
				},
			},
			" - ",
			schema.SPAN{
				CharStyles: &schema.CharStyles{
					FIELDNAMEAttr: "R_BIS",
				},
			},
		},
	}}
}

func setupPSIFormatRecZusRGes() *schema.INFOType {
	info := &schema.INFOType{
		FOXType: &schema.FOXType{
			SPAN: []*schema.SPAN{
				{InnerText: "Rechnungsbetrag (in EUR inkl. MwSt)"},
			},
		},
	}
	return info
}

func setupPSIFormatRecZusRDAT() string {
	return "Rechnungsdatum"
}

func setupPSIFormatRecZusRNum() string {
	return "Rechnungsnummer"
}

func setupPSIFormatRecZusAktuellAktGesTPC() *schema.FOXType {
	return &schema.FOXType{XMLChoices: &schema.XMLChoices{
		XMLChoices: []interface{}{
			"Bruttoleistungen anderer Anbieter ",
			schema.SUP{
				InnerText: "A",
			},
			" (inkl. MwSt.)",
		},
	}}
}

func setupPSIFormatRecZusAktuell() *schema.AKTUELL {
	return &schema.AKTUELL{
		AKTGESTPC: []*schema.FOXType{
			setupPSIFormatRecZusAktuellAktGesTPC(),
		},
		AKTGESB: []*schema.FOXType{
			setupPSIFormatRecZusAktuellGes().FOXType,
		},
	}
}

func setupPSIFormatRecZus() *schema.RECZUS {
	return &schema.RECZUS{
		AKTUELL: []*schema.AKTUELL{
			setupPSIFormatRecZusAktuell(),
		},
		RDAT: []string{
			setupPSIFormatRecZusRDAT(),
		},
		RGES: []*schema.INFOType{
			setupPSIFormatRecZusRGes(),
		},
		RNUM: []string{
			setupPSIFormatRecZusRNum(),
		},
		LZ: []string{
			"Leistungszeitraum",
		},
	}
}

func setupPSIFormatMwsZus() *schema.MWSZUS {
	return &schema.MWSZUS{
		NAM: schema.NAM{XMLChoices: &schema.XMLChoices{
			XMLChoices: []interface{}{
				"Mehrwehrtsteuer ",
				schema.SPAN{
					CharStyles: &schema.CharStyles{
						FIELDNAMEAttr: "SAT",
					},
				},
				"% (von ",
				schema.SPAN{
					CharStyles: &schema.CharStyles{
						FIELDNAMEAttr: "R_GES",
					},
				},
				schema.SPAN{
					CharStyles: &schema.CharStyles{
						FIELDNAMEAttr: "R_WAE",
					},
				},
				")",
			},
		}},
	}
}

func setupPSIFormatKunZusKzst(id int, text string) *schema.KZST {
	return &schema.KZST{IDAttr: id, Value: text}
}

func setupPSIFormatKunZus(vBillHeader *models.NormalInvoiceOtherDataHeader) *schema.KUNZUS {
	kunzus := schema.KUNZUS{
		FDAT: []string{
			"Ihre Mobilfunknummer",
		},
		ANUM: []string{
			"Kundennummer",
		},
		KNUM: []string{
			"Ihre Kundenkontonummer",
		},
		MNUM: []string{
			"Fällig am",
		},
		KZST: []*schema.KZST{
			setupPSIFormatKunZusKzst(1, "Zusatz 1"),
			setupPSIFormatKunZusKzst(2, "Zusatz 2"),
		},
		ANZNUM: []string{
			"Anzahl Ihrer Rufnummern",
		},
		USTNUM: []string{
			"Ihre Umsatzsteuer-ID:",
		},
	}

	if vBillHeader.FrameContractID != "" {
		kunzus.RVNUM = []string{
			"Rahmenvertragsnummer",
		}
	}
	return &kunzus
}

func setupPSIFormatDwBelDwRabBet(text string) *schema.FOXType {
	return &schema.FOXType{
		InnerText: text,
	}
}

func setupPSIFormatDwBelDwZlsZusGes(text string) *schema.GES {
	return &schema.GES{
		INFOType: &schema.INFOType{
			FOXType: &schema.FOXType{InnerText: text},
		},
	}
}

func setupPSIFormatDwBelDwZlsZus() *schema.ZUSType {
	return &schema.ZUSType{GES: []*schema.GES{
		setupPSIFormatDwBelDwZlsZusGes("Zusatzleistungen"),
	}}
}

func setupPSIFormatDwBelDwGgbZusGes(text string) *schema.GES {
	return &schema.GES{
		INFOType: &schema.INFOType{
			FOXType: &schema.FOXType{InnerText: text},
		},
	}
}
