package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : catalyst
 */

import "time"

type Products struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Price     float64   `db:"price"`
	BrandID   int64     `db:"brand_id"`
	BrandName string    `db:"brand_name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
