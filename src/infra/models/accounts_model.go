package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : catalyst
 */

type Accounts struct {
	ID            int64   `gorm:"id"`
	AccountNumber string  `gorm:"account_number"`
	CustomerName  string  `gorm:"customer_name"`
	Balance       float64 `gorm:"balance"`
}
