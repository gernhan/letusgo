package xml_handlers

import (
	"log"
	"strconv"
	"strings"

	"github.com/gernhan/xml/db"
	"github.com/gernhan/xml/entities/views"
	"github.com/gernhan/xml/models"

	"github.com/gernhan/xml/xml/schema"
)

func createBrandSpecificGenericTexts(textModulesMap map[string]string) *schema.GenericTextsType {
	genericTextsType := schema.GenericTextsType{}

	for _, gtl := range schema.BrandGenericTextsLabelsList {
		genericTextsType.Text = append(genericTextsType.Text, createTextTypeById(gtl, textModulesMap[gtl]))
	}

	return &genericTextsType
}

func createTextTypeById(id, text string) *schema.TextType {
	return &schema.TextType{
		ID:      id,
		Content: text,
	}
}

func createInvoiceAddress(bill views.VExportingBillsV3) *schema.InvoiceAddressType {
	invoiceAddressType := schema.InvoiceAddressType{
		Salutation:      bill.SalutationString,
		Firstname:       bill.FirstName,
		Lastname:        bill.LastName,
		Street:          bill.Street,
		StreetNumber:    bill.HouseNumber,
		PostalCode:      bill.ZIP,
		City:            bill.City,
		Country:         bill.Alpha2,
		AddressAddition: bill.AdditionalAddress,
	}
	return &invoiceAddressType
}

func createEmailToAddress(email string) *schema.EmailToAddressType {
	emailToAddressType := schema.EmailToAddressType{
		NotificationEmail: email,
	}
	return &emailToAddressType
}

func createInvoiceCustomer(bill views.VExportingBillsV3) *schema.CustomerType {
	customerType := schema.CustomerType{
		CustomerNumber: strconv.FormatInt(bill.CustomerID, 10),
		AccountNumber:  bill.InvoiceReference,
	}

	customerType.InvoiceAddress = createInvoiceAddress(bill)
	customerType.EmailToAddress = createEmailToAddress(bill.Email)
	return &customerType
}

func createInvoiceItemsColumnType(textModulesMap map[string]string) *schema.InvoiceItemsColumnType {
	invoiceItemsColumnType := schema.InvoiceItemsColumnType{
		Amount:  textModulesMap["InvoiceItems.Amount"],
		Header:  textModulesMap["InvoiceItems.Header"],
		TaxRate: textModulesMap["InvoiceItems.TaxRate"],
	}
	return &invoiceItemsColumnType
}

func createInvoiceTextsType(textModulesMap map[string]string) *schema.InvoiceTextsType {
	invoiceTextsType := schema.InvoiceTextsType{
		InvoiceAmount:    textModulesMap["Invoice.InvoiceAmount"],
		InvoiceNumber:    textModulesMap["Invoice.InvoiceNumber"],
		InvoiceDate:      textModulesMap["Invoice.InvoiceDate"],
		InvoiceTitle:     textModulesMap["Invoice.InvoiceTitle"],
		InvoiceNetAmount: textModulesMap["Invoice.InvoiceNetAmount"],
		TaxAmount:        textModulesMap["Invoice.TaxAmount"],
		InvoicePeriod:    textModulesMap["Invoice.InvoicePeriod"],
	}
	return &invoiceTextsType
}

func createCustomerTextsType(textModulesMap map[string]string) *schema.CustomerTextsType {
	customerTextsType := schema.CustomerTextsType{
		CustomerNumber: textModulesMap["Customer.CustomerNumber"],
		AccountNumber:  textModulesMap["Customer.AccountNumber"],
	}
	return &customerTextsType
}

