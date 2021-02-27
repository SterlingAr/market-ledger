package ledger

import "github.com/google/logger"

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

func matchingAlgorithm() error {
	var (
		sellOrders []*SellOrder
	)

	err := db.Model(&sellOrders).Relation("Invoice").Select()

	if err != nil {
		return err
	}

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer txCloseLog(tx)

soLoop:
	for _, so := range sellOrders {
		bids, err := getInvoiceBids(so.Invoice)

		if err != nil {
			return err
		}

		var total float64 = 0

		for _, bid := range bids {
			var surplus float64

			if so.Financed {
				break soLoop
			}

			total += bid.InvestmentValue

			reservedBalance := investmentDiscount(bid.InvestmentValue, bid.Discount)

			bid.Investor.Balance -= reservedBalance

			// persist how much party-Y owes investor-X, which should be the real investment value (without the discount)
			bid.ReservedBalance = reservedBalance

			_, err := tx.Model(bid).Where("position = ?", bid.Position).Update()
			if err != nil {
				return err
			}

			_, err = tx.Model(bid.Investor).WherePK().Update()
			if err != nil {
				return err
			}

			if total >= so.Invoice.NeededValue {
				if total == so.Invoice.NeededValue {

				} else {

					surplus =  total - so.Invoice.NeededValue

					bid.InvestmentValue -= surplus

					reservedBalance = investmentDiscount(bid.InvestmentValue, bid.Discount)
					bid.ReservedBalance = reservedBalance

					bid.Investor.Balance += investmentDiscount(surplus, bid.Discount)

					_, err := tx.Model(bid).Where("position = ?", bid.Position).Update()
					if err != nil {
						return err
					}

					_, err = tx.Model(bid.Investor).WherePK().Update()
					if err != nil {
						return err
					}
				}

				so.Financed = true
				_, err = tx.Model(so).WherePK().Update()
				if err != nil {
					logger.Error(err)
				}

				break soLoop
			}
		}
	}
	return tx.Commit()
}

// apply discount to the investment Value
func investmentDiscount(value float64, discount float64) float64 {
	return value - (value * (discount/100))
}
