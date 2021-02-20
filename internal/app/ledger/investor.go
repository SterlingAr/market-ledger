package ledger

type Investor struct {
	tableName struct{} `pg:"ledger.investor"`
	ID        uint64
	Balance   int
	Name      string
	Bids  []*Bid `pg:"rel:has-many"`
}

type Bid struct {
	tableName       struct{}  `pg:"ledger.bids"`
	ID              uint64
	Position        int 	  `pg:",pk"`
	InvestmentValue int
	ExpectedProfit  int
	InvestorID      uint64    `pg:",pk"`
	Investor        *Investor `pg:"rel:has-one"`
	InvoiceID       uint64    `pg:",pk"`
	Invoice         *Invoice  `pg:"rel:has-one"`
}

func (i *Investor) newBid(invoice *Invoice, bid *Bid) error {
	// matchingAlgorithm test
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
