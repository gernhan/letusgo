package xml_handlers

import (
	"encoding/xml"
	"github.com/gernhan/xml/models"
	"github.com/gernhan/xml/xml/schema"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

func getPsiRechnung(
	vBillHeader models.NormalInvoiceOtherDataHeader,
	vTaxes models.TaxesDto,
	vBillNoneUsageServiceList []models.NormalInvoiceNonUsageDto,
	vBillDiscountsPerServiceList []models.NormalInvoiceDiscountDto,
	vBillSummary models.NormalInvoiceBillSummaryDto,
	isProforma bool,
) schema.RECHNUNG {
	psiRechnung := schema.RECHNUNG{Value: make([]interface{}, 0)}

	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungAdrZus(vBillHeader))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungArchive(vBillHeader))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungInfoRefs())
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungKtoBel(vBillDiscountsPerServiceList))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungKunZus(vBillHeader))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungMwsZus(vTaxes))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungRecZus(vBillHeader, vTaxes, isProforma))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungSiBel(vBillHeader, vBillSummary))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungDwBel(vBillHeader, vBillNoneUsageServiceList))
	psiRechnung.Value = append(psiRechnung.Value, setupPSIRechnungZahZus(vBillHeader))

	return psiRechnung
}

func setupPSIRechnungZahZusInfoRef(id, pos int) *schema.INFOType {
	infoType := schema.INFOType{
		IDAttr:  id,
		POSAttr: pos,
	}
	return &infoType
}

func setupPSIRechnungZahZus(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.ZAHZUS {
	zahzus := schema.ZAHZUS{
		XMLName: xml.Name{},
		Choices: make([]schema.XMLChoice, 0),
	}

	if vBillHeader.PaymentMode == 1 {
		zahzus.Choices = append(zahzus.Choices, schema.XMLChoice{
			XMLName: xml.Name{Local: "INFOREF"},
			Value:   setupPSIRechnungZahZusInfoRef(1017, 1),
		})

		if vBillHeader.IBAN != "" {
			zahzus.Choices = append(zahzus.Choices, schema.XMLChoice{
				XMLName: xml.Name{Local: "CUST_IBAN"},
				Value:   vBillHeader.IBAN,
			})
		}

		if !vBillHeader.InvoiceDueDate.IsZero() {
			zahzus.Choices = append(zahzus.Choices, schema.XMLChoice{
				XMLName: xml.Name{Local: "ZAH_ZIE"},
				Value:   dateToString(vBillHeader.InvoiceDate),
			})
		}

		if vBillHeader.MandateID != "" {
			zahzus.Choices = append(zahzus.Choices, schema.XMLChoice{
				XMLName: xml.Name{Local: "MANDATE_ID"},
				Value:   vBillHeader.MandateID,
			})

		}
	} else {
		zahzus.Choices = append(zahzus.Choices, schema.XMLChoice{
			XMLName: xml.Name{Local: "INFOREF"},
			Value:   setupPSIRechnungZahZusInfoRef(3, 1),
		})

		if !vBillHeader.InvoiceDueDate.IsZero() {
			zahzus.Choices = append(zahzus.Choices, schema.XMLChoice{
				XMLName: xml.Name{Local: "ZIE_DAT"},
				Value:   dateToString(vBillHeader.InvoiceDueDate),
			})
		}
	}

	return &zahzus
}

func setupPSIRechnungSiBel(vBillHeader models.NormalInvoiceOtherDataHeader, vBillSummary models.NormalInvoiceBillSummaryDto) *schema.SIBEL {
	sibel := schema.SIBEL{
		Choices: make([]schema.XMLChoice, 0),
	}
	sibel.Choices = append(sibel.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "CONT_DAT_BEL"},
		Value:   setupPSIRechnungSibelContBelContData(vBillHeader, vBillSummary),
	})
	return &sibel
}

