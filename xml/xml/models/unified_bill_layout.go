package xml

import (
	"encoding/xml"
	"fmt"
	"regexp"
)

const SchemaLocationString = "xmlns:fw=\"http://www.formware.de/fox\""

const (
	Legal              string = "Legal"
	PaymentDebit       string = "PaymentDebit"
	PaymentCredit      string = "PaymentCredit"
	Klarna             string = "Klarna"
	Paypal             string = "Paypal"
	PaymentCreditCard  string = "PaymentCreditCard"
	BriefText          string = "Brieftext"
	HotlineNormal      string = "HotlineNormal"
	Hotline            string = "Hotline"
	HotlineFootnote    string = "HotlineFootnote"
	HotlineVip         string = "HotlineVIP"
	HotlineVipFootnote string = "HotlineVIPFootnote"
	Advert             string = "Advert"
)

var BrandGenericTextsLabelsList = []string{
	Legal,
	PaymentDebit,
	PaymentCredit,
	Klarna,
	Paypal,
	PaymentCreditCard,
	BriefText,
	HotlineNormal,
	HotlineFootnote,
	HotlineVip,
	HotlineVipFootnote,
	Advert,
}

type ItemizedBillTextType struct {
	Line LineTextType `xml:"Line"`
}

type LineTextType struct {
	Date     string `xml:"Date"`
	Time     string `xml:"Time"`
	Type     string `xml:"Type"`
	BNum     string `xml:"BNum"`
	Duration string `xml:"Duration"`
	Amount   string `xml:"Amount"`
	Comment  string `xml:"Comment"`
}

func ValidateDate(d string) bool {
	pattern := `^(0?[1-9]|[1-2][0-9]|3[0-1])\.(0?[1-9]|1[0-2])\.\d{4}$`
	match, _ := regexp.MatchString(pattern, d)
	return match
}

func ValidateTime(t string) bool {
	pattern := `^([01]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`
	match, _ := regexp.MatchString(pattern, t)
	return match
}

// ValidateTaxRate validates the TaxRateType using a regular expression.
func ValidateTaxRate(tr string) bool {
	pattern := `[0-9,]{1,3}\%`
	match, _ := regexp.MatchString(pattern, tr)
	return match
}

// XML representation

// DateTypeXML represents the XML structure for DateType.
type DateTypeXML struct {
	XMLName xml.Name `xml:"DateType"`
	Value   string   `xml:",chardata"`
}

// TimeTypeXML represents the XML structure for TimeType.
type TimeTypeXML struct {
	XMLName xml.Name `xml:"TimeType"`
	Value   string   `xml:",chardata"`
}

// TaxRateTypeXML represents the XML structure for TaxRateType.
type TaxRateTypeXML struct {
	XMLName xml.Name `xml:"TaxRateType"`
	Value   string   `xml:",chardata"`
}

type BrandInvoiceTextsType struct {
	BrandID         string                 `xml:"BrandId,attr"`
	Email           EmailType              `xml:"Email"`
	Sender          MarkupTextType         `xml:"Sender"`
	Footer          MarkupTextType         `xml:"Footer"`
	GenericTexts    GenericTextsType       `xml:"GenericTexts"`
	Invoice         InvoiceTextsType       `xml:"Invoice"`
	Customer        CustomerTextsType      `xml:"Customer"`
	ContractAddress string                 `xml:"ContractAddress"`
	DeliveryAddress string                 `xml:"DeliveryAddress"`
	InvoiceItems    InvoiceItemsColumnType `xml:"InvoiceItems"`
	TaxedAmount     MarkupTextType         `xml:"TaxedAmount"`
	Subscriptions   SubscriptionsTextType  `xml:"Subscriptions"`
	ThirdParties    ThirdPartiesTextType   `xml:"ThirdParties"`
}

// Invoices struct
type Invoices struct {
	BrandInvoiceTexts []BrandInvoiceTextsType `xml:"BrandInvoiceTexts"`
	Invoice           []Invoice               `xml:"Invoice"`
}

