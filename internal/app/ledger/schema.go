package ledger

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/logger"
)

var db *pg.DB

func cleanDB() {
	err := deleteTables()
	if err != nil {
		logger.Error(err)
	}
	err = createTables()
	if err != nil {
		logger.Error(err)
	}
}

func createTables() error {
	models := []interface{}{
		(*Issuer)(nil),
		(*Invoice)(nil),
		(*SellOrder)(nil),
		(*Investor)(nil),
		(*Bid)(nil),
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Fatal(err)
	}
	defer tx.Close()

	for _, model := range models {
		err := tx.Model(model).CreateTable(&orm.CreateTableOptions{
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func deleteTables() error {
	models := []interface{}{
		(*Issuer)(nil),
		(*Invoice)(nil),
		(*SellOrder)(nil),
		(*Investor)(nil),
		(*Bid)(nil),
	}
	tx, err := db.Begin()

	if err != nil {
		logger.Fatal(err)
	}

	defer tx.Close()

	for _, model := range models {
		err := tx.Model(model).DropTable(&orm.DropTableOptions{
			IfExists: true,
			Cascade:  true,
		})
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func txCloseLog(tx *pg.Tx) {
	err := tx.Close()
	if err != nil {
		logger.Error(err)
	}
}
