package models

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

type PrepaidInvoiceOtherData struct {
	DateFrom            time.Time `json:"rechnungszeitraum_von"`
	DateTo              time.Time `json:"rechnungszeitraum_bis"`
	Taxes               string    `json:"taxes"`
	BillSummary         string    `json:"bill_summary"`
	ItemizedBillCDRs    string    `json:"itemized_bill_cdrs"`
	ItemizedBillCharges string    `json:"itemized_bill_charges"`
	ItemizedBillTopUps  string    `json:"itemized_bill_top_ups"`
	ThirdParties        string    `json:"third_parties"`
}

// Custom unmarshaler for the time.Time fields
func (p *PrepaidInvoiceOtherData) UnmarshalJSON(data []byte) error {
	type Alias PrepaidInvoiceOtherData
	aux := &struct {
		DateFrom string `json:"rechnungszeitraum_von"`
		DateTo   string `json:"rechnungszeitraum_bis"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error
	p.DateFrom, err = time.Parse("2006-01-02T15:04:05", aux.DateFrom)
	if err != nil {
		return err
	}

	p.DateTo, err = time.Parse("2006-01-02T15:04:05", aux.DateTo)
	if err != nil {
		return err
	}

	if len(p.ItemizedBillCDRs) == 0 {
		p.ItemizedBillCDRs = "[]"
	}

	if len(p.ItemizedBillCharges) == 0 {
		p.ItemizedBillCharges = "[]"
	}

	if len(p.ItemizedBillTopUps) == 0 {
		p.ItemizedBillTopUps = "[]"
	}

	if len(p.Taxes) == 0 {
		p.Taxes = "[]"
	}

	if len(p.ThirdParties) == 0 {
		p.ThirdParties = "[]"
	}

	return nil
}

func main() {
	jsonStr := `{
		"rechnungszeitraum_von": "2023-06-01T00:00:00",
		"rechnungszeitraum_bis": "2023-06-30T23:59:59",
		"taxes": [],
		"bill_summary": [],
		"itemized_bill_cdrs": [],
		"itemized_bill_charges": [],
		"itemized_bill_top_ups": [],
		"third_parties": []
	}`

	var data PrepaidInvoiceOtherData
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Parsed DateFrom:", data.DateFrom)
	fmt.Println("Parsed DateTo:", data.DateTo)
}

type TaxesDto struct {
	ID      int64           `json:"id"`
	Bill    int64           `json:"bill"`
	TaxRate float64         `json:"tax_rate"`
	Sum     decimal.Decimal `json:"sum"`
	SumNet  decimal.Decimal `json:"sum_net"`
	Tax     decimal.Decimal `json:"tax"`
}

type PrepaidBillSummaryDto struct {
	ID              int64           `json:"id"`
	Bill            int64           `json:"bill"`
	AmountNet       decimal.Decimal `json:"amount_net"`
	Amount          decimal.Decimal `json:"amount"`
	InvoicePosition string          `json:"invoice_position"`
	TaxRate         decimal.Decimal `json:"tax_rate"`
	MSISDN          string          `json:"msisdn"`
	ProductName     string          `json:"product_name"`
}

type PrepaidItemizedBillCdr struct {
	ID          int64           `json:"id"`
	BillID      int64           `json:"bill_id"`
	StartDt     time.Time       `json:"start_dt"`
	UnitCount   string          `json:"unit_count"`
	AmountGross decimal.Decimal `json:"amount_gross"`
	BNumber     string          `json:"b_number"`
	Type        string          `json:"type"`
	Comment     string          `json:"comment"`
	ServiceName string          `json:"service_name"`
}

type PrepaidItemizedBillCharge struct {
	ID          int64           `json:"id"`
	BillID      int64           `json:"bill_id"`
	StartDt     time.Time       `json:"start_dt"`
	UnitCount   string          `json:"unit_count"`
	AmountGross decimal.Decimal `json:"amount_gross"`
	BNumber     string          `json:"b_number"`
	Type        string          `json:"type"`
	Comment     string          `json:"comment"`
	ServiceName string          `json:"service_name"`
}

type PrepaidItemizedBillTopUp struct {
	ID          int64           `json:"id"`
	BillID      int64           `json:"bill_id"`
	ServiceID   string          `json:"service_id"`
	StartDt     time.Time       `json:"start_dt"`
	UnitCount   string          `json:"unit_count"`
	AmountGross decimal.Decimal `json:"amount_gross"`
	BNumber     string          `json:"b_number"`
	Type        string          `json:"type"`
	Comment     string          `json:"comment"`
	ServiceName string          `json:"service_name"`
}

type PrepaidThirdParty struct {
	ID           string          `json:"id"`
	ProviderID   string          `json:"provider_id"`
	Sum          decimal.Decimal `json:"sum"`
	PartnerText  string          `json:"partner_text"`
	ProviderName string          `json:"provider_name"`
	StreetName   string          `json:"street_name"`
	StreetNumber string          `json:"street_number"`
	ZipCode      string          `json:"zip_code"`
	City         string          `json:"city"`
	Country      string          `json:"country"`
	PhoneNumber  string          `json:"phone_number"`
	EmailAddress string          `json:"email_address"`
	InvoiceText  string          `json:"invoice_text"`
}
