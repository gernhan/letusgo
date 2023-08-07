package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type LineType string

const (
	LineTypeCDR    LineType = "CDR"
	LineTypeCharge LineType = "CHARGE"
	LineTypeTopUp  LineType = "TOPUP"
)

type ItemizedBillDto struct {
	ID          int64
	BillID      int64
	StartDt     time.Time
	UnitCount   string
	AmountGross decimal.Decimal
	BNumber     string
	Type        string
	Comment     string
	ServiceName string
	LineType    LineType
}

func NewItemizedBillDtoFromPrepaidItemizedBillCdr(obj PrepaidItemizedBillCdr) *ItemizedBillDto {
	return &ItemizedBillDto{
		ID:          obj.ID,
		BillID:      obj.BillID,
		StartDt:     obj.StartDt,
		UnitCount:   obj.UnitCount,
		AmountGross: obj.AmountGross,
		BNumber:     obj.BNumber,
		Type:        obj.Type,
		Comment:     obj.Comment,
		ServiceName: obj.ServiceName,
		LineType:    LineTypeCDR,
	}
}

func NewItemizedBillDtoFromPrepaidItemizedBillCharge(obj PrepaidItemizedBillCharge) *ItemizedBillDto {
	return &ItemizedBillDto{
		ID:          obj.ID,
		BillID:      obj.BillID,
		StartDt:     obj.StartDt,
		UnitCount:   obj.UnitCount,
		AmountGross: obj.AmountGross,
		BNumber:     obj.BNumber,
		Type:        obj.Type,
		Comment:     obj.Comment,
		ServiceName: obj.ServiceName,
		LineType:    LineTypeCharge,
	}
}

func NewItemizedBillDtoFromPrepaidItemizedBillTopUp(obj PrepaidItemizedBillTopUp) *ItemizedBillDto {
	return &ItemizedBillDto{
		ID:          obj.ID,
		BillID:      obj.BillID,
		StartDt:     obj.StartDt,
		UnitCount:   obj.UnitCount,
		AmountGross: obj.AmountGross,
		BNumber:     obj.BNumber,
		Type:        obj.Type,
		Comment:     obj.Comment,
		ServiceName: obj.ServiceName,
		LineType:    LineTypeTopUp,
	}
}