type EmailType struct {
	From    MarkupTextType `xml:"From,omitempty"`
	ReplyTo MarkupTextType `xml:"ReplyTo,omitempty"`
	Subject MarkupTextType `xml:"Subject,omitempty"`
}

type MarkupTextType struct {
	Content string `xml:",innerxml"`
}

type GenericTextsType struct {
	Text []TextType `xml:"Text,omitempty"`
}

type TextType struct {
	Content string `xml:",innerxml"`
	ID      string `xml:"ID,attr,omitempty"`
	Pos     string `xml:"Pos,attr,omitempty"`
	Ref     string `xml:"Ref,attr,omitempty"`
}

type InvoiceItemsColumnType struct {
	Header  string `xml:"Header"`
	Amount  string `xml:"Amount"`
	TaxRate string `xml:"TaxRate"`
}

type InvoiceTextsType struct {
	InvoiceTitle     string `xml:"InvoiceTitle,attr,omitempty"`
	InvoiceNumber    string `xml:"InvoiceNumber,attr,omitempty"`
	InvoiceDate      string `xml:"InvoiceDate,attr,omitempty"`
	InvoiceNetAmount string `xml:"InvoiceNetAmount,attr,omitempty"`
	InvoiceAmount    string `xml:"InvoiceAmount,attr,omitempty"`
	TaxAmount        string `xml:"TaxAmount,attr,omitempty"`
	InvoicePeriod    string `xml:"InvoicePeriod,attr,omitempty"`
}

type CustomerTextsType struct {
	AccountNumber  string `xml:"AccountNumber,attr,omitempty"`
	CustomerNumber string `xml:"CustomerNumber,attr,omitempty"`
}

type Invoice struct {
	Customer         CustomerType     `xml:"Customer"`
	TaxedAmount      TaxedAmounts     `xml:"TaxedAmounts"`
	Subscriptions    Subscriptions    `xml:"Subscriptions,omitempty"`
	ThirdParties     ThirdParties     `xml:"ThirdParties,omitempty"`
	InvoiceItems     InvoiceItemsType `xml:"InvoiceItems,omitempty"`
	Texts            TextsType        `xml:"Texts,omitempty"`
	InvoiceNumber    string           `xml:"InvoiceNumber,attr,omitempty"`
	InvoiceDate      string           `xml:"InvoiceDate,attr,omitempty"`
	StartDate        string           `xml:"StartDate,attr,omitempty"`
	EndDate          string           `xml:"EndDate,attr,omitempty"`
	InvoiceAmount    string           `xml:"InvoiceAmount,attr,omitempty"`
	InvoiceNetAmount string           `xml:"InvoiceNetAmount,attr,omitempty"`
	TaxAmount        string           `xml:"TaxAmount,attr,omitempty"`
	Header           string           `xml:"Header,attr,omitempty"`
	InvoiceTitle     string           `xml:"InvoiceTitle,attr,omitempty"`
	Layout           string           `xml:"Layout,attr,omitempty"`
	IsOriginal       bool             `xml:"IsOriginal,attr,omitempty"`
	BrandId          string           `xml:"BrandId,attr,omitempty"`
	Medium           string           `xml:"Medium,attr,omitempty"`
}

type CustomerType struct {
	InvoiceAddress  InvoiceAddressType  `xml:"InvoiceAddress"`
	ShippingAddress ShippingAddressType `xml:"ShippingAddress,omitempty"`
	EmailToAddress  EmailToAddressType  `xml:"EmailToAddress,omitempty"`
	AccountNumber   string              `xml:"AccountNumber,attr,required"`
	CustomerNumber  string              `xml:"CustomerNumber,attr,required"`
	IBAN            string              `xml:"IBAN,attr,omitempty"`
	SEPAMandateId   string              `xml:"SEPAMandateId,attr,omitempty"`
}