func setupPSIRechnungSibelContBelContData(vBillHeader models.NormalInvoiceOtherDataHeader, vBillSummary models.NormalInvoiceBillSummaryDto) *schema.CONTDATA {
	startDateBilling := vBillSummary.StartDtBilling
	var contractEndDate *time.Time
	if !vBillSummary.ContractEndDt.Equal(time.Time{}) {
		contractEndDate = &vBillSummary.ContractEndDt
	}
	var endDtRequested *time.Time
	if !vBillSummary.EndDtRequested.Equal(time.Time{}) {
		endDtRequested = &vBillSummary.EndDtRequested
	}

	contData := schema.CONTDATA{}
	contData.TRAV = append(contData.TRAV, setupPSIRechnungSibelContBelContDataTrav(1, vBillHeader.CustomerID))
	contData.TRAV = append(contData.TRAV, setupPSIRechnungSibelContBelContDataTrav(2, vBillSummary.ServiceName))
	contData.TRAV = append(contData.TRAV, setupPSIRechnungSibelContBelContDataTrav(3, dateToString(startDateBilling)))
	if contractEndDate != nil {
		contData.TRAV = append(contData.TRAV, setupPSIRechnungSibelContBelContDataTrav(5, dateToString(*contractEndDate)))
	}
	if endDtRequested != nil {
		contData.TRAV = append(contData.TRAV, setupPSIRechnungSibelContBelContDataTrav(6, dateToString(*endDtRequested)))
	}

	return &contData
}

func setupPSIRechnungSibelContBelContDataTrav(id int, text string) *schema.TRAV {
	trav := schema.TRAV{}
	trav.IDAttr = id
	trav.Value = text

	return &trav
}

func setupPSIRechnungRecZusRDat(vBillHeader models.NormalInvoiceOtherDataHeader) string {
	invoiceDate := vBillHeader.InvoiceDate // Assume it's already a time.Time
	sdf := "02-01-2006"                    // Format for "DD-MM-YYYY"
	return invoiceDate.Format(sdf)
}

func setupPSIRechnungRecZusEDat(vBillHeader models.NormalInvoiceOtherDataHeader) string {
	invoiceDate := vBillHeader.InvoiceDate // Assume it's already a time.Time
	sdf := "02-01-2006"                    // Format for "DD-MM-YYYY"
	return invoiceDate.Format(sdf)
}

func setupPSIRechnungRecZusAktuell(vTaxes models.TaxesDto) *schema.AKTUELL {
	aktuell := schema.AKTUELL{}

	aktuell.Choices = append(aktuell.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "AKT_GES_TPC"},
		Value:   "1",
	})

	aktuell.Choices = append(aktuell.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "AKT_GES_TPC"},
		Value:   vTaxes.Sum,
	})

	aktuell.Choices = append(aktuell.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "AKT_GES_TPC"},
		Value:   vTaxes.SumNet,
	})

	return &aktuell
}

func setupPSIRechnungRecZus(vBillHeader models.NormalInvoiceOtherDataHeader, vTaxes models.TaxesDto, isProforma bool) *schema.RECZUS {
	reczus := schema.RECZUS{}

	aktuell := setupPSIRechnungRecZusAktuell(vTaxes)
	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "AKTUELL"},
		Value:   aktuell,
	})

	rdat := setupPSIRechnungRecZusRDat(vBillHeader)
	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_DAT"},
		Value:   rdat,
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_NUM_RST"},
		Value:   vBillHeader.InvoiceID,
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_UEB_Z"},
		Value:   "Rechnung",
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "E_DAT"},
		Value:   setupPSIRechnungRecZusEDat(vBillHeader),
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "AKT_GES_TPC"},
		Value:   "",
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "LZ"},
		Value:   setupPSIRechnungRecZusLZ(vBillHeader),
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_VON"},
		Value:   dateToString(vBillHeader.RechnungszeitraumVon),
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_BIS"},
		Value:   dateToString(vBillHeader.RechnungszeitraumBis),
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_GES"},
		Value:   vBillHeader.Amount,
	})

	reczus.Choices = append(reczus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_WAE"},
		Value:   vBillHeader.Currency,
	})

	if isProforma {
		reczus.Choices = append(reczus.Choices, schema.XMLChoice{
			XMLName: xml.Name{Local: "R_TYP"},
			Value:   "PROFORMA",
		})
	}

	return &reczus
}

