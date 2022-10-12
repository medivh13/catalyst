package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : catalyst
 */

type OrderReqModel struct {
	OrderID int64 `db:"id"`
	// OrderCode string `db:"order_code"`
}

type Orders struct {
	ID         int64   `db:"id"`
	OrderCode  string  `db:"order_code"`
	Total      float64 `db:"total"`
	Product    string  `db:"product"`
	Quantity   int64   `db:"quantity"`
	TotalPrice float64 `db:"total_price"`
}
