package ledger

import (
	"fmt"
	"github.com/google/logger"
	"testing"
)

func TestGetLedger(t *testing.T) {
	// for all sell orders, before running the matching algorithm for that sale order
	// print something similar:

	//party-A has a €1000 invoice invoice-1 that should be financed;
	//● investor-1 has €450 reserved for the purchase of invoice-1;
	//● investor-2 has €270 reserved for the purchase of invoice-1;
	//● investor-4 has €285 reserved for the purchase of invoice-1;
	err := newSellOrderTestData()
	if err != nil {
		t.Error(err)
	}

	// entries in the ledger should be displayed in chronological order
	err = matchingAlgorithm()
	if err != nil {
		t.Error(err)
	}
	ledger := getLedger()

	// investor check
	//if ledger.Entries[0].InvestorName != "investor-1" {
	//	t.Errorf("expected investor name %v, actual %v", "investor-1", ledger.Entries[0].InvestorName)
	//}
	//
	//if ledger.Entries[1].InvestorName != "investor-2" {
	//	t.Errorf("expected investor name %v, actual %v", "investor-2", ledger.Entries[1].InvestorName)
	//}
	//
	//if ledger.Entries[2].InvestorName != "investor-4" {
	//	t.Errorf("expected investor name %v, actual %v", "investor-4", ledger.Entries[2].InvestorName)
	//}
	//
	//// invoice check
	//if ledger.Entries[0].InvoiceName != "invoice-1" {
	//	t.Errorf("expected invoice name %v, actual %v", "invoice-1", ledger.Entries[0].InvoiceName)
	//}
	//
	//if ledger.Entries[1].InvoiceName != "invoice-1" {
	//	t.Errorf("expected invoice name %v, actual %v", "invoice-1", ledger.Entries[1].InvoiceName)
	//}
	//
	//if ledger.Entries[2].InvoiceName != "invoice-1" {
	//	t.Errorf("expected invoice name %v, actual %v", "invoice-1", ledger.Entries[2].InvoiceName)
	//}

	// reserved value
	if ledger.Entries[0].InvestedValue != 450 {
		t.Errorf("expected InvestedValue %v, actual %v", 450, ledger.Entries[0].InvestedValue)
	}

	if ledger.Entries[1].InvestedValue != 270 {
		t.Errorf("expected InvestedValue %v, actual %v", 270, ledger.Entries[1].InvestedValue)
	}

	if ledger.Entries[2].InvestedValue != 190 {
		t.Errorf("expected InvestedValue %v, actual %v", 190, ledger.Entries[2].InvestedValue)
	}

	// expected profit
	if ledger.Entries[0].ExpectedProfit != 50 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 50, ledger.Entries[0].ExpectedProfit)
	}

	if ledger.Entries[1].ExpectedProfit != 30 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 30, ledger.Entries[1].ExpectedProfit)
	}

	if ledger.Entries[2].ExpectedProfit != 10 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 10, ledger.Entries[2].ExpectedProfit)
	}

}

func getLedger() Ledger {
	var (
		sellOrders []*SellOrder
		ledger Ledger
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
			entry := LedgerEntry{
				InvestedValue:  bid.InvestmentValue,
				ExpectedProfit: calcProfit(bid.InvestmentValue, bid.ProfitPercentage),
				InvestorName:   bid.Investor.Name,
				InvoiceName:    bid.Invoice.Name,
			}

			ledger.Entries = append(ledger.Entries, entry)
		}
	}
	return ledger
}

func calcProfit(investedValue float64, profitPercentage float64) float64 {
	return investedValue * (profitPercentage / 100)
}

type Ledger struct {
	Entries []LedgerEntry
}

type LedgerEntry struct {
	InvestedValue  float64
	ExpectedProfit float64
	InvestorName   string
	InvoiceName    string
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

			if total >= so.Invoice.NeededValue {
				if total == so.Invoice.NeededValue {

				} else {

					surplus =  total - so.Invoice.NeededValue

					bid.InvestmentValue -= surplus
					bid.Investor.Balance += surplus

					_, err := tx.Model(bid).Where("position = ?", bid.Position).Update()
					if err != nil {
						logger.Error(err)
					}

					_, err = tx.Model(bid.Investor).WherePK().Update()
					if err != nil {
						logger.Error(err)
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

func newSellOrderTestData() error {
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)

	if err != nil {
		return err
	}

	var investors = make(map[string]*Investor)
	for i := 1; i < 5; i++ {
		name := fmt.Sprintf("investor-%v", i)
		investor := &Investor{
			Balance: 1000,
			Name:    name,
		}

		err := newInvestor(investor)
		if err != nil {
			return err
		}
		investors[name] = investor
	}

	invoice := &Invoice{
		FaceValue:   1000,
		NeededValue: 900,
	}

	err = sellInvoice(issuer, invoice)

	if err != nil {
		return err
	}

	_, err = newSellOrder(invoice)

	if err != nil {
		return err
	}

	err = investors["investor-1"].newBid(invoice, &Bid{
		InvestmentValue:  450,
		ProfitPercentage: 10,
	})

	if err != nil {
		return err
	}

	err = investors["investor-2"].newBid(invoice, &Bid{
		InvestmentValue:  270,
		ProfitPercentage: 10,
	})

	if err != nil {
		return err
	}

	err = investors["investor-3"].newBid(invoice, &Bid{
		InvestmentValue:  175,
		ProfitPercentage: 14.29,
	}) // for this instance, its expected for the bid to be rejected

	if err == nil {
		return err
	}

	err = investors["investor-4"].newBid(invoice, &Bid{
		InvestmentValue:  285,
		ProfitPercentage: 10,
	})
	if err != nil {
		return err
	}

	// 1 new sell order
	// 4 new bids, one for each investor
	return nil
}
