package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type NormalInvoiceOtherDataHeader struct {
	// done
	X400Admd string
	X400Anam string
	X400Lan  string
	X400Vnam string
	X400Nnam string
	X400Org  string
	X400Ou   string

	// xrech_things
	XrechLeitweg     string
	XrechEmail       string
	XrechAuftragnr   string
	XrechLieferant   string
	XrechVersandName string

	// stm_things
	InvServiceContactTitle  string
	InvServiceContactWeb    string
	InvServiceContactChat   string
	InvServiceContactMail   string
	InvServiceContactPhone  string
	InvServiceContactFooter string
	BankTransferInfo        string
	DirectDebitInfo         string

	// more personal info
	Street      string
	HouseNumber string
	Firma       string
	ZipCode     string
	City        string
	FirstName   string
	LastName    string
	Country     string
	Salutation  string
	Department  string
	Branch      string

	// more_account_things
	ServiceCustomer    string
	AccountAddress     string
	AccountCompanyName string
	AccountZip         string
	AccountCity        string
	AccountFirstName   string
	AccountLastName    string
	AccountCountry     string
	AccountSalutation  string
	AccountDepartment  string
	AccountBranch      string

	// in bill
	// as customer number
	CustomerID  string
	Billmanager int64
	HasEvn      int64
	VatIDNumber string
	IBAN        string
	MandateID   string // should be a different data type, but using string for simplicity
	// as email
	Contact   string
	Amount    decimal.Decimal
	AmountNet decimal.Decimal
	AccountID int64

	// most outer level
	Additional1          string
	Additional2          string
	Kontonummer          string
	Bankname             string
	PaymentMode          int64
	CustID               string
	Billmedia            int64
	BIC                  string
	FrameContractID      string
	InvoiceDueDate       time.Time
	InvoiceDate          time.Time
	Currency             string
	InvoiceID            string
	RechnungszeitraumVon time.Time
	RechnungszeitraumBis time.Time
	AccountHouseNumber   string
	AccountStreet        string
	XrechVersand         string
}

type PreparedData struct {
	VBillHeader               *NormalInvoiceOtherDataHeader
	VTaxes                    *TaxesDto
	VBillSummary              *NormalInvoiceBillSummaryDto
	VBillNoneUsageServiceList []*NormalInvoiceNonUsageDto
	VDiscountsPerServiceList  []*NormalInvoiceDiscountDto
}

func NewPreparedData(
	vBillHeader *NormalInvoiceOtherDataHeader,
	vTaxes *TaxesDto,
	vBillSummary *NormalInvoiceBillSummaryDto,
	vBillNoneUsageServiceList []*NormalInvoiceNonUsageDto,
	vDiscountsPerServiceList []*NormalInvoiceDiscountDto,
) *PreparedData {
	return &PreparedData{
		VBillHeader:               vBillHeader,
		VTaxes:                    vTaxes,
		VBillSummary:              vBillSummary,
		VBillNoneUsageServiceList: vBillNoneUsageServiceList,
		VDiscountsPerServiceList:  vDiscountsPerServiceList,
	}
}

type NormalInvoiceDiscountDto struct {
	ID                   int64               `json:"id"`
	Bill                 int64               `json:"bill"`
	DisplayValue         string              `json:"display_value"`
	OptDisplayValue      string              `json:"opt_display_value"`
	Discount             int64               `json:"discount"`
	DiscountService      int64               `json:"discount_service"`
	DiscountDisplayValue string              `json:"discount_display_value"`
	DiscountPercentage   decimal.NullDecimal `json:"discount_percentage"`
}

type NormalInvoiceNonUsageDto struct {
	ID                        int64               `json:"id"`
	Bill                      int64               `json:"bill"`
	ChargeMode                int64               `json:"charge_mode"`
	DisplayValue              string              `json:"display_value"`
	OptDisplayValue           string              `json:"opt_display_value"`
	FromDate                  time.Time           `json:"from_date"`
	ToDate                    time.Time           `json:"to_date"`
	AmountNet                 decimal.NullDecimal `json:"amount_net"`
	DiscountedAmountNet       decimal.NullDecimal `json:"discounted_amount_net"`
	ComponentOf               int64               `json:"component_of"`
	ServiceName               string              `json:"service_name"`
	Qty                       int64               `json:"qty"`
	ServiceSegment            int64               `json:"service_segment"`
	InvoiceGroup              int64               `json:"invoice_group"`
	ComponentOfChargeMode     int64               `json:"component_of_charge_mode"`
	ComponentOfServiceSegment int64               `json:"component_of_service_segment"`
}

