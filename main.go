package main

import (
	"context"
	"database/sql"

	usecases "catalyst/src/app/use_cases"

	"catalyst/src/infra/config"
	postgres "catalyst/src/infra/persistence/postgress"
	pgBrand "catalyst/src/infra/persistence/postgress/brand"
	pgOrder "catalyst/src/infra/persistence/postgress/order"
	pgProduct "catalyst/src/infra/persistence/postgress/product"

	"catalyst/src/interface/rest"

	ms_log "catalyst/src/infra/log"

	brandUc "catalyst/src/app/use_cases/brand"
	orderUc "catalyst/src/app/use_cases/order"
	productUc "catalyst/src/app/use_cases/product"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

//reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	brandRepository := pgBrand.NewBrandsRepository(postgresdb.Conn)
	productRepository := pgProduct.NewProductsRepository(postgresdb.Conn)
	orderRepository := pgOrder.NewOrderRepository(postgresdb.Conn)
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{

			BrandUseCase:   brandUc.NewBrandUseCase(brandRepository),
			ProductUseCase: productUc.NewProductUseCase(productRepository),
			OrderUseCase:   orderUc.NewOrderUseCase(orderRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
