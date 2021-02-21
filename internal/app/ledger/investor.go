package ledger

import "errors"

type Investor struct {
	tableName struct{} `pg:"ledger.investor"`
	ID        uint64
	Balance   int
	Name      string
	Bids  []*Bid `pg:"rel:has-many"`
}

type Bid struct {
	tableName struct{} `pg:"ledger.bids"`
	//ID              uint64
	Position         int `pg:",pk"`
	InvestmentValue  int
	ProfitPercentage float64
	InvestorID       uint64    `pg:",pk"`
	Investor         *Investor `pg:"rel:has-one"`
	InvoiceID        uint64    `pg:",pk"`
	Invoice          *Invoice  `pg:"rel:has-one"`
}

func (i *Investor) newBid(invoice *Invoice, bid *Bid) error {

	err := matchingAlgorithm(invoice, i, bid)

	if err != nil {
		return err
	}

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer txCloseLog(tx)

	bid.Investor = i
	bid.InvestorID = i.ID

	bid.Invoice = invoice
	bid.InvoiceID = invoice.ID

	_, err = tx.Model(bid).
		Returning("*").
		Insert()

	if err != nil {
		return err
	}

	return tx.Commit()
}

func matchingAlgorithm(invoice *Invoice, i *Investor, bid *Bid) error {
	if bid.InvestmentValue > i.Balance {
		return errors.New("insufficient balance")
	}

	invoiceDiscount := calcDiscount(float64(invoice.FaceValue), float64(invoice.NeededValue))

	if bid.ProfitPercentage > invoiceDiscount {
		return errors.New("bid discount is bigger than invoice discount")
	}
	return nil
}

func getInvestor(name string) (*Investor, error) {
	var investor Investor

	err := db.Model(&investor).
		Where("name = ?", name).
		Select()

	return &investor, err
}

func getInvestorBids(investor * Investor) ([]*Bid, error) {
	var (

		bids []*Bid
	)

	err := db.Model(&bids).
		Where("investor_id = ?", investor.ID).
		Relation("Investor").
		Relation("Invoice").
		Select()

	return bids, err
}

func newInvestor(investor *Investor) error {
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer txCloseLog(tx)

	_, err = tx.Model(investor).
		Returning("*").
		Insert()

	if err != nil {
		return err
	}

	return tx.Commit()
}


func calcDiscount(i float64, i2 float64) float64 {
	var x, y, z, p float64

	if i > i2 {
		y = i2
		x = i
	} else {
		y = i
		x = i2
	}

	z = x - y

	p = (z/x) * 100

	return p
}