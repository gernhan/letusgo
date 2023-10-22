package xml_handlers

import (
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/gernhan/xml/entities/views"
	"github.com/gernhan/xml/models"
	"github.com/gernhan/xml/xml/schema"
	"github.com/shopspring/decimal"
)

// VAS_EXCLUDED and VAS_REFUND_EXCLUDED are constants
const (
	VAS_EXCLUDED        = "Bezahlen per Handyrechnung inkl. MwSt. (siehe Details)"
	VAS_REFUND_EXCLUDED = "Gutschriften für Bezahlen per Handyrechnung inkl. MwSt."
)

// Constants
const (
	DefaultIsOriginal = "true"
	VasRefundExcluded = "Gutschriften für Bezahlen per Handyrechnung inkl. MwSt."
	VasExcluded       = "Bezahlen per Handyrechnung inkl. MwSt. (siehe Details)"
)

// BillMedia represents the enumeration-like behavior for BillMedia
type BillMedia struct {
	ID    int64
	Value string
}

var (
	Papierrechnung   = BillMedia{ID: 2, Value: "Papierrechnung"}
	Online           = BillMedia{ID: 4, Value: "Online"}
	XRechnung        = BillMedia{ID: 5, Value: "Rechnung"}
	EDIFACT95APAPIER = BillMedia{ID: 6, Value: "DIFACT95APAPIER"}
	EDIFACT95BPAPIER = BillMedia{ID: 7, Value: "DIFACT95BPAPIER"}
	EDIFACT99B       = BillMedia{ID: 8, Value: "DIFACT99B"}
)

var SupportedBillMedias = []BillMedia{
	Papierrechnung,
	Online,
}

type InvoiceBuilder struct {
	invoice schema.Invoice
}

func NewInvoiceBuilderV3() *InvoiceBuilder {
	builder := &InvoiceBuilder{}
	return builder
}

// Assuming you have already defined the "Invoice" struct and other types used in the method.
func (b *InvoiceBuilder) header(bill views.VExportingBillsV3, otherData models.PrepaidInvoiceOtherData) *InvoiceBuilder {
	b.invoice.InvoiceNumber = bill.BillNumber
	b.invoice.InvoiceDate = formatDate(bill.InvoiceDate)
	b.invoice.StartDate = formatDate(otherData.DateFrom)
	b.invoice.EndDate = formatDate(otherData.DateTo)
	b.invoice.InvoiceAmount = amountString(&bill.Amount)
	b.invoice.InvoiceNetAmount = amountString(&bill.AmountNet)
	sub := bill.Amount.Sub(bill.AmountNet)
	b.invoice.TaxAmount = amountString(&sub)
	b.invoice.IsOriginal = DefaultIsOriginal // Will become relevant when implementing bill copies - until then, always true
	b.invoice.BrandId = bill.AccountBrand
	b.invoice.Layout = "PrepaidBill"

	var billMediaValue string
	for _, bm := range SupportedBillMedias {
		if bill.BillMedia.Equal(decimal.NewFromInt(int64(bm.ID))) {
			billMediaValue = bm.Value
			break
		}
	}
	if billMediaValue == "" {
		billMediaValue = Papierrechnung.Value // Default to "PAPIERRECHNUNG" if no matching billmedia is found
	}
	b.invoice.Medium = billMediaValue
	return b
}

// formatDate is a helper function to format the date as "dd.MM.yyyy".
func formatDate(date time.Time) string {
	return date.Format("02.01.2006")
}

// formatTime is a helper function to format the time as "hh:mm:ss".
func formatTime(dateTime time.Time) string {
	return dateTime.Format("15:04:05")
}

func (b *InvoiceBuilder) customer(customerType *schema.CustomerType) *InvoiceBuilder {
	b.invoice.Customer = customerType
	return b
}

func (b *InvoiceBuilder) taxes(taxes []models.TaxesDto) *InvoiceBuilder {
	taxedAmounts := schema.TaxedAmounts{}
	var totalNetAmount, totalTaxAmount, totalGrossAmount decimal.Decimal

	for _, tax := range taxes {
		taxedAmount := schema.TaxedAmount{}

		if tax.Tax.Cmp(decimal.NewFromInt(0)) == 0 {
			taxedAmount.TaxRate = "0%"
		} else {
			//taxRate := tax.TaxRate * 100
			taxRate := tax.TaxRate
			taxRateInt := int64(taxRate)
			sb := strings.Builder{}
			sb.WriteString(big.NewInt(taxRateInt).String())
			sb.WriteString("%")
			taxedAmount.TaxRate = sb.String()
		}

		if tax.SumNet.Cmp(decimal.NewFromInt(0)) == 0 {
			taxedAmount.NetAmount = "0,00"
		} else {
			taxedAmount.NetAmount = amountString(toMonetary(&tax.SumNet))
		}

		if tax.Tax.Cmp(decimal.NewFromInt(0)) == 0 {
			taxedAmount.TaxAmount = "0,00"
		} else {
			taxedAmount.TaxAmount = amountString(toMonetary(&tax.Tax))
		}

		taxedAmount.GrossAmount = amountString(toMonetary(&tax.Sum))
		taxedAmounts.TaxedAmount = append(taxedAmounts.TaxedAmount, &taxedAmount)

		totalNetAmount = totalNetAmount.Add(tax.SumNet)
		totalTaxAmount = totalTaxAmount.Add(tax.Tax)
		totalGrossAmount = totalGrossAmount.Add(tax.Sum)
	}

	if totalNetAmount.Cmp(decimal.NewFromInt(0)) == 0 {
		taxedAmounts.SumNetAmount = "0,00"
	} else {
		taxedAmounts.SumNetAmount = amountString(toMonetary(&totalNetAmount))
	}

	if totalTaxAmount.Cmp(decimal.NewFromInt(0)) == 0 {
		taxedAmounts.SumTaxAmount = "0,00"
	} else {
		taxedAmounts.SumTaxAmount = amountString(toMonetary(&totalTaxAmount))
	}

	taxedAmounts.SumGrossAmount = amountString(toMonetary(&totalGrossAmount))

	b.invoice.TaxedAmounts = &taxedAmounts

	return b
}

