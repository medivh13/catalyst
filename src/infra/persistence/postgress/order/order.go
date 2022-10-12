package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/order"
	dtoDetail "catalyst/src/app/dtos/order_detail"
	repositories "catalyst/src/domain/repositories"
	models "catalyst/src/infra/models"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	CreateOrder       = `INSERT INTO jam_tangan.orders (order_code, total, created_at) VALUES ($1, $2, now()) Returning id`
	CreateOrderDetail = `INSERT INTO jam_tangan.order_details (order_id, product_id, quantity, created_at) VALUES ($1, $2, $3, now())`
	GetOrderByID      = `SELECT o.id, o.order_code, o.total, p.name as product, od.quantity, (od.quantity * p.price) as total_price
	from jam_tangan.orders o JOIN jam_tangan.order_details od
	ON o.id = od.order_id
	JOIN jam_tangan.products p ON p.id = od.product_id
	Where o.id = $1`
)

var statement PreparedStatement

type PreparedStatement struct {
	getOrderByID *sqlx.Stmt
}

type OrderRepo struct {
	Connection *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) repositories.OrderRepository {
	repo := &OrderRepo{db}
	InitPreparedStatement(repo)
	return repo
}

func (p *OrderRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *OrderRepo) {
	statement = PreparedStatement{
		getOrderByID: m.Preparex(GetOrderByID),
	}
}

func (p *OrderRepo) CreateOrder(dataOrder *dto.CreateOrderReqDTO) (*models.OrderReqModel, error) {
	var total float64
	orderCode := uuid.New()
	var orderData models.OrderReqModel
	for _, val := range dataOrder.Data {
		total += (val.Price * float64(val.Quantity))
	}

	tx, err := p.Connection.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx CreateOrder : ", err.Error())
		return nil, err
	}

	err = tx.QueryRow(CreateOrder, orderCode, total).Scan(&orderData.OrderID)

	fmt.Println(orderData, "here")
	if err != nil {
		tx.Rollback()
		log.Println("Failed Query Create Order: ", err.Error())
		return nil, err
	}
	for _, val := range dataOrder.Data {
		_, err = tx.Exec(CreateOrderDetail, orderData.OrderID, val.ProductID, val.Quantity)
		if err != nil {
			tx.Rollback()
			log.Println("Failed Query Create Order Detail : ", err.Error())
			return nil, err
		}
	}
	dataResult := &models.OrderReqModel{
		OrderID: orderData.OrderID,
	}
	return dataResult, tx.Commit()
}

func (p *OrderRepo) GetOrderByID(dataOrder *dtoDetail.GetOrderReqDTO) ([]*models.Orders, error) {
	var data []*models.Orders

	err := statement.getOrderByID.Select(&data, dataOrder.OrderID)

	if err != nil {
		log.Println("Failed Query GetOrderByID : ", err.Error())
		return nil, err
	}

	return data, nil
}
