package xml_handlers

import (
	"encoding/json"
	"github.com/gernhan/xml/entities/views"
	"github.com/gernhan/xml/models"
	"github.com/gernhan/xml/xml/schema"
	"github.com/shopspring/decimal"
	"strings"
)

type PrepaidInvoiceHandler struct{}

var PrepaidHandler PrepaidInvoiceHandler

func InitPrepaidHandler() {
	PrepaidHandler = PrepaidInvoiceHandler{}
}

func (p PrepaidInvoiceHandler) PrepareXmlData(bills []views.VExportingBillsV3) (string, error) {
	brandIdsAndExternalIds := make(map[int64]models.BrandAndExternalIds)
	invoices := schema.Invoices{Invoice: make([]*schema.Invoice, 0)}

	for _, bill := range bills {
		brandIdsAndExternalIds[bill.BrandID] = models.BrandAndExternalIds{BrandID: bill.BrandID, ExternalID: bill.AccountBrand}

		var otherData models.PrepaidInvoiceOtherData
		if err := json.Unmarshal([]byte(bill.OtherData), &otherData); err != nil {
			return "", err
		}

		var taxes []models.TaxesDto
		if err := json.Unmarshal([]byte(otherData.Taxes), &taxes); err != nil {
			return "", err
		}

		var summaries []models.PrepaidBillSummaryDto
		if err := json.Unmarshal([]byte(otherData.BillSummary), &summaries); err != nil {
			return "", err
		}

		var iCdrs []models.PrepaidItemizedBillCdr
		if len(otherData.ItemizedBillCDRs) == 0 {
			otherData.ItemizedBillCDRs = "[]"
		}
		if err := json.Unmarshal([]byte(otherData.ItemizedBillCDRs), &iCdrs); err != nil {
			return "", err
		}

		var iCharges []models.PrepaidItemizedBillCharge
		if err := json.Unmarshal([]byte(otherData.ItemizedBillCharges), &iCharges); err != nil {
			return "", err
		}

		var iTopUps []models.PrepaidItemizedBillTopUp
		if err := json.Unmarshal([]byte(otherData.ItemizedBillTopUps), &iTopUps); err != nil {
			return "", err
		}

		var thirdParties []models.PrepaidThirdParty
		if err := json.Unmarshal([]byte(otherData.ThirdParties), &thirdParties); err != nil {
			return "", err
		}

		prepaidItemizedBills := make([]models.ItemizedBillDto, 0)
		for _, obj := range iCdrs {
			prepaidItemizedBills = append(prepaidItemizedBills, *models.NewItemizedBillDtoFromPrepaidItemizedBillCdr(obj))
		}
		for _, obj := range iCharges {
			prepaidItemizedBills = append(prepaidItemizedBills, *models.NewItemizedBillDtoFromPrepaidItemizedBillCharge(obj))
		}
		for _, obj := range iTopUps {
			prepaidItemizedBills = append(prepaidItemizedBills, *models.NewItemizedBillDtoFromPrepaidItemizedBillTopUp(obj))
		}

		ib := &InvoiceBuilder{}
		ib.header(bill, otherData).
			customer(createInvoiceCustomer(bill)).
			taxes(taxes).
			subscriptions(summaries, prepaidItemizedBills).
			texts(createTexts()).
			thirdParties(*p.createThirdParties(thirdParties))

		invoices.Invoice = append(invoices.Invoice, ib.build())
	}

	body, err := convertToXML(invoices)
	if err != nil {
		return "", err
	}
	newBody := doInvoiceSpecificReplacements(body, "Invoices")
	header := generateHeader(brandIdsAndExternalIds)
	footer := generateFooter()

	sb := strings.Builder{}
	sb.WriteString(header)
	sb.WriteString(newBody)
	sb.WriteString(footer)
	return sb.String(), nil
}

// Implement the createThirdParties function
func (p PrepaidInvoiceHandler) createThirdParties(prepaidBillThirdParties []models.PrepaidThirdParty) *schema.ThirdParties {
	// Group prepaidBillThirdParties by ProviderId
	groupedThirdParties := make(map[string][]models.PrepaidThirdParty)
	for _, prepaidThirdParty := range prepaidBillThirdParties {
		groupedThirdParties[prepaidThirdParty.ProviderID] = append(groupedThirdParties[prepaidThirdParty.ProviderID], prepaidThirdParty)
	}

	thirdParties := &schema.ThirdParties{}
	for _, thirdPartyEntries := range groupedThirdParties {
		thirdParty := schema.ThirdParty{}
		if len(thirdPartyEntries) > 0 {
			providerClass := thirdPartyEntries[0]

			thirdParty.Name = providerClass.ProviderName
			contactDataBuilder := strings.Builder{}
			contactDataBuilder.WriteString(providerClass.ProviderName)
			contactDataBuilder.WriteString(" ")
			contactDataBuilder.WriteString(providerClass.StreetName)
			contactDataBuilder.WriteString(" ")
			contactDataBuilder.WriteString(providerClass.StreetNumber)
			contactDataBuilder.WriteString(", ")
			contactDataBuilder.WriteString(providerClass.ZipCode)
			contactDataBuilder.WriteString(" ")
			contactDataBuilder.WriteString(providerClass.City)
			contactDataBuilder.WriteString(", ")
			contactDataBuilder.WriteString(providerClass.Country)
			contactDataBuilder.WriteString(", ")
			contactDataBuilder.WriteString(providerClass.PhoneNumber)
			contactDataBuilder.WriteString(", ")
			contactDataBuilder.WriteString(providerClass.EmailAddress)
			thirdParty.ContactData = contactDataBuilder.String()
		}
		thirdParty.Detail = ""

		totalGrossAmount := decimal.NewFromFloat(0)
		for _, services := range thirdPartyEntries {
			totalGrossAmount = totalGrossAmount.Add(services.Sum)

			thirdPartyService := schema.ThirdPartyService{
				Name:        services.InvoiceText,
				Detail:      services.PartnerText,
				GrossAmount: amountString(&services.Sum),
			}
			thirdParty.ThirdPartyService = append(thirdParty.ThirdPartyService, &thirdPartyService)
		}
		thirdParty.TotalGrossAmount = amountString(&totalGrossAmount)

		thirdParties.ThirdParty = append(thirdParties.ThirdParty, &thirdParty)
	}

	return thirdParties
}