type NormalInvoiceBillSummaryDto struct {
	ID                int64     `json:"id"`
	Bill              int64     `json:"bill_id"`
	ChargeBilledID    int64     `json:"charge_billed_id"`
	Customer          int64     `json:"customer"`
	Account           int64     `json:"account"`
	ContractID        string    `json:"contract_id"`
	Service           string    `json:"service"`
	ContractEndDt     time.Time `json:"contract_end_dt"`
	ServiceName       string    `json:"service_name"`
	StartDtBilling    time.Time `json:"start_dt_billing"`
	TerminationPeriod string    `json:"termination_period"`
	EndDtRequested    time.Time `json:"end_dt_requested"`
}

type CopyableData interface {
	Copy(vBillHeader *NormalInvoiceOtherDataHeader)
}

type NormalInvoiceMoreAccountThings struct {
	ServiceCustomer    string
	AccountAddress     string
	AccountCompanyName string
	AccountZip         string
	AccountCity        string
	AccountFirstName   string
	AccountLastName    string
	AccountSalutation  string
	AccountStreet      string
	AccountHouseNumber string
	AccountCountry     string
	AccountDepartment  string
	AccountBranch      string
}

func (obj *NormalInvoiceMoreAccountThings) Copy(vBillHeader *NormalInvoiceOtherDataHeader) {
	vBillHeader.ServiceCustomer = obj.ServiceCustomer
	vBillHeader.AccountAddress = obj.AccountAddress
	vBillHeader.AccountCompanyName = obj.AccountCompanyName
	vBillHeader.AccountZip = obj.AccountZip
	vBillHeader.AccountCity = obj.AccountCity
	vBillHeader.AccountFirstName = obj.AccountFirstName
	vBillHeader.AccountLastName = obj.AccountLastName
	vBillHeader.AccountSalutation = obj.AccountSalutation
	vBillHeader.AccountStreet = obj.AccountStreet
	vBillHeader.AccountHouseNumber = obj.AccountHouseNumber
	vBillHeader.AccountCountry = obj.AccountCountry
	vBillHeader.AccountDepartment = obj.AccountDepartment
	vBillHeader.AccountBranch = obj.AccountBranch
}

type NormalInvoiceMorePersonalInfo struct {
	FirstName   string
	LastName    string
	Firma       string
	Street      string
	HouseNumber string
	ZipCode     string
	City        string
}

func (obj *NormalInvoiceMorePersonalInfo) Copy(vBillHeader *NormalInvoiceOtherDataHeader) {
	vBillHeader.FirstName = obj.FirstName
	vBillHeader.LastName = obj.LastName
	vBillHeader.Firma = obj.Firma
	vBillHeader.Street = obj.Street
	vBillHeader.HouseNumber = obj.HouseNumber
	vBillHeader.ZipCode = obj.ZipCode
	vBillHeader.City = obj.City
}

type NormalInvoiceXrechThings struct {
	XrechLeitweg     string
	XrechEmail       string
	XrechVersand     string
	XrechLieferant   string
	XrechVersandName string
}

func (obj *NormalInvoiceXrechThings) Copy(vBillHeader *NormalInvoiceOtherDataHeader) {
	vBillHeader.XrechLeitweg = obj.XrechLeitweg
	vBillHeader.XrechEmail = obj.XrechEmail
	vBillHeader.XrechVersand = obj.XrechVersand
	vBillHeader.XrechLieferant = obj.XrechLieferant
	vBillHeader.XrechVersandName = obj.XrechVersandName
}

type NormalInvoiceX400Things struct {
	X400Admd string
	X400Anam string
	X400Lan  string
	X400Vnam string
	X400Nnam string
	X400Org  string
	X400Ou   string
}

