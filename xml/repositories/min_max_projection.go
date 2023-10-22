package repositories

import (
	"context"

	"github.com/gernhan/xml/entities/views"

	"github.com/gernhan/xml/db"
)

func FindMinMaxBillId(billRunID int64, status int) (views.MinMaxProjection, error) {
	pool := db.GetPool()
	query := `
		SELECT
			COALESCE(MAX(db.id), -1) AS max,
			COALESCE(MIN(db.id), -1) AS min,
			COUNT(db.id) AS total
		FROM
			d_billruns dbr
		JOIN
			d_bills db ON dbr.id = db.billrun
		WHERE
			dbr.id = $1
			and db.pdf_status = $2`

	var result views.MinMaxProjection
	err := pool.QueryRow(context.Background(), query, billRunID, status).Scan(&result.Max, &result.Min, &result.Total)
	if err != nil {
		return views.MinMaxProjection{}, err
	}

	return result, nil
}
