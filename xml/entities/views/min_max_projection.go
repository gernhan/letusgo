package views

type MinMaxProjection struct {
	Max   int64 `db:"max"`
	Min   int64 `db:"min"`
	Total int64 `db:"total"`
}
