package ledger

type Invoice struct {
	tableName   struct{} `pg:"ledger.invoices"`
	ID          uint64
	Name        string `pg:",notnull,unique"`
	FaceValue   float64
	NeededValue float64
	IssuerID    uint64
	Issuer      *Issuer `pg:"rel:has-one"`
	Bids        []*Bid  `pg:"rel:has-many"`
}

func getInvoiceBids(invoice *Invoice) ([]*Bid, error) {
	var (
		bids []*Bid
	)

	err := db.Model(&bids).
		Where("invoice_id = ?", invoice.ID).
		Relation("Investor").
		Relation("Invoice").
		Select()

	return bids, err
}
