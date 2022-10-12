package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/brand"
	repositories "catalyst/src/domain/repositories"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	CreateBrand = `INSERT INTO jam_tangan.brands (name, created_at) VALUES ($1, now())`
)

var statement PreparedStatement

type PreparedStatement struct {
	createBrand *sqlx.Stmt
}

type BrandsRepo struct {
	Connection *sqlx.DB
}

func NewBrandsRepository(db *sqlx.DB) repositories.BrandRepository {
	repo := &BrandsRepo{db}
	InitPreparedStatement(repo)
	return repo
}

func (p *BrandsRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *BrandsRepo) {
	statement = PreparedStatement{
		createBrand: m.Preparex(CreateBrand),
	}
}

func (p *BrandsRepo) CreateBrand(dataBrand *dto.CreateBrandReqDTO) error {

	_, err := statement.createBrand.Exec(dataBrand.Name)

	if err != nil {
		log.Println("Failed Query CreateBrand : ", err.Error())
		return err
	}

	return nil
}