func setupPSIRechnungRecZusLZ(vBillHeader models.NormalInvoiceOtherDataHeader) string {
	billPeriod := dateToString(vBillHeader.RechnungszeitraumVon) + " - " + dateToString(vBillHeader.RechnungszeitraumBis)
	return billPeriod
}

func setupPSIRechnungDwBel(vBillHeader models.NormalInvoiceOtherDataHeader, vBillNoneUsageServiceList []models.NormalInvoiceNonUsageDto) *schema.DWBEL {
	dwbel := schema.DWBEL{}

	dwbel.Choices = append(dwbel.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "GES_DW_BEL_N"},
		Value:   vBillHeader.AmountNet,
	})

	dw := setupPSIRechnungDwBelDw(vBillNoneUsageServiceList)
	dwbel.Choices = append(dwbel.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "GES_DW_BEL_N"},
		Value:   dw,
	})

	return &dwbel
}

func setupPSIRechnungDwBelDwGgbZus(vBillNoneUsageServiceList []models.NormalInvoiceNonUsageDto) *schema.GGBZUSType {
	ggbZus := schema.GGBZUSType{}
	sumGesB := decimal.NewFromFloat(0)
	sumGesN := decimal.NewFromFloat(0)

	rcList := make([]models.NormalInvoiceNonUsageDto, 0)

	for _, vBillNoneUsageService := range vBillNoneUsageServiceList {
		if vBillNoneUsageService.ChargeMode == 2 {
			rcList = append(rcList, vBillNoneUsageService)
		}
	}

	for _, vBillNoneUsageService := range rcList {
		amountNet := vBillNoneUsageService.AmountNet
		discountedAmountNet := vBillNoneUsageService.DiscountedAmountNet
		sumGesB = sumGesB.Add(amountNet.Decimal)
		sumGesN = sumGesN.Add(discountedAmountNet.Decimal)
		ggbZus.Choices = append(ggbZus.Choices, schema.XMLChoice{
			XMLName: xml.Name{Local: "GGB"},
			Value:   setupPSIRechnungDwBelDwGgbZusGgb(vBillNoneUsageService),
		})
	}

	ggbZus.Choices = append(ggbZus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "GES_B"},
		Value:   sumGesB.String(),
	})
	ggbZus.Choices = append(ggbZus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "GES_N"},
		Value:   sumGesN.String(),
	})

	return &ggbZus
}

func setupPSIRechnungDwBelDw(vBillNoneUsageServiceList []models.NormalInvoiceNonUsageDto) *schema.DW {
	dw := schema.DW{}

	ggbZus := setupPSIRechnungDwBelDwGgbZus(vBillNoneUsageServiceList)
	dw.Choices = append(dw.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "GGB_ZUS"},
		Value:   ggbZus,
	})

	zslZus := setupPSIRechnungDwBelDwZslZus(vBillNoneUsageServiceList)
	dw.Choices = append(dw.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "ZSL_ZUS"},
		Value:   zslZus,
	})

	return &dw
}

func setupPSIRechnungDwBelDwZslZus(vBillNoneUsageServiceList []models.NormalInvoiceNonUsageDto) *schema.ZSLZUSType {
	zslZus := schema.ZSLZUSType{}
	sumGesB := decimal.NewFromInt(0)
	sumGesN := decimal.NewFromInt(0)
	otcList := make([]models.NormalInvoiceNonUsageDto, 0)

	for _, vBillNoneUsageService := range vBillNoneUsageServiceList {
		if vBillNoneUsageService.ChargeMode == 1 {
			otcList = append(otcList, vBillNoneUsageService)
		}
	}

	for _, vBillNoneUsageService := range otcList {
		amountNet := vBillNoneUsageService.AmountNet
		discountedAmountNet := vBillNoneUsageService.DiscountedAmountNet

		sumGesB = sumGesB.Add(amountNet.Decimal)
		sumGesN = sumGesN.Add(discountedAmountNet.Decimal)

		zslZus.Choices = append(zslZus.Choices, schema.XMLChoice{
			XMLName: xml.Name{Local: "ZSL"},
			Value:   setupPSIRechnungDwBelDwZslZusZsl(vBillNoneUsageService),
		})
	}

	zslZus.Choices = append(zslZus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "GES_B"},
		Value:   sumGesB.String(),
	})
	zslZus.Choices = append(zslZus.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "GES_N"},
		Value:   sumGesN.String(),
	})

	return &zslZus
}