func createThirdPartiesTextType(textModulesMap map[string]string) *schema.ThirdPartiesTextType {
	thirdPartiesTextType := schema.ThirdPartiesTextType{
		Header: textModulesMap["ThirdParties.Header"],
		Text:   textModulesMap["ThirdParties.Text"],
	}

	thirdPartyTextType := schema.ThirdPartyTextType{
		ThirdPartyAmountInvoice: textModulesMap["ThirdParty.ThirdPartyAmountInvoice"],
		ThirdPartyAmountEVN:     textModulesMap["ThirdParty.ThirdPartyAmountEVN"],
		Service:                 textModulesMap["ThirdParty.Service"],
		Detail:                  textModulesMap["ThirdParty.Detail"],
		Amount:                  textModulesMap["ThirdParty.Amount"],
		Footnote:                textModulesMap["ThirdParty.Footnote"],
		ThirdPartyDetails:       textModulesMap["ThirdParty.ThirdPartyDetails"],
	}

	thirdPartiesTextType.ThirdParty = &thirdPartyTextType
	return &thirdPartiesTextType
}

func createBrandInvoiceTextsTypeCommon(brandID string, textModulesMap map[string]string) *schema.BrandInvoiceTextsType {
	brandInvoiceTextsType := schema.BrandInvoiceTextsType{}
	brandInvoiceTextsType.BrandId = brandID

	email := schema.EmailType{}
	email.From = &schema.MarkupTextType{Content: textModulesMap["Email.From"]}
	email.ReplyTo = &schema.MarkupTextType{Content: textModulesMap["Email.replyTo"]}
	email.Subject = &schema.MarkupTextType{Content: textModulesMap["Email.Subject"]}

	brandInvoiceTextsType.Email = &email

	footer := schema.MarkupTextType{}
	footer.Content = textModulesMap["ImmediateBillFooter"]

	brandInvoiceTextsType.Footer = &footer

	sender := schema.MarkupTextType{}
	sender.Content = textModulesMap["ImmediateBillSender"]
	brandInvoiceTextsType.Sender = &sender

	markupTextType := schema.MarkupTextType{}
	markupTextType.Content = textModulesMap["TaxedAmount"]
	brandInvoiceTextsType.TaxedAmount = &markupTextType

	brandInvoiceTextsType.Invoice = createInvoiceTextsType(textModulesMap)
	brandInvoiceTextsType.Customer = createCustomerTextsType(textModulesMap)
	brandInvoiceTextsType.InvoiceItems = createInvoiceItemsColumnType(textModulesMap)
	brandInvoiceTextsType.ThirdParties = createThirdPartiesTextType(textModulesMap)

	brandInvoiceTextsType.DeliveryAddress = textModulesMap["DeliveryAddress"]
	brandInvoiceTextsType.ContractAddress = textModulesMap["ContractAddress"]

	brandInvoiceTextsType.GenericTexts = createBrandSpecificGenericTexts(textModulesMap)

	return &brandInvoiceTextsType
}

func createTexts() *schema.TextsType {
	textsType := schema.TextsType{}

	briefText := schema.TextType{
		Ref: schema.BriefText,
		Pos: schema.BriefText,
	}
	textsType.Text = append(textsType.Text, &briefText)

	hotline := schema.TextType{
		Ref: schema.Hotline,
		Pos: schema.Hotline,
	}
	textsType.Text = append(textsType.Text, &hotline)

	hotlineFootnote := schema.TextType{
		Ref: schema.HotlineFootnote,
		Pos: schema.HotlineFootnote,
	}
	textsType.Text = append(textsType.Text, &hotlineFootnote)

	return &textsType
}

func doStandardReplacements(str string) string {
	if strings.Contains(str, "ns2:") {
		str = strings.ReplaceAll(str, "ns2:", "")
	}
	if strings.Contains(str, "ns4:") {
		str = strings.ReplaceAll(str, "ns4:", "")
	}
	if strings.Contains(str, "xsi:schemaLocation") {
		sub := str[strings.Index(str, "xsi:schemaLocation"):strings.Index(str, ">")]
		str = strings.ReplaceAll(str, sub, schema.LocationString)
	}
	return str
}

func doTagReplacements(str string) string {
	if strings.Contains(str, "BrandInvoiceTextsType") {
		str = strings.ReplaceAll(str, "BrandInvoiceTextsType", "BrandInvoiceTexts")
	}
	if strings.Contains(str, schema.LocationString) {
		str = strings.ReplaceAll(str, schema.LocationString, "")
	}
	return str
}

