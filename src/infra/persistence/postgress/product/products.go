package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/product"
	repositories "catalyst/src/domain/repositories"
	models "catalyst/src/infra/models"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	CreateProduct       = `INSERT INTO jam_tangan.products (name, price, brand_id, created_at) VALUES ($1, $2, $3, now())`
	GetSingleProduct    = `SELECT id, name, price, brand_id from jam_tangan.products where id = $1 and deleted_at ISNULL`
	GetProductByBrandId = `SELECT id, name, price, brand_id from jam_tangan.products where brand_id = $1 and deleted_at ISNULL`
)

var statement PreparedStatement

type PreparedStatement struct {
	createProduct       *sqlx.Stmt
	getSingleProduct    *sqlx.Stmt
	getProductByBrandId *sqlx.Stmt
}

type ProductsRepo struct {
	Connection *sqlx.DB
}

func NewProductsRepository(db *sqlx.DB) repositories.ProductRepository {
	repo := &ProductsRepo{db}
	InitPreparedStatement(repo)
	return repo
}

func (p *ProductsRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *ProductsRepo) {
	statement = PreparedStatement{
		createProduct:       m.Preparex(CreateProduct),
		getSingleProduct:    m.Preparex(GetSingleProduct),
		getProductByBrandId: m.Preparex(GetProductByBrandId),
	}
}

func (p *ProductsRepo) CreateProduct(dataProduct *dto.CreateProductReqDTO) error {

	_, err := statement.createProduct.Exec(dataProduct.Name, dataProduct.Price, dataProduct.BrandId)

	if err != nil {
		log.Println("Failed Query CreateProduct : ", err.Error())
		return err
	}

	return nil
}

func (p *ProductsRepo) GetSingleProduct(dataProduct *dto.GetProductReqDTO) (*models.Products, error) {
	var data []*models.Products

	err := statement.getSingleProduct.Select(&data, dataProduct.ID)

	if err != nil {
		log.Println("Failed Query GetSingleProduct : ", err.Error())
		return nil, err
	}

	return data[0], nil
}

func (p *ProductsRepo) GetProductByBrandId(dataProduct *dto.GetProductReqDTO) ([]*models.Products, error) {
	var data []*models.Products

	err := statement.getProductByBrandId.Select(&data, dataProduct.ID)

	if err != nil {
		log.Println("Failed Query GetProductByBrandId : ", err.Error())
		return nil, err
	}

	return data, nil
}
