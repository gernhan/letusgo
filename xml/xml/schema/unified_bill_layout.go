package schema

const LocationString = "xmlns:fw=\"http://www.formware.de/fox\""

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

// Invoices ...
type Invoices struct {
	BrandInvoiceTexts []*BrandInvoiceTextsType `xml:"BrandInvoiceTexts"`
	Invoice           []*Invoice               `xml:"Invoice"`
}

// DateType ...
type DateType string

// TimeType ...
type TimeType string

// TaxRateType is Tax rate in percent with percentage sign (e.g. 19% or 7,7%)
type TaxRateType string

// InvoiceTextsType ...
type InvoiceTextsType struct {
	InvoiceTitle     string `xml:"InvoiceTitle,attr,omitempty"`
	InvoiceNumber    string `xml:"InvoiceNumber,attr,omitempty"`
	InvoiceDate      string `xml:"InvoiceDate,attr,omitempty"`
	InvoiceNetAmount string `xml:"InvoiceNetAmount,attr,omitempty"`
	InvoiceAmount    string `xml:"InvoiceAmount,attr,omitempty"`
	TaxAmount        string `xml:"TaxAmount,attr,omitempty"`
	InvoicePeriod    string `xml:"InvoicePeriod,attr,omitempty"`
}

// CustomerTextsType ...
type CustomerTextsType struct {
	AccountNumber  string `xml:"AccountNumber,attr,omitempty"`
	CustomerNumber string `xml:"CustomerNumber,attr,omitempty"`
}

// Invoice ...
type Invoice struct {
	InvoiceNumber    string            `xml:"InvoiceNumber,attr,omitempty"`
	InvoiceDate      string            `xml:"InvoiceDate,attr,omitempty"`
	StartDate        string            `xml:"StartDate,attr,omitempty"`
	EndDate          string            `xml:"EndDate,attr,omitempty"`
	InvoiceAmount    string            `xml:"InvoiceAmount,attr,omitempty"`
	InvoiceNetAmount string            `xml:"InvoiceNetAmount,attr,omitempty"`
	TaxAmount        string            `xml:"TaxAmount,attr,omitempty"`
	Header           string            `xml:"Header,attr,omitempty"`
	InvoiceTitle     string            `xml:"InvoiceTitle,attr,omitempty"`
	Layout           string            `xml:"Layout,attr,omitempty"`
	IsOriginal       string            `xml:"IsOriginal,attr,omitempty"`
	BrandId          string            `xml:"BrandId,attr,omitempty"`
	Medium           string            `xml:"Medium,attr,omitempty"`
	Customer         *CustomerType     `xml:"Customer"`
	TaxedAmounts     *TaxedAmounts     `xml:"TaxedAmounts"`
	Subscriptions    *Subscriptions    `xml:"Subscriptions"`
	ThirdParties     *ThirdParties     `xml:"ThirdParties"`
	InvoiceItems     *InvoiceItemsType `xml:"InvoiceItems"`
	Texts            *TextsType        `xml:"Texts"`
}

// TaxedAmounts ...
type TaxedAmounts struct {
	SumNetAmount   string         `xml:"SumNetAmount,attr,omitempty"`
	SumTaxAmount   string         `xml:"SumTaxAmount,attr,omitempty"`
	SumGrossAmount string         `xml:"SumGrossAmount,attr,omitempty"`
	TaxedAmount    []*TaxedAmount `xml:"TaxedAmount"`
}

// TaxedAmount is Overview of the tax rates and sums of net, gross and tax amounts in this invoice
type TaxedAmount struct {
	TaxRate     string `xml:"TaxRate,attr,omitempty"`
	NetAmount   string `xml:"NetAmount,attr,omitempty"`
	TaxAmount   string `xml:"TaxAmount,attr,omitempty"`
	GrossAmount string `xml:"GrossAmount,attr,omitempty"`
}

// Subscriptions ...
type Subscriptions struct {
	Subscription []*Subscription `xml:"Subscription"`
}

// Subscription ...
type Subscription struct {
	Msisdn         string        `xml:"Msisdn,attr,omitempty"`
	SumGrossAmount string        `xml:"SumGrossAmount,attr,omitempty"`
	Product        string        `xml:"Product,attr,omitempty"`
	LineItems      *LineItems    `xml:"LineItems"`
	ItemizedBill   *ItemizedBill `xml:"ItemizedBill"`
}

