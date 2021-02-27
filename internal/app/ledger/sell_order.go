package ledger

type SellOrder struct {
	tableName struct{} `pg:"ledger.sell_orders"`
	ID        uint64
	InvoiceID uint64   `pg:",notnull"`
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
	var sellOrder *SellOrder
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
		Financed:  false,
		InvoiceID: invoice.ID,
		Invoice:   invoice,
	}

	_, err = tx.Model(sellOrder).
		Returning("*").
		Insert()

	if err != nil {
		return sellOrder, err
	}

	return sellOrder, tx.Commit()
}

//
//func matchingAlgorithm(bid *Bid) error {
//
//	invoiceDiscount := calcDiscount(float64(bid.Invoice.FaceValue), float64(bid.Invoice.NeededValue))
//
//	if bid.ProfitPercentage > invoiceDiscount {
//		err := rejectBid(bid)
//		if err != nil {
//			return err
//		}
//		return errors.New("bid discount is bigger than invoice discount")
//	}
//
//	return nil
//}

// change bid status, return money to investor
func rejectBid(bid *Bid) error {
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer txCloseLog(tx)

	//bid.Investor.Balance += bid.InvestmentValue

	bid.Status = "rejected"
	
	_, err = tx.Model(bid).Update()

	if err != nil {
		return err
	}

	_, err = tx.Model(bid.Investor).Update()

	if err != nil {
		return err
	}

	return tx.Commit()
}