package xml_handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	views "github.com/gernhan/xml/entities/views"
	"github.com/gernhan/xml/models"
	"github.com/gernhan/xml/utils"
	schema "github.com/gernhan/xml/xml/schema"
)

type NormalInvoiceHandler struct{}

var NormalHandler NormalInvoiceHandler

func InitNormalHandler() {
	NormalHandler = NormalInvoiceHandler{}
}

type PreparedData struct {
	VBillHeader               *models.NormalInvoiceOtherDataHeader
	VTaxes                    *models.TaxesDto
	VBillSummary              *models.NormalInvoiceBillSummaryDto
	VBillNoneUsageServiceList []*models.NormalInvoiceNonUsageDto
	VDiscountsPerServiceList  []*models.NormalInvoiceDiscountDto
}

func (handler NormalInvoiceHandler) getPreparedData(brandIdsAndExternalIds map[int64]models.BrandAndExternalIds, bill views.VExportingBillsV3) (*PreparedData, error) {
	vBillHeader := models.NormalInvoiceOtherDataHeader{}
	bill.Copy(&vBillHeader)

	brandIdsAndExternalIds[bill.BrandID] = models.BrandAndExternalIds{
		BrandID:    bill.BrandID,
		ExternalID: bill.AccountBrand,
	}

	var otherData *models.NormalInvoiceOtherData
	if err := utils.FromJSON(bill.OtherData, otherData); err != nil {
		return nil, err
	}
	otherData.Copy(&vBillHeader)

	var vTaxes models.TaxesDto
	if err := utils.FromJSON(otherData.Taxes, &vTaxes); err != nil {
		return nil, err
	}

	var vBillSummaryV3s []*models.NormalInvoiceBillSummaryDto
	if err := utils.FromJSON(otherData.Summary, &vBillSummaryV3s); err != nil {
		return nil, err
	}

	var vBillSummary *models.NormalInvoiceBillSummaryDto
	if len(vBillSummaryV3s) > 0 {
		vBillSummary = vBillSummaryV3s[0]
	}

	var vBillNoneUsageServiceList []*models.NormalInvoiceNonUsageDto
	if err := utils.FromJSON(otherData.NonUsage, &vBillNoneUsageServiceList); err != nil {
		return nil, err
	}

	var vDiscountsPerServiceList []*models.NormalInvoiceDiscountDto
	if err := utils.FromJSON(otherData.Discount, &vDiscountsPerServiceList); err != nil {
		return nil, err
	}

	if otherData.StmThings != "" {
		var obj *models.NormalInvoiceStmThings
		if err := utils.FromJSON(otherData.StmThings, obj); err == nil {
			obj.Copy(&vBillHeader)
		}
	}

	if otherData.X400Things != "" {
		var obj *models.NormalInvoiceX400Things
		if err := utils.FromJSON(otherData.X400Things, obj); err == nil {
			obj.Copy(&vBillHeader)
		}
	}

	if otherData.XrechThings != "" {
		var obj *models.NormalInvoiceXrechThings
		if err := utils.FromJSON(otherData.XrechThings, obj); err == nil {
			obj.Copy(&vBillHeader)
		}
	}

	if otherData.MorePersonalInfo != "" {
		var obj *models.NormalInvoiceMorePersonalInfo
		if err := utils.FromJSON(otherData.MorePersonalInfo, obj); err == nil {
			obj.Copy(&vBillHeader)
		}
	}

	if otherData.MoreAccountThings != "" {
		var obj models.NormalInvoiceMoreAccountThings
		if err := json.Unmarshal([]byte(otherData.MoreAccountThings), &obj); err == nil {
			obj.Copy(&vBillHeader)
		}
	}

	return &PreparedData{
		VBillHeader:               &vBillHeader,
		VTaxes:                    &vTaxes,
		VBillSummary:              vBillSummary,
		VBillNoneUsageServiceList: vBillNoneUsageServiceList,
		VDiscountsPerServiceList:  vDiscountsPerServiceList,
	}, nil
}

func (handler NormalInvoiceHandler) PrepareXMLData(bills []views.VExportingBillsV3, proforma bool) (string, error) {
	// TODO
	//counter := 0
	var exception error
	brandIdsAndExternalIds := make(map[int64]models.BrandAndExternalIds)

	var psi schema.PSI
	for _, bill := range bills {
		preparedData, err := handler.getPreparedData(brandIdsAndExternalIds, bill)
		if err != nil || preparedData == nil {
			if err != nil {
				return "", fmt.Errorf("error creating xml for bill %v: %v", bill.ID, err)
			}
			continue
		}

		// TODO
		//if counter == 0 {
		//	psi = schema.PSI{
		//		O2BANK: O2Bank,
		//		FORMAT: {getPSIFormat(preparedData.VBillHeader)},
		//	}
		//}
		//
		//psi.RECHNUNG = append(psi.RECHNUNG, getPsiRechnung(preparedData.VBillHeader, preparedData.VTaxes,
		//	preparedData.VBillNoneUsageServiceList, preparedData.VDiscountsPerServiceList, preparedData.VBillSummary, proforma))
		//counter++
	}

	if exception != nil {
		return "", fmt.Errorf("xml exporting exception: %v", exception)
	}

	header := generateHeader(brandIdsAndExternalIds)
	xmlBytes, err := xml.MarshalIndent(psi, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling PSI: %v", err)
	}

	result := fmt.Sprintf("%s\n%s\n%s", header, xml.Header, string(xmlBytes))
	return result, nil
}