// LineItems is Aggregated line items to be displayed in the invoice section.
type LineItems struct {
	Item []*Item `xml:"Item"`
}

// Item ...
type Item struct {
	Description string `xml:"Description,attr,omitempty"`
	NetAmount   string `xml:"NetAmount,attr,omitempty"`
	TaxRate     string `xml:"TaxRate,attr,omitempty"`
	GrossAmount string `xml:"GrossAmount,attr,omitempty"`
	GrossOnly   bool   `xml:"GrossOnly,attr,omitempty"`
}

// ItemizedBill is Itemized bill (EVN / EGN)
type ItemizedBill struct {
	Line []*Line `xml:"Line"`
}

// Line ...
type Line struct {
	Date     string `xml:"Date"`
	Time     string `xml:"Time"`
	Duration string `xml:"Duration"`
	Type     string `xml:"Type"`
	BNum     string `xml:"BNum"`
	Amount   string `xml:"Amount"`
	Comment  string `xml:"Comment"`
}

// ThirdParties ...
type ThirdParties struct {
	ThirdParty []*ThirdParty `xml:"ThirdParty"`
}

// ThirdParty ...
type ThirdParty struct {
	Name              string               `xml:"Name,attr,omitempty"`
	Detail            string               `xml:"Detail,attr,omitempty"`
	TotalGrossAmount  string               `xml:"TotalGrossAmount,attr,omitempty"`
	ContactData       string               `xml:"ContactData,attr,omitempty"`
	ThirdPartyService []*ThirdPartyService `xml:"ThirdPartyService"`
}

// ThirdPartyService ...
type ThirdPartyService struct {
	Name        string `xml:"Name,attr,omitempty"`
	Detail      string `xml:"Detail,attr,omitempty"`
	NetAmount   string `xml:"NetAmount,attr,omitempty"`
	TaxRate     string `xml:"TaxRate,attr,omitempty"`
	TaxAmount   string `xml:"TaxAmount,attr,omitempty"`
	GrossAmount string `xml:"GrossAmount,attr,omitempty"`
}

// BrandInvoiceTextsType ...
type BrandInvoiceTextsType struct {
	BrandId         string                  `xml:"BrandId,attr"`
	Email           *EmailType              `xml:"Email"`
	Sender          *MarkupTextType         `xml:"Sender"`
	Footer          *MarkupTextType         `xml:"Footer"`
	GenericTexts    *GenericTextsType       `xml:"GenericTexts"`
	Invoice         *InvoiceTextsType       `xml:"Invoice"`
	Customer        *CustomerTextsType      `xml:"Customer"`
	ContractAddress string                  `xml:"ContractAddress"`
	DeliveryAddress string                  `xml:"DeliveryAddress"`
	InvoiceItems    *InvoiceItemsColumnType `xml:"InvoiceItems"`
	TaxedAmount     *MarkupTextType         `xml:"TaxedAmount"`
	Subscriptions   *SubscriptionsTextType  `xml:"Subscriptions"`
	ThirdParties    *ThirdPartiesTextType   `xml:"ThirdParties"`
}

// ThirdPartiesTextType ...
type ThirdPartiesTextType struct {
	Header     string              `xml:"Header"`
	Text       string              `xml:"Text"`
	ThirdParty *ThirdPartyTextType `xml:"ThirdParty"`
}

// ThirdPartyTextType ...
type ThirdPartyTextType struct {
	ThirdPartyAmountInvoice string `xml:"ThirdPartyAmountInvoice"`
	ThirdPartyAmountEVN     string `xml:"ThirdPartyAmountEVN"`
	Service                 string `xml:"Service"`
	Detail                  string `xml:"Detail"`
	Amount                  string `xml:"Amount"`
	Footnote                string `xml:"Footnote"`
	ThirdPartyDetails       string `xml:"ThirdPartyDetails"`
}

// SubscriptionsTextType ...
type SubscriptionsTextType struct {
	Subscription *SubscriptionTextType `xml:"Subscription"`
}

// SubscriptionTextType ...
type SubscriptionTextType struct {
	Msisdn         string                `xml:"Msisdn,attr,omitempty"`
	SumGrossAmount string                `xml:"SumGrossAmount,attr,omitempty"`
	Product        string                `xml:"Product,attr,omitempty"`
	ItemizedBill   *ItemizedBillTextType `xml:"ItemizedBill"`
}