type InvoiceAddressType struct {
	Value           string `xml:",chardata"`
	Salutation      string `xml:"Salutation,attr,omitempty"`
	Firstname       string `xml:"Firstname,attr,omitempty"`
	Lastname        string `xml:"Lastname,attr,omitempty"`
	Street          string `xml:"Street,attr,omitempty"`
	StreetNumber    string `xml:"StreetNumber,attr,omitempty"`
	AddressAddition string `xml:"AddressAddition,attr,omitempty"`
	PostalCode      string `xml:"PostalCode,attr,omitempty"`
	City            string `xml:"City,attr,omitempty"`
	Country         string `xml:"Country,attr,omitempty"`
}

type ShippingAddressType struct {
	Value           string `xml:",chardata"`
	Salutation      string `xml:"Salutation,omitempty"`
	Firstname       string `xml:"Firstname,omitempty"`
	Lastname        string `xml:"Lastname,omitempty"`
	Street          string `xml:"Street,omitempty"`
	StreetNumber    string `xml:"StreetNumber,omitempty"`
	AddressAddition string `xml:"AddressAddition,omitempty"`
	PostalCode      string `xml:"PostalCode,omitempty"`
	City            string `xml:"City,omitempty"`
	Country         string `xml:"Country,omitempty"`
}

type EmailToAddressType struct {
	Value             string `xml:",chardata"`
	NotificationEmail string `xml:"NotificationEmail,attr,omitempty"`
	InvoiceEmail      string `xml:"InvoiceEmail,attr,omitempty"`
	RemittanceEmail   string `xml:"RemittanceEmail,attr,omitempty"`
}

// TaxedAmount represents the tax rate and amounts in an invoice
type TaxedAmount struct {
	TaxRate     string `xml:"TaxRate,attr,omitempty"`
	NetAmount   string `xml:"NetAmount,attr,omitempty"`
	TaxAmount   string `xml:"TaxAmount,attr,omitempty"`
	GrossAmount string `xml:"GrossAmount,attr,omitempty"`
}

// TaxedAmounts represents a collection of TaxedAmount elements
type TaxedAmounts struct {
	TaxedAmount    []TaxedAmount `xml:"TaxedAmount"`
	SumNetAmount   string        `xml:"SumNetAmount,attr,omitempty"`
	SumTaxAmount   string        `xml:"SumTaxAmount,attr,omitempty"`
	SumGrossAmount string        `xml:"SumGrossAmount,attr,omitempty"`
}

// validateGermanDecimal checks if the GermanDecimal format is valid
func validateGermanDecimal(dec string) bool {
	re := regexp.MustCompile(`^\d+(\.\d{1,2})?$`)
	return re.MatchString(string(dec))
}

