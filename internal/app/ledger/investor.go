package ledger

import "errors"

type Investor struct {
	tableName struct{} `pg:"ledger.investor"`
	ID        uint64
	Balance   float64
	Name      string
	Bids      []*Bid `pg:"rel:has-many"`
}

type Bid struct {
	tableName struct{} `pg:"ledger.bids"`
	//ID              uint64
	Position        int `pg:",pk"`
	InvestmentValue float64
	ReservedBalance float64
	// I am not sure what bid discount is
	// as a bidder I bid 500 euros
	// I have a bid discount of 10%
	// that means I won't be paying 500 euros? but only 450
	// if the bid goes through, my balance is reduced by 450 euros
	Discount        float64
	InvestorID      uint64    `pg:",pk"`
	Investor        *Investor `pg:"rel:has-one"`
	InvoiceID       uint64    `pg:",pk"`
	Invoice         *Invoice  `pg:"rel:has-one"`
	Status          string
}

func (i *Investor) newBid(invoice *Invoice, bid *Bid) error {
	if bid.InvestmentValue > i.Balance {
		return errors.New("insufficient balance")
	}

	invoiceDiscount := calcDiscount(invoice.FaceValue, invoice.NeededValue)

	if bid.Discount > invoiceDiscount {
		err := rejectBid(bid)
		if err != nil {
			return err
		}
	}

	_, err := getSellOrder(invoice.ID)

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

func getInvestor(name string) (*Investor, error) {
	var investor Investor

	err := db.Model(&investor).
		Where("name = ?", name).
		Select()

	return &investor, err
}

func getInvestorBids(investor *Investor) ([]*Bid, error) {
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

	p = (z / x) * 100

	return p
}
