package v1

type Issuer struct {
	tableName struct{} `pg:"ledger.issuers"`
	ID        uint64
	Name      string     `pg:",notnull,unique"`
	Invoices  []*Invoice `pg:"rel:has-many"`
}

func newIssuer(issuer *Issuer) error {
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer txCloseLog(tx)

	_, err = tx.Model(issuer).
		Returning("*").
		Insert()

	if err != nil {
		return err
	}

	return tx.Commit()
}

func getIssuer(name string) (*Issuer, error) {
	var iss Issuer

	err := db.Model(&iss).
		Where("name = ?", name).
		Relation("Invoices").
		Select()

	return &iss, err
}

func sellInvoice(issuer *Issuer, invoice *Invoice) error {
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer txCloseLog(tx)

	invoice.Issuer = issuer
	invoice.IssuerID = issuer.ID
	// issuer.
	_, err = tx.Model(invoice).
		Returning("*").
		Insert()

	if err != nil {
		return err
	}

	return tx.Commit()
}
