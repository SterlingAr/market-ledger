package ledger

import (
	"fmt"
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

	entries, err := getLedger()

	// investor check
	if entries[0].InvestorName != "investor-1" {
		t.Errorf("expected investor name %v, actual %v", "investor-1", entries[0].InvestorName)
	}

	if entries[1].InvestorName != "investor-2" {
		t.Errorf("expected investor name %v, actual %v", "investor-2", entries[0].InvestorName)
	}

	if entries[2].InvestorName != "investor-4" {
		t.Errorf("expected investor name %v, actual %v", "investor-4", entries[0].InvestorName)
	}

	// invoice check
	if entries[0].InvoiceName != "invoice-1" {
		t.Errorf("expected invoice name %v, actual %v", "invoice-1", entries[0].InvoiceName)
	}

	if entries[1].InvoiceName != "invoice-2" {
		t.Errorf("expected invoice name %v, actual %v", "invoice-2", entries[0].InvoiceName)
	}

	if entries[2].InvoiceName != "invoice-4" {
		t.Errorf("expected invoice name %v, actual %v", "invoice-4", entries[0].InvoiceName)
	}

	// reserved value
	if entries[0].InvestedValue != 450 {
		t.Errorf("expected InvestedValue %v, actual %v", 450, entries[0].InvestedValue)
	}

	if entries[1].InvestedValue != 270 {
		t.Errorf("expected InvestedValue %v, actual %v", 270, entries[0].InvestedValue)
	}

	if entries[2].InvestedValue != 190 {
		t.Errorf("expected InvestedValue %v, actual %v", 190, entries[0].InvestedValue)
	}

	// expected profit
	if entries[0].ExpectedProfit != 50 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 50, entries[0].ExpectedProfit)
	}

	if entries[1].ExpectedProfit != 30 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 30, entries[0].ExpectedProfit)
	}

	if entries[2].ExpectedProfit != 10 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 10, entries[0].ExpectedProfit)
	}

}

func newSellOrderTestData() error {
	issuer := &Issuer{
		Name:      "party-A",
	}

	err := newIssuer(issuer)

	if err != nil {
		return err
	}

	var investors = make(map[string]*Investor)
	for i:= 1; i < 5; i++ {
		name := fmt.Sprintf("investor-%v", i)
		investor := &Investor{
			Balance:   1000,
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