func setupPSIRechnungDwBelDwZslZusZsl(vBillNoneUsageService models.NormalInvoiceNonUsageDto) *schema.RBelType {
	rbelType := schema.RBelType{}
	infoRNam := schema.INFOType{}

	displayValue := vBillNoneUsageService.OptDisplayValue
	if displayValue == "" {
		displayValue = vBillNoneUsageService.DisplayValue
	}
	infoRNam.Choices = make([]schema.XMLChoice, 0)
	infoRNam.Choices = append(infoRNam.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "FOX"},
		Value:   displayValue,
	})

	rbelType.Choices = append(rbelType.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_NAM"},
		Value:   infoRNam,
	})
	rbelType.Choices = append(rbelType.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "BET_N"},
		Value:   vBillNoneUsageService.AmountNet,
	})
	rbelType.Choices = append(rbelType.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "BET_N"},
		Value:   vBillNoneUsageService.DiscountedAmountNet,
	})

	return &rbelType
}

func setupPSIRechnungDwBelDwGgbZusGgb(vBillNoneUsageService models.NormalInvoiceNonUsageDto) *schema.RBelType {
	rbelType := schema.RBelType{}
	infoRNam := schema.INFOType{}

	displayValue := vBillNoneUsageService.OptDisplayValue
	if displayValue == "" {
		displayValue = vBillNoneUsageService.DisplayValue
	}

	infoRNam.Choices = make([]schema.XMLChoice, 0)
	infoRNam.Choices = append(infoRNam.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "SUP"},
		Value:   displayValue,
	})

	rbelType.Choices = append(rbelType.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_NAM"},
		Value:   infoRNam,
	})

	rbelType.Choices = append(rbelType.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "BET_B"},
		Value:   vBillNoneUsageService.AmountNet,
	})

	rbelType.Choices = append(rbelType.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "BET_N"},
		Value:   vBillNoneUsageService.DiscountedAmountNet,
	})

	return &rbelType
}

func dateToString(date time.Time) string {
	return date.Format("02-01-2006")
}

func setupPSIRechnungMwsZus(vTaxes models.TaxesDto) *schema.MWSZUS {
	return &schema.MWSZUS{KOPF: []*schema.KOPF{{
		INFOType: &schema.INFOType{
			IDAttr:  0,
			POSAttr: 0,
			FOXType: &schema.FOXType{Choices: []schema.XMLChoice{
				{
					XMLName: xml.Name{Local: "KOPF"},
					Value:   setupPSIRechnungMwsZusAktuell(vTaxes),
				},
			}},
		},
	}}}
}

func setupPSIRechnungMwsZusAktuell(vTaxes models.TaxesDto) *schema.MWS {
	aktuellMWS := schema.MWS{
		Choices: make([]schema.XMLChoice, 0),
	}
	aktuellMWS.Choices = append(aktuellMWS.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "BET"},
		Value:   vTaxes.Tax,
	})
	aktuellMWS.Choices = append(aktuellMWS.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "BET"},
		Value:   vTaxes.SumNet,
	})
	aktuellMWS.Choices = append(aktuellMWS.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "BET"},
		Value:   vTaxes.TaxRate,
	})
	return &aktuellMWS
}

