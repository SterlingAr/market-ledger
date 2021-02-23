package ledger

import "errors"

type SellOrder struct {
	tableName struct{} `pg:"ledger.sell_orders"`
	ID        uint64
	InvoiceID uint64	`pg:",notnull"`
	Invoice   *Invoice `pg:"rel:has-one"`
	Financed  bool
}

func getSellOrder(id uint64) (*SellOrder, error) {
	var so SellOrder
	err := db.Model(&so).
		Where("invoice_id = ?", id).
		Select()
	return &so, err
}

func newSellOrder(invoice *Invoice) (*SellOrder, error) {
	var sellOrder * SellOrder
	tx, err := db.Begin()

	if err != nil {
		return sellOrder, err
	}

	defer txCloseLog(tx)

	err = tx.Model(invoice).
		Where("id = ?", invoice.ID).
		Select()

	if err != nil {
		return sellOrder, err
	}

	sellOrder = &SellOrder{
		Financed: false,
		InvoiceID: invoice.ID,
		Invoice: invoice,
	}

	_, err = tx.Model(sellOrder).
		Returning("*").
		Insert()

	if err != nil {
		return sellOrder, err
	}

	return sellOrder, tx.Commit()
}

func matchingAlgorithm(invoice *Invoice, i *Investor, bid *Bid) error {

	invoiceDiscount := calcDiscount(float64(invoice.FaceValue), float64(invoice.NeededValue))

	if bid.ProfitPercentage > invoiceDiscount {
		return errors.New("bid discount is bigger than invoice discount")
	}

	return nil
}