func (x400Things *NormalInvoiceX400Things) Copy(vBillHeader *NormalInvoiceOtherDataHeader) {
	vBillHeader.X400Admd = (x400Things.X400Admd)
	vBillHeader.X400Anam = (x400Things.X400Anam)
	vBillHeader.X400Lan = (x400Things.X400Lan)
	vBillHeader.X400Vnam = (x400Things.X400Vnam)
	vBillHeader.X400Nnam = (x400Things.X400Nnam)
	vBillHeader.X400Org = (x400Things.X400Org)
	vBillHeader.X400Ou = (x400Things.X400Ou)
}

type NormalInvoiceStmThings struct {
	InvServiceContactTitle  string
	InvServiceContactWeb    string
	InvServiceContactChat   string
	InvServiceContactMail   string
	InvServiceContactPhone  string
	InvServiceContactFooter string
	BankTransferInfo        string
	DirectDebitInfo         string
}

func (stmThings *NormalInvoiceStmThings) Copy(vBillHeader *NormalInvoiceOtherDataHeader) {
	vBillHeader.InvServiceContactTitle = (stmThings.InvServiceContactTitle)
	vBillHeader.InvServiceContactWeb = (stmThings.InvServiceContactWeb)
	vBillHeader.InvServiceContactChat = (stmThings.InvServiceContactChat)
	vBillHeader.InvServiceContactMail = (stmThings.InvServiceContactMail)
	vBillHeader.InvServiceContactPhone = (stmThings.InvServiceContactPhone)
	vBillHeader.InvServiceContactFooter = (stmThings.InvServiceContactFooter)
	vBillHeader.BankTransferInfo = (stmThings.BankTransferInfo)
	vBillHeader.DirectDebitInfo = (stmThings.DirectDebitInfo)
}

type NormalInvoiceOtherData struct {
	NonUsage             string    `json:"nonUsage"`
	Discount             string    `json:"discount"`
	Summary              string    `json:"summary"`
	Taxes                string    `json:"taxes"`
	StmThings            string    `json:"stmThings"`
	X400Things           string    `json:"x400Things"`
	XrechThings          string    `json:"xrechThings"`
	MorePersonalInfo     string    `json:"morePersonalInfo"`
	MoreAccountThings    string    `json:"moreAccountThings"`
	Additional1          string    `json:"additional1"`
	Additional2          string    `json:"additional2"`
	Kontonummer          string    `json:"kontonummer"`
	Bankname             string    `json:"bankname"`
	PaymentMode          int64     `json:"paymentMode"`
	CustId               string    `json:"custId"`
	Billmedia            int64     `json:"billmedia"`
	Bic                  string    `json:"bic"`
	FrameContractId      string    `json:"frameContractId"`
	InvoiceDueDate       time.Time `json:"invoiceDueDate"`
	InvoiceDate          time.Time `json:"invoiceDate"`
	Currency             string    `json:"currency"`
	InvoiceId            string    `json:"invoiceId"`
	RechnungszeitraumVon time.Time `json:"rechnungszeitraumVon"`
	RechnungszeitraumBis time.Time `json:"rechnungszeitraumBis"`
}

func (otherData *NormalInvoiceOtherData) Copy(vBillHeader *NormalInvoiceOtherDataHeader) {
	vBillHeader.Additional1 = otherData.Additional1
	vBillHeader.Additional2 = otherData.Additional2
	vBillHeader.Kontonummer = otherData.Kontonummer
	vBillHeader.Bankname = otherData.Bankname
	vBillHeader.PaymentMode = otherData.PaymentMode
	vBillHeader.CustID = otherData.CustId
	vBillHeader.Billmedia = otherData.Billmedia
	vBillHeader.BIC = otherData.Bic
	vBillHeader.FrameContractID = otherData.FrameContractId
	vBillHeader.InvoiceDueDate = otherData.InvoiceDueDate
	vBillHeader.InvoiceDate = otherData.InvoiceDate
	vBillHeader.Currency = otherData.Currency
	vBillHeader.InvoiceID = otherData.InvoiceId
	vBillHeader.RechnungszeitraumVon = otherData.RechnungszeitraumVon
	vBillHeader.RechnungszeitraumBis = otherData.RechnungszeitraumBis
}