func setupPSIRechnungKunZus(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.KUNZUS {
	obj := schema.KUNZUS{}

	obj.Choices = append(obj.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "KNUM"},
		Value:   vBillHeader.CustomerID,
	})

	if vBillHeader.FrameContractID != "" {
		obj.Choices = append(obj.Choices, schema.XMLChoice{
			XMLName: xml.Name{Local: "RV_NUM"},
			Value:   vBillHeader.FrameContractID,
		})
	}

	obj.Choices = append(obj.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "F_DAT"},
		Value:   vBillHeader.InvoiceDueDate,
	})

	obj.Choices = append(obj.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "UST_NUM"},
		Value:   vBillHeader.VatIDNumber,
	})

	obj.Choices = append(obj.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "A_NUM"},
		Value:   vBillHeader.AccountID,
	})

	obj.Choices = append(obj.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "K_ZST"},
		Value:   vBillHeader.Additional1,
	})

	obj.Choices = append(obj.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "K_ZST"},
		Value:   vBillHeader.Additional2,
	})

	return &obj
}

func setupPSIRechnungKtoBel(vBillDiscountsPerServiceList []models.NormalInvoiceDiscountDto) *schema.KTOBEL {
	ktoBel := schema.KTOBEL{}
	ktoBel.KTO = append(ktoBel.KTO, setupPSIRechnungKtoBelKto(vBillDiscountsPerServiceList))
	return &ktoBel
}

func setupPSIRechnungKtoBelKto(vBillDiscountsPerServiceList []models.NormalInvoiceDiscountDto) *schema.KTO {
	kto := schema.KTO{Choices: make([]schema.XMLChoice, 0)}
	kto.Choices = append(kto.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "RAB_ZUS"},
		Value:   setupPSIRechnungKtoBelKtoRabZus(vBillDiscountsPerServiceList),
	})
	return &kto
}

func setupPSIRechnungKtoBelKtoRabZus(vBillDiscountsPerServiceList []models.NormalInvoiceDiscountDto) *schema.RABZUSType {
	rabzus := schema.RABZUSType{Choices: make([]schema.XMLChoice, 0)}

	for _, vBillDiscountsPerService := range vBillDiscountsPerServiceList {
		rabzus.Choices = append(rabzus.Choices, schema.XMLChoice{
			XMLName: xml.Name{Local: "RAB"},
			Value: setupPSIRechnungKtoBelKtoRabZusRab(
				vBillDiscountsPerService.DiscountDisplayValue,
				vBillDiscountsPerService.DiscountPercentage.Decimal.InexactFloat64(),
			),
		})
	}

	return &rabzus
}

func setupPSIRechnungKtoBelKtoRabZusRab(rnam string, percentage float64) *schema.RBelType {
	rab := schema.RBelType{}
	value := schema.INFOType{}
	value.Choices = append(value.Choices, schema.XMLChoice{
		XMLName: xml.Name{Local: "R_NAM"},
		Value:   rnam,
	})

	rab.Choices = append(rab.Choices, schema.XMLChoice{
		XMLName: xml.Name{
			Local: "PER",
		},
		Value: percentage,
	})

	rab.Choices = append(rab.Choices, schema.XMLChoice{
		XMLName: xml.Name{
			Local: "R_NAM",
		},
		Value: value,
	})
	return &rab
}

