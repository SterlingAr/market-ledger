package v1

import "github.com/google/logger"

type Ledger struct {
	Entries []Entry
}

type Entry struct {
	InvestorName    string
	InvoiceName     string
	InvestedBalance float64
	ReservedBalance float64
	ExpectedProfit  float64
}

func getLedger() Ledger {
	var (
		sellOrders []*SellOrder
		ledger     Ledger
	)

	err := db.Model(&sellOrders).Relation("Invoice").Select()

	if err != nil {
		logger.Error(err)
	}

	for _, so := range sellOrders {
		bids, err := getInvoiceBids(so.Invoice)

		if err != nil {
			logger.Error(err)
			continue
		}

		for _, bid := range bids {
			entry := Entry{
				InvestedBalance:  bid.InvestmentValue,
				ReservedBalance: bid.ReservedBalance,
				ExpectedProfit: bid.InvestmentValue - bid.ReservedBalance,
				InvestorName:   bid.Investor.Name,
				InvoiceName:    bid.Invoice.Name,
			}

			ledger.Entries = append(ledger.Entries, entry)
		}
	}
	return ledger
}