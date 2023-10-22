package views

import (
	"database/sql"
	"time"

	"github.com/gernhan/xml/models"
	"github.com/shopspring/decimal"
)

type VExportingBillsV3 struct {
	ID                int64           `db:"id" json:"id"`
	BillRunID         int64           `db:"bill_run_id" json:"billRunId"`
	BillNumber        string          `db:"bill_number" json:"billNumber"`
	PdfStatus         int             `db:"pdf_status" json:"pdfStatus"`
	InvoiceDate       time.Time       `db:"invoice_date" json:"invoiceDate"`
	BillPeriodFrom    time.Time       `db:"bill_period_from" json:"billPeriodFrom"`
	BillPeriodTo      time.Time       `db:"bill_period_to" json:"billPeriodTo"`
	AccountID         int64           `db:"account_id" json:"accountId"`
	Amount            decimal.Decimal `db:"amount" json:"amount"`
	AmountNet         decimal.Decimal `db:"amount_net" json:"amountNet"`
	AccountBrand      string          `db:"account_brand" json:"accountBrand"`
	BillMedia         decimal.Decimal `db:"bill_media" json:"billMedia"`
	CustomerID        int64           `db:"customer_id" json:"customerId"`
	CustomerNumber    string          `db:"customer_number" json:"customerNumber"`
	CustomerClient    int64           `db:"customer_client" json:"customerClient"`
	CustomerLanguage  int64           `db:"customer_language" json:"customerLanguage"`
	InvoiceReference  string          `db:"invoice_reference" json:"invoiceReference"`
	SalutationString  string          `db:"salutation_string" json:"salutationString"`
	FirstName         string          `db:"first_name" json:"firstName"`
	LastName          string          `db:"last_name" json:"lastName"`
	Street            string          `db:"street" json:"street"`
	HouseNumber       string          `db:"house_number" json:"houseNumber"`
	FreeText          sql.NullString  `db:"free_text" json:"freeText"`
	ZIP               string          `db:"zip" json:"zip"`
	City              string          `db:"city" json:"city"`
	Alpha2            string          `db:"alpha2" json:"alpha2"`
	AdditionalAddress string          `db:"additional_address" json:"additionalAddress"`
	Email             string          `db:"email" json:"email"`
	IBAN              string          `db:"iban" json:"iban"`
	MandateID         sql.NullString  `db:"mandate_id" json:"mandateId"`
	ProductDetails    string          `db:"product_details" json:"productDetails"`
	BillType          int64           `db:"bill_type" json:"billType"`
	HistoricBill      int             `db:"historic_bill" json:"historicBill"`
	BillFrequency     int64           `db:"bill_frequency" json:"billFrequency"`
	BrandID           int64           `db:"brand_id" json:"brandId"`
	Billmanager       int             `db:"billmanager" json:"billmanager"`
	HasEVN            int64           `db:"has_evn" json:"hasEvn"`
	VatIDNumber       string          `db:"vat_id_number" json:"vatIdNumber"`
	OtherData         string          `db:"other_data" json:"otherData"`
}

func (obj VExportingBillsV3) Copy(vBillHeader *models.NormalInvoiceOtherDataHeader) {
	vBillHeader.CustomerID = obj.CustomerNumber
	vBillHeader.Billmanager = obj.Billmanager
	vBillHeader.HasEvn = obj.HasEVN
	vBillHeader.VatIDNumber = obj.VatIDNumber
	vBillHeader.IBAN = obj.IBAN
	vBillHeader.MandateID = obj.MandateID.String
	vBillHeader.Contact = obj.Email
	vBillHeader.Amount = obj.Amount
	vBillHeader.AmountNet = obj.AmountNet
	vBillHeader.AccountID = obj.AccountID
}