// ItemizedBillTextType ...
type ItemizedBillTextType struct {
	Line *LineTextType `xml:"Line"`
}

// LineTextType ...
type LineTextType struct {
	Date     string `xml:"Date"`
	Time     string `xml:"Time"`
	Type     string `xml:"Type"`
	BNum     string `xml:"BNum"`
	Duration string `xml:"Duration"`
	Amount   string `xml:"Amount"`
	Comment  string `xml:"Comment"`
}

// GenericTextsType is Dieser Text würde verwendet werden, falls keine
// BrandId-Spezifische Texte vorhanden sind
type GenericTextsType struct {
	Text []*TextType `xml:"Text"`
}

// MarkupTextType ...
type MarkupTextType struct {
	Content string `xml:",innerxml"`
}

// EmailType ...
type EmailType struct {
	From    *MarkupTextType `xml:"From"`
	ReplyTo *MarkupTextType `xml:"ReplyTo"`
	Subject *MarkupTextType `xml:"Subject"`
}

// CustomerType ...
type CustomerType struct {
	AccountNumber   string               `xml:"AccountNumber,attr"`
	CustomerNumber  string               `xml:"CustomerNumber,attr"`
	IBAN            string               `xml:"IBAN,attr,omitempty"`
	SEPAMandateId   string               `xml:"SEPAMandateId,attr,omitempty"`
	InvoiceAddress  *InvoiceAddressType  `xml:"InvoiceAddress"`
	ShippingAddress *ShippingAddressType `xml:"ShippingAddress"`
	EmailToAddress  *EmailToAddressType  `xml:"EmailToAddress"`
}

// InvoiceAddressType ...
type InvoiceAddressType struct {
	Salutation      string `xml:"Salutation,attr,omitempty"`
	Firstname       string `xml:"Firstname,attr,omitempty"`
	Lastname        string `xml:"Lastname,attr,omitempty"`
	Street          string `xml:"Street,attr,omitempty"`
	StreetNumber    string `xml:"StreetNumber,attr,omitempty"`
	AddressAddition string `xml:"AddressAddition,attr,omitempty"`
	PostalCode      string `xml:"PostalCode,attr,omitempty"`
	City            string `xml:"City,attr,omitempty"`
	Country         string `xml:"Country,attr,omitempty"`
	Value           string `xml:",chardata"`
}

// ShippingAddressType ...
type ShippingAddressType struct {
	Salutation      string `xml:"Salutation,attr,omitempty"`
	Firstname       string `xml:"Firstname,attr,omitempty"`
	Lastname        string `xml:"Lastname,attr,omitempty"`
	Street          string `xml:"Street,attr,omitempty"`
	StreetNumber    string `xml:"StreetNumber,attr,omitempty"`
	AddressAddition string `xml:"AddressAddition,attr,omitempty"`
	PostalCode      string `xml:"PostalCode,attr,omitempty"`
	City            string `xml:"City,attr,omitempty"`
	Country         string `xml:"Country,attr,omitempty"`
	Value           string `xml:",chardata"`
}

// EmailToAddressType ...
type EmailToAddressType struct {
	NotificationEmail string `xml:"NotificationEmail,attr,omitempty"`
	Value             string `xml:",chardata"`
}

// TextType ...
type TextType struct {
	Content string `xml:",innerxml"`

	ID  string `xml:"ID,attr,omitempty"`
	Pos string `xml:"Pos,attr,omitempty"`
	Ref string `xml:"Ref,attr,omitempty"`
}

// GermanDecimal ...
type GermanDecimal string

// InvoiceItemsType ...
type InvoiceItemsType struct {
	InvoiceItem []*InvoiceItemType `xml:"InvoiceItem"`
}

// InvoiceItemType ...
type InvoiceItemType struct {
	Product        string          `xml:"Product"`
	ProductDetails *MarkupTextType `xml:"ProductDetails"`
	NetAmount      string          `xml:"NetAmount"`
	TaxRate        string          `xml:"TaxRate"`
}

// TextsType is Texte können hier direkt angegeben oder aus den
// globalen Texten referenziert werden
type TextsType struct {
	Text []*TextType `xml:"Text"`
}

// InvoiceItemsColumnType ...
type InvoiceItemsColumnType struct {
	Header  string `xml:"Header"`
	Amount  string `xml:"Amount"`
	TaxRate string `xml:"TaxRate"`
}
