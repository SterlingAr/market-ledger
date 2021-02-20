package ledger

import (
	"context"
	"fmt"
	"github.com/SterlingAr/market-ledger/internal/pkg/database"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/google/logger"
	"os"

	"testing"
)

var db *pg.DB
func TestMain(m *testing.M) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	db = database.Connection(ctx, database.ConnectionParams{
		User:     "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     "5432",
		Database: "shine",
		Schema:   "world",
	})
	err := database.CreateSchema(db, "market")
	if err != nil {
		logger.Fatal(err)
	}
	cleanDB()
	os.Exit(m.Run())
}

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
	createTx, err := db.Begin()
	if err != nil {
		return err
	}
	defer createTx.Close()
	for _, model := range []interface{}{
		(*Issuer)(nil),
	} {
		err := createTx.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return fmt.Errorf("%v, %v", err, createTx.Rollback())
		}
	}
	return createTx.Commit()
}

func deleteTables() error {
	deleteTx, err := db.Begin()
	if err != nil {
		return err
	}
	defer deleteTx.Close()

	for _, model := range []interface{}{
		(*Issuer)(nil),
	} {
		err := deleteTx.DropTable(model, &orm.DropTableOptions{
			IfExists: true,
			Cascade:  true,
		})
		if err != nil {
			return fmt.Errorf("%v, %v", err, deleteTx.Rollback())
		}
	}
	return deleteTx.Commit()
}


// only a representation
type LedgerJSON struct{
	//issuers []*Issuer
}

type Issuer struct {
	tableName        struct{} `pg:"market.issuers"`
	ID               uint64
	Name string  `pg:",notnull,unique"`
	Invoices []*Invoice
}

type Invoice struct {
	tableName        struct{} `pg:"market.invoices"`
	ID               uint64
	FaceValue int
	IssuerID uint64
	Issuer *Issuer
}

type SellOrder struct {}

type Investor struct{
	Balance int
}

type Bid struct {}

func TestnewIssuer(t *testing.T)  {

	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	iss, err := getIssuer(issuer.Name)
	if err != nil {
		t.Error(err)
	}

	if iss.Name != issuer.Name {
		t.Errorf("expected name =%v, actual name =%v", issuer.Name, iss.Name)
	}
}

func TestIssueInvoice(t *testing.T) {
	// An issuer party-A wants to finance an invoice with number invoice-1. The face value
	//of the invoice is €1000.
	//We need to record in the ledger that party-A has a €1000 invoice invoice-1 that
	//should be financed;

	issuer := &Issuer{
		Name: "party-A",
	}

	invoice := &Invoice{
		FaceValue: 1000,
	}

	err := sellInvoice(issuer, invoice)

	if err != nil {
		t.Errorf("expected to sell invoice, got error %v", err)
	}


	iss, err := getIssuer(issuer.Name)

	if err != nil {
		t.Errorf("expected to find issuer %v, got error %v", issuer.Name, err)
	}

	var ivc []Invoice
	ivc, err = getInvoices(iss, 900, 1000)

	if len(ivc) == 0 {
		t.Errorf("expected invoices = %v, actual invoices =%v", 1, len(ivc))
	}

}


func txCloseLog(tx * pg.Tx)  {
	err := tx.Close()
	if err != nil {
		logger.Error(err)
	}
}

func newIssuer(issuer *Issuer) error {
	newTx, err := db.Begin()

	if err != nil {
		return err
	}

	defer txCloseLog(newTx)

	_, err = newTx.Model(issuer).Returning("*").Insert()
	if err != nil {
		return  err
	}

	return newTx.Commit()
}

func getIssuer(name string) (*Issuer, error) {
	var iss * Issuer
	return iss, nil
}

func sellInvoice(issuer * Issuer, invoice * Invoice) error {
	// if issuer exists
	//
	return nil
}

func getInvoices(iss * Issuer, min int, max int) ([]Invoice, error) {
	var ivcs []Invoice
	return ivcs, nil
}