func toMonetary(f *decimal.Decimal) *decimal.Decimal {
	round := f.Round(2)
	return &round
}

func amountString(f *decimal.Decimal) string {
	return replaceDecimalPointWithComma(f.StringFixed(2))
}

func replaceDecimalPointWithComma(s string) string {
	if len(s) == 0 {
		return s
	}

	// Find the position of the decimal point
	dotIndex := -1
	for i := range s {
		if s[i] == '.' {
			dotIndex = i
			break
		}
	}

	// If the decimal point is found, replace it with a comma
	if dotIndex >= 0 {
		s = s[:dotIndex] + "," + s[dotIndex+1:]
	}

	return s
}

func (b *InvoiceBuilder) texts(textsType *schema.TextsType) *InvoiceBuilder {
	b.invoice.Texts = textsType
	return b
}

func (b *InvoiceBuilder) thirdParties(thirdParties schema.ThirdParties) *InvoiceBuilder {
	b.invoice.ThirdParties = &thirdParties
	return b
}

func (b *InvoiceBuilder) build() *schema.Invoice {
	return &b.invoice
}

func amountStringFourDigits(amount decimal.Decimal) string {
	return amount.StringFixed(2)
}

// subscriptions represents the Go implementation of the subscriptions method
func (b *InvoiceBuilder) subscriptions(prepaidBillSummaries []models.PrepaidBillSummaryDto, prepaidItemizedBills []models.ItemizedBillDto) *InvoiceBuilder {
	subscriptions := schema.Subscriptions{}

	totalGrossAmount := decimal.NewFromInt(0)

	// group by msisdn
	MSISDNs := make(map[string]string)
	for _, summary := range prepaidBillSummaries {
		MSISDNs[summary.MSISDN] = summary.MSISDN
	}
	msisdnList := make([]string, 0, len(MSISDNs))
	for msisdn := range MSISDNs {
		msisdnList = append(msisdnList, msisdn)
	}
	sort.Strings(msisdnList)

	for _, msisdn := range msisdnList {
		subscription := schema.Subscription{}
		lineItems := schema.LineItems{}
		prepaidBillSummariesSameMSISDNs := make([]models.PrepaidBillSummaryDto, 0)
		for _, summary := range prepaidBillSummaries {
			if summary.MSISDN == msisdn {
				prepaidBillSummariesSameMSISDNs = append(prepaidBillSummariesSameMSISDNs, summary)
			}
		}
		prepaidItemizedBillsSameMSISDNs := make([]models.ItemizedBillDto, 0)
		for _, itemizedBill := range prepaidItemizedBills {
			if itemizedBill.ServiceName == msisdn {
				prepaidItemizedBillsSameMSISDNs = append(prepaidItemizedBillsSameMSISDNs, itemizedBill)
			}
		}

		for _, summary := range prepaidBillSummariesSameMSISDNs {
			item := schema.Item{}
			item.Description = summary.InvoicePosition
			item.NetAmount = amountStringFourDigits(summary.AmountNet)

			sb := strings.Builder{}
			sb.WriteString(summary.TaxRate.StringFixed(2))
			sb.WriteString("%")
			item.TaxRate = sb.String()
			item.GrossAmount = amountString(toMonetary(&summary.Amount))
			if VAS_EXCLUDED == summary.InvoicePosition || VAS_REFUND_EXCLUDED == summary.InvoicePosition {
				item.NetAmount = "0"
				item.TaxRate = "0%"
			}

			lineItems.Item = append(lineItems.Item, &item)

			totalGrossAmount.Add(summary.Amount)
			if subscription.Msisdn == "" {
				subscription.Msisdn = summary.MSISDN
			}

			if subscription.Product == "" {
				subscription.Product = summary.ProductName
			}
		}
		subscription.SumGrossAmount = amountString(toMonetary(&totalGrossAmount))
		subscription.LineItems = &lineItems

		itemizedBill := schema.ItemizedBill{}
		for _, itemizedBillLine := range prepaidItemizedBillsSameMSISDNs {
			line := schema.Line{}
			line.Date = formatDate(itemizedBillLine.StartDt)
			line.Time = formatTime(itemizedBillLine.StartDt)
			line.Type = itemizedBillLine.Type
			line.BNum = itemizedBillLine.BNumber
			line.Duration = itemizedBillLine.UnitCount
			line.Amount = amountString(toMonetary(&itemizedBillLine.AmountGross))
			line.Comment = itemizedBillLine.Comment

			itemizedBill.Line = append(itemizedBill.Line, &line)
		}
		subscription.ItemizedBill = &itemizedBill

		subscriptions.Subscription = append(subscriptions.Subscription, &subscription)
	}
	b.invoice.Subscriptions = &subscriptions

	return b
}
