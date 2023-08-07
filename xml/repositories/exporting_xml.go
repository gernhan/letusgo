package repositories

import (
	"context"
	"github.com/gernhan/xml/entities/views"
	"log"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func FindByBillRunV3(pool *pgxpool.Pool, billRunID, minID, maxID int64) (<-chan views.VExportingBillsV3, <-chan error) {
	query := `SELECT * FROM v_exporting_bills_v3
		WHERE bill_run_id = $1 AND id <= $2 AND id >= $3`

	rows, err := pool.Query(context.Background(), query, billRunID, maxID, minID)
	if err != nil {
		errCh := make(chan error, 1)
		errCh <- err
		close(errCh)
		return nil, errCh
	}

	billCh := make(chan views.VExportingBillsV3)
	errCh := make(chan error, 1)

	go func() {
		defer rows.Close()
		defer close(billCh)
		defer close(errCh)

		bCount := int64(0)
		log.Println("start streaming...")
		startTime := time.Now()
		for rows.Next() {
			var bill views.VExportingBillsV3
			if err := rows.Scan(
				&bill.ID, &bill.BillRunID, &bill.BillNumber, &bill.InvoiceDate, &bill.BillPeriodFrom, &bill.BillPeriodTo,
				&bill.AccountID, &bill.Amount, &bill.AmountNet, &bill.AccountBrand, &bill.BillMedia, &bill.CustomerID,
				&bill.CustomerNumber, &bill.CustomerClient, &bill.CustomerLanguage, &bill.InvoiceReference,
				&bill.SalutationString, &bill.FirstName, &bill.LastName, &bill.Street, &bill.HouseNumber, &bill.FreeText,
				&bill.ZIP, &bill.City, &bill.Alpha2, &bill.AdditionalAddress, &bill.Email, &bill.IBAN, &bill.MandateID,
				&bill.ProductDetails, &bill.BillType, &bill.HistoricBill, &bill.BillFrequency, &bill.BrandID,
				&bill.Billmanager, &bill.HasEVN, &bill.VatIDNumber, &bill.OtherData,
			); err != nil {
				errCh <- err
				return
			}

			billCh <- bill
			atomic.AddInt64(&bCount, 1)
		}
		endTime := time.Now()
		executionTime := endTime.Sub(startTime)
		log.Printf("got %v entities! in %v\n", bCount, executionTime)
	}()

	return billCh, errCh
}