// NewGermanDecimal creates a new GermanDecimal
func NewGermanDecimal(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

// IsZero checks if the GermanDecimal is zero
func IsZero(dec string) bool {
	return dec == "0" || dec == "0.00"
}

// IsValid checks if the GermanDecimal is valid
func IsValid(dec string) bool {
	return validateGermanDecimal(dec)
}

// Subscriptions represents the Subscriptions struct
type Subscriptions struct {
	Subscription []Subscription `xml:"Subscription"`
}

// Subscription represents the Subscription struct
type Subscription struct {
	Msisdn         string       `xml:"Msisdn,attr"`
	SumGrossAmount string       `xml:"SumGrossAmount,attr"`
	Product        string       `xml:"Product,attr"`
	LineItems      LineItems    `xml:"LineItems"`
	ItemizedBill   ItemizedBill `xml:"ItemizedBill"`
}

// LineItems represents the LineItems struct
type LineItems struct {
	Item []Item `xml:"Item"`
}

// Item represents the Item struct
type Item struct {
	Description string `xml:"Description,attr"`
	NetAmount   string `xml:"NetAmount,attr"`
	TaxRate     string `xml:"TaxRate,attr"`
	GrossAmount string `xml:"GrossAmount,attr"`
	GrossOnly   bool   `xml:"GrossOnly,attr"`
}

// ItemizedBill represents the ItemizedBill struct
type ItemizedBill struct {
	Line []Line `xml:"Line"`
}

// Line represents the Line struct
type Line struct {
	Date     string `xml:"Date"`
	Time     string `xml:"Time"`
	Duration string `xml:"Duration"`
	Type     string `xml:"Type"`
	BNum     string `xml:"BNum"`
	Amount   string `xml:"Amount"`
	Comment  string `xml:"Comment"`
}

// Define Go equivalents of the XML complex types
type ThirdParties struct {
	ThirdParty []ThirdParty `xml:"ThirdParty"`
}

type ThirdParty struct {
	ThirdPartyServices []ThirdPartyService `xml:"ThirdPartyService"`
	Name               string              `xml:"Name,attr"`
	Detail             string              `xml:"Detail,attr"`
	TotalGrossAmount   string              `xml:"TotalGrossAmount,attr"`
	ContactData        string              `xml:"ContactData,attr"`
}

type ThirdPartyService struct {
	Name        string `xml:"Name,attr"`
	Detail      string `xml:"Detail,attr"`
	NetAmount   string `xml:"NetAmount,attr"`
	TaxRate     string `xml:"TaxRate,attr"`
	TaxAmount   string `xml:"TaxAmount,attr"`
	GrossAmount string `xml:"GrossAmount,attr"`
}

type ThirdPartyItemType struct {
	Quantity             float64 `xml:"Quantity,attr,omitempty"`
	CategoryName         string  `xml:"CategoryName,attr,omitempty"`
	ServiceStartDate     string  `xml:"ServiceStartDate,attr,omitempty"`
	ServiceEndDate       string  `xml:"ServiceEndDate,attr,omitempty"`
	UnitPrice            string  `xml:"UnitPrice,attr,omitempty"`
	NetAmount            string  `xml:"NetAmount,attr,omitempty"`
	TaxAmount            string  `xml:"TaxAmount,attr,omitempty"`
	InvoiceAmount        string  `xml:"InvoiceAmount,attr,omitempty"`
	ThirdPartyIdentifier string  `xml:"ThirdPartyIdentifier,attr,omitempty"`
}

type InvoiceItemsType struct {
	InvoiceItem []InvoiceItemType `xml:"InvoiceItem,omitempty"`
}

type InvoiceItemType struct {
	Position      float64 `xml:"Position,attr,omitempty"`
	Quantity      float64 `xml:"Quantity,attr,omitempty"`
	Unit          string  `xml:"Unit,attr,omitempty"`
	Description   string  `xml:"Description,attr,omitempty"`
	UnitPrice     string  `xml:"UnitPrice,attr,omitempty"`
	NetAmount     string  `xml:"NetAmount,attr,omitempty"`
	TaxRate       string  `xml:"TaxRate,attr,omitempty"`
	TaxAmount     string  `xml:"TaxAmount,attr,omitempty"`
	InvoiceAmount string  `xml:"InvoiceAmount,attr,omitempty"`
}

type SubscriptionsTextType struct {
	Subscription SubscriptionTextType `xml:"Subscription"`
}

type SubscriptionTextType struct {
	ItemizedBill   ItemizedBillTextType `xml:"ItemizedBill"`
	Msisdn         string               `xml:"Msisdn,attr"`
	SumGrossAmount string               `xml:"SumGrossAmount,attr"`
	Product        string               `xml:"Product,attr"`
}

type ThirdPartyTextType struct {
	ThirdPartyAmountInvoice string `xml:"ThirdPartyAmountInvoice"`
	ThirdPartyAmountEVN     string `xml:"ThirdPartyAmountEVN"`
	Service                 string `xml:"Service"`
	Detail                  string `xml:"Detail"`
	Amount                  string `xml:"Amount"`
	Footnote                string `xml:"Footnote"`
	ThirdPartyDetails       string `xml:"ThirdPartyDetails"`
}

type ThirdPartiesTextType struct {
	Header     string             `xml:"Header"`
	Text       string             `xml:"Text"`
	ThirdParty ThirdPartyTextType `xml:"ThirdParty"`
}

type TextsType struct {
	Text []TextType `xml:"Text,omitempty"`
}