func doInvoiceSpecificReplacements(xml, rootTagName string) string {
	return removeRootTag(xml, rootTagName)
}

func removeRootTag(xml, rootTagName string) string {
	// Find the start and end positions of the root tag including any attributes
	sb := strings.Builder{}

	sb.WriteString("<")
	sb.WriteString(rootTagName)
	startPos := strings.Index(xml, sb.String())
	endPos := strings.Index(xml, ">") + 1

	// Extract the substring containing the root tag with attributes
	sub := xml[startPos:endPos]

	// Remove the root tag along with its attributes from the original XML
	xml = strings.ReplaceAll(xml, sub, "")

	sb = strings.Builder{}
	sb.WriteString("</")
	sb.WriteString(rootTagName)
	sb.WriteString(">")
	// Remove the closing root tag
	xml = strings.ReplaceAll(xml, sb.String(), "")

	// Trim any leading or trailing whitespace
	xml = strings.TrimSpace(xml)

	return xml
}

func generateHeader(brandIds map[int64]models.BrandAndExternalIds) string {
	sb := strings.Builder{}
	sb.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	sb.WriteString("\n")
	sb.WriteString("<Invoices ")
	sb.WriteString(schema.LocationString)
	sb.WriteString(">\n")
	xmlHeader := sb.String()
	var xmlStrings []string

	for _, brandAndExternalIds := range brandIds {
		brands, err := db.FindAllByClientAndLanguageAndProductLineBrand(int64(1), int64(1), brandAndExternalIds.BrandID)
		textModulesMap := make(map[string]string)
		for _, brand := range brands {
			textModulesMap[brand.TextModuleName] = brand.TextTemplate
		}

		brandInvoiceTextsType := createBrandInvoiceTextsType(brandAndExternalIds.ExternalID, textModulesMap)
		brandInvoiceTextsXml, err := convertToXML(brandInvoiceTextsType)
		if err != nil {
			log.Fatal(err)
		}
		brandInvoiceTextsXml = doTagReplacements(brandInvoiceTextsXml)
		xmlStrings = append(xmlStrings, brandInvoiceTextsXml)
	}

	sb = strings.Builder{}
	sb.WriteString(xmlHeader)
	sb.WriteString(strings.Join(xmlStrings, ""))
	sb.WriteString("\n")
	return sb.String()
}

func generateFooter() string {
	return "\n</Invoices>"
}

func createBrandInvoiceTextsType(brandID string, textModulesMap map[string]string) *schema.BrandInvoiceTextsType {
	brandInvoiceTextsType := createBrandInvoiceTextsTypeCommon(brandID, textModulesMap)
	brandInvoiceTextsType.Subscriptions = createSubscriptions(textModulesMap)

	return brandInvoiceTextsType
}

func createSubscriptions(textModulesMap map[string]string) *schema.SubscriptionsTextType {
	lineTextType := schema.LineTextType{
		Date:     textModulesMap["Line.Date"],
		Time:     textModulesMap["Line.Time"],
		Type:     textModulesMap["Line.Type"],
		BNum:     textModulesMap["Line.BNum"],
		Duration: textModulesMap["Line.Duration"],
		Amount:   textModulesMap["Line.Amount"],
		Comment:  textModulesMap["Line.Comment"],
	}

	itemizedBillTextType := schema.ItemizedBillTextType{
		Line: &lineTextType,
	}

	subscriptionTextType := schema.SubscriptionTextType{
		ItemizedBill:   &itemizedBillTextType,
		Msisdn:         textModulesMap["Subscription.Msisdn"],
		SumGrossAmount: textModulesMap["Subscription.SumGrossAmount"],
		Product:        textModulesMap["Subscription.Product"],
	}

	subscriptionsTextType := schema.SubscriptionsTextType{
		Subscription: &subscriptionTextType,
	}

	return &subscriptionsTextType
}
