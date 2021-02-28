package v1

import (
	"fmt"
	"testing"
)

func TestGetLedger(t *testing.T) {
	cleanDB()
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
	if ledger.Entries[0].InvestorName != "investor-1" {
		t.Errorf("expected investor name %v, actual %v", "investor-1", ledger.Entries[0].InvestorName)
	}

	if ledger.Entries[1].InvestorName != "investor-2" {
		t.Errorf("expected investor name %v, actual %v", "investor-2", ledger.Entries[1].InvestorName)
	}

	if ledger.Entries[2].InvestorName != "investor-4" {
		t.Errorf("expected investor name %v, actual %v", "investor-4", ledger.Entries[2].InvestorName)
	}

	// invoice check
	if ledger.Entries[0].InvoiceName != "invoice-1" {
		t.Errorf("expected invoice name %v, actual %v", "invoice-1", ledger.Entries[0].InvoiceName)
	}

	if ledger.Entries[1].InvoiceName != "invoice-1" {
		t.Errorf("expected invoice name %v, actual %v", "invoice-1", ledger.Entries[1].InvoiceName)
	}

	if ledger.Entries[2].InvoiceName != "invoice-1" {
		t.Errorf("expected invoice name %v, actual %v", "invoice-1", ledger.Entries[2].InvoiceName)
	}

	// invested value
	if ledger.Entries[0].InvestedBalance != 500 {
		t.Errorf("expected InvestedValue %v, actual %v", 500, ledger.Entries[0].InvestedBalance)
	}

	if ledger.Entries[1].InvestedBalance != 300 {
		t.Errorf("expected InvestedValue %v, actual %v", 300, ledger.Entries[1].InvestedBalance)
	}

	if ledger.Entries[2].InvestedBalance != 100 {
		t.Errorf("expected InvestedValue %v, actual %v", 100, ledger.Entries[2].InvestedBalance)
	}

	if ledger.Entries[0].ReservedBalance != 450 {
		t.Errorf("expected InvestedValue %v, actual %v", 450, ledger.Entries[0].ReservedBalance)
	}

	// reserved value (invested value - discount)
	if ledger.Entries[1].ReservedBalance != 270 {
		t.Errorf("expected InvestedValue %v, actual %v", 270, ledger.Entries[1].ReservedBalance)
	}

	if ledger.Entries[2].ReservedBalance != 95 {
		t.Errorf("expected InvestedValue %v, actual %v", 95, ledger.Entries[2].ReservedBalance)
	}

	// expected profit
	if ledger.Entries[0].ExpectedProfit != 50 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 50, ledger.Entries[0].ExpectedProfit)
	}

	if ledger.Entries[1].ExpectedProfit != 30 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 30, ledger.Entries[1].ExpectedProfit)
	}

	if ledger.Entries[2].ExpectedProfit != 5 {
		t.Errorf("expected ExpectedProfit %v, actual %v", 10, ledger.Entries[2].ExpectedProfit)
	}
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
		Name:        "invoice-1",
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
		InvestmentValue: 500,
		Discount:        10,
	})

	if err != nil {
		return err
	}

	err = investors["investor-2"].newBid(invoice, &Bid{
		InvestmentValue: 300,
		Discount:        10,
	})

	if err != nil {
		return err
	}

	err = investors["investor-3"].newBid(invoice, &Bid{
		InvestmentValue: 200,
		Discount:        14.29,
	}) // for this instance, its expected for the bid to be rejected

	if err == nil {
		return err
	}

	err = investors["investor-4"].newBid(invoice, &Bid{
		InvestmentValue: 300,
		Discount:        5,
	})

	if err != nil {
		return err
	}

	// 1 new sell order
	// 4 new bids, one for each investor
	return nil
}