// setupPSIRechnungAdrZus creates and returns an ADRZUS2 object.
func setupPSIRechnungAdrZus(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.ADRZUS {
	adrzus := schema.ADRZUS{VADROrRADROrCDADR: make([]interface{}, 0)}
	adrzus.VADROrRADROrCDADR = append(adrzus.VADROrRADROrCDADR, setupPSIRechnungAdrZusRadr(vBillHeader))

	switch vBillHeader.Billmedia {
	case Online.ID:
		adrzus.VADROrRADROrCDADR = append(adrzus.VADROrRADROrCDADR, setupPSIRechnungAdrZusEadr(vBillHeader))
	case XRechnung.ID:
		adrzus.VADROrRADROrCDADR = append(adrzus.VADROrRADROrCDADR, setupPSIRechnungAdrZusXRECHadr(vBillHeader))
	case EDIFACT95APAPIER.ID:
		adrzus.VADROrRADROrCDADR = append(adrzus.VADROrRADROrCDADR, setupPSIRechnungAdrZusX400adr("95A", vBillHeader))
	case EDIFACT95BPAPIER.ID:
		adrzus.VADROrRADROrCDADR = append(adrzus.VADROrRADROrCDADR, setupPSIRechnungAdrZusX400adr("95B", vBillHeader))
	case EDIFACT99B.ID:
		adrzus.VADROrRADROrCDADR = append(adrzus.VADROrRADROrCDADR, setupPSIRechnungAdrZusX400adr("99B", vBillHeader))
	}

	if vBillHeader.CustID != vBillHeader.ServiceCustomer {
		adrzus.VADROrRADROrCDADR = append(adrzus.VADROrRADROrCDADR, setupPSIRechnungAdrZusVadr(vBillHeader))
	}

	return &adrzus
}

func setupPSIRechnungAdrZusEadr(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.EADR {
	eadr := schema.EADR{}
	eadr.IDAttr = 1
	eadr.EVNAttr = 1

	if vBillHeader.Contact != "" {
		eadr.ADR = vBillHeader.Contact
	}

	return &eadr
}

// setupPSIRechnungAdrZusRadr function sets up the RADR object based on vBillHeader values
func setupPSIRechnungAdrZusRadr(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.RADR {

	radr := schema.RADR{}
	radr.IDAttr = 1

	if vBillHeader.Street != "" && vBillHeader.HouseNumber != "" {
		radr.ADR1 = vBillHeader.Street + " " + vBillHeader.HouseNumber
	}
	if vBillHeader.Firma != "" {
		radr.FIR = vBillHeader.Firma
	}
	if vBillHeader.ZipCode != "" {
		radr.PLZ = vBillHeader.ZipCode
	}
	if vBillHeader.City != "" {
		radr.STA = vBillHeader.City
	}
	if vBillHeader.FirstName != "" {
		radr.VNAM = vBillHeader.FirstName
	}
	if vBillHeader.LastName != "" {
		radr.NNAM = vBillHeader.LastName
	}
	if vBillHeader.Country != "" {
		radr.LAN = vBillHeader.Country
	}
	if vBillHeader.Salutation != "" {
		radr.ANR = vBillHeader.Salutation
	}
	if vBillHeader.Department != "" {
		radr.ABT = vBillHeader.Department
	}
	if vBillHeader.Branch != "" {
		radr.NIE = vBillHeader.Branch
	}

	return &radr
}

// setupPSIRechnungAdrZusVadr function sets up the VADR object based on vBillHeader values
func setupPSIRechnungAdrZusVadr(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.VADR {
	vadr := schema.VADR{}
	vadr.IDAttr = 1

	if vBillHeader.AccountAddress != "" {
		vadr.ADR1 = vBillHeader.AccountAddress
	}
	if vBillHeader.AccountCompanyName != "" {
		vadr.FIR = vBillHeader.AccountCompanyName
	}
	if vBillHeader.AccountZip != "" {
		vadr.PLZ = vBillHeader.AccountZip
	}
	if vBillHeader.AccountCity != "" {
		vadr.STA = vBillHeader.AccountCity
	}
	if vBillHeader.AccountFirstName != "" {
		vadr.VNAM = vBillHeader.AccountFirstName
	}
	if vBillHeader.AccountLastName != "" {
		vadr.NNAM = vBillHeader.AccountLastName
	}
	if vBillHeader.AccountCountry != "" {
		vadr.LAN = vBillHeader.AccountCountry
	}
	if vBillHeader.AccountSalutation != "" {
		vadr.ANR = vBillHeader.AccountSalutation
	}
	if vBillHeader.AccountDepartment != "" {
		vadr.ABT = vBillHeader.AccountDepartment
	}
	if vBillHeader.AccountBranch != "" {
		vadr.NIE = vBillHeader.AccountBranch
	}

	return &vadr
}

func setupPSIRechnungAdrZusX400adr(format string, vBillHeader models.NormalInvoiceOtherDataHeader) *schema.X400ADR {
	x400adr := schema.X400ADR{}
	x400adr.EvnAttr = 1
	x400adr.IDAttr = 1
	x400adr.RecAttr = 1
	x400adr.FormatAttr = format
	x400adr.ADMD = make([]string, 0)
	x400adr.ADMD = append(x400adr.ADMD, vBillHeader.X400Admd)
	x400adr.ANAM = make([]string, 0)
	x400adr.ANAM = append(x400adr.ANAM, vBillHeader.X400Anam)
	x400adr.LAN = make([]string, 0)
	x400adr.LAN = append(x400adr.LAN, vBillHeader.X400Lan)
	x400adr.VNAM = make([]string, 0)
	x400adr.VNAM = append(x400adr.VNAM, vBillHeader.X400Vnam)
	x400adr.NNAM = make([]string, 0)
	x400adr.NNAM = append(x400adr.NNAM, vBillHeader.X400Nnam)
	x400adr.ORG = make([]string, 0)
	x400adr.ORG = append(x400adr.ORG, vBillHeader.X400Org)
	x400adr.OU = make([]string, 0)
	x400adr.OU = append(x400adr.OU, vBillHeader.X400Ou)

	return &x400adr
}

// setupPSIRechnungAdrZusXRECHadr creates and returns an XRECHADR object.
func setupPSIRechnungAdrZusXRECHadr(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.XRECHADR {
	xrechadr := schema.XRECHADR{}
	xrechadr.IDAttr = 1
	xrechadr.XLEITWEG = make([]string, 0)
	xrechadr.XLEITWEG = append(xrechadr.XLEITWEG, vBillHeader.XrechLeitweg)
	xrechadr.XEMAIL = make([]string, 0)
	xrechadr.XEMAIL = append(xrechadr.XLEITWEG, vBillHeader.XrechEmail)
	xrechadr.XAUFTRAGNR = make([]string, 0)
	xrechadr.XAUFTRAGNR = append(xrechadr.XLEITWEG, vBillHeader.XrechAuftragnr)
	xrechadr.XLIEFERANT = make([]string, 0)
	xrechadr.XLIEFERANT = append(xrechadr.XLEITWEG, vBillHeader.XrechLieferant)
	xrechadr.XVERSAND = make([]string, 0)
	xrechadr.XVERSAND = append(xrechadr.XLEITWEG, vBillHeader.XrechVersand)
	return &xrechadr
}

// setupPSIRechnungArchive creates and returns an ARCHIVE object.
func setupPSIRechnungArchive(vBillHeader models.NormalInvoiceOtherDataHeader) *schema.ARCHIVE {
	archive := schema.ARCHIVE{
		BMAttr:  vBillHeader.Billmanager,
		CSVAttr: 0,
	}
	if vBillHeader.HasEvn == 2 {
		archive.CSVAttr = 1
	}

	return &archive
}

// setupPSIRechnungInfoRef creates and returns an INFOREF object.
func setupPSIRechnungInfoRef(id, pos int) *schema.INFOREF {
	info := schema.INFOREF{
		IDAttr:  id,
		POSAttr: pos,
	}
	return &info
}

// setupPSIRechnungInfoRefs creates and returns a slice of INFOREF objects.
func setupPSIRechnungInfoRefs() []*schema.INFOREF {
	infoRefList := []*schema.INFOREF{
		setupPSIRechnungInfoRef(1, 0),
		setupPSIRechnungInfoRef(5, 5),
		setupPSIRechnungInfoRef(9, 3),
		setupPSIRechnungInfoRef(10, 10),
		setupPSIRechnungInfoRef(15, 15),
		setupPSIRechnungInfoRef(16, 2),
		setupPSIRechnungInfoRef(325, 325),
		setupPSIRechnungInfoRef(1017, 1),
		setupPSIRechnungInfoRef(1020, 315),
		setupPSIRechnungInfoRef(1021, 316),
		setupPSIRechnungInfoRef(1022, 317),
		setupPSIRechnungInfoRef(1023, 319),
	}

	return infoRefList
}

func convertXMLEntities(text string) string {
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&nbsp;", " ")
	text = strings.ReplaceAll(text, "&ouml;", "ö")
	text = strings.ReplaceAll(text, "&auml;", "ä")
	text = strings.ReplaceAll(text, "&uuml;", "ü")
	return text
}
