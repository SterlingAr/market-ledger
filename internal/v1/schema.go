package v1

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/logger"
)

func seedDB() error {
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)

	if err != nil {
		return err
	}

	var investors = make(map[string]*Investor)
	for i := 1; i < 5; i++ {
		name := fmt.Sprintf("investor-%v", i)
		investor := &Investor{
			Balance: 1000,
			Name:    name,
		}

		err := newInvestor(investor)
		if err != nil {
			return err
		}
		investors[name] = investor
	}
	return nil
}

func cleanDB() {
	err := deleteTables()
	if err != nil {
		logger.Fatal(err)
	}
	err = createTables()
	if err != nil {
		logger.Fatal(err)
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
