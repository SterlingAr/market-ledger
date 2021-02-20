package ledger

import "testing"

func TestNewInvestor(t *testing.T) {
	cleanDB()
	investor := &Investor{
		Name: "investor-1",
		Balance: 1000,
	}

	err := newInvestor(investor)

	if err != nil {
		t.Error(err)
	}

	inv, err := getInvestor(investor.Name)

	if err != nil {
		t.Error(err)
	}

	if inv.Name != investor.Name {
		t.Errorf("expected name =%v, actual name =%v", investor.Name, inv.Name)
	}

}

var investors = []*Investor{
	&Investor{
		Balance:   5000,
		Name:      "investor-1",
	},
}

func TestNewBid(t *testing.T) {
	cleanDB()
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		FaceValue: 1000,
		NeededValue: 900,
	}

	err = sellInvoice(issuer, invoice)

	if err != nil {
		t.Error(err)
	}

	investor := &Investor{
		Name: "investor-1",
		Balance: 1000,
	}

	err = newInvestor(investor)

	if err != nil {
		t.Error(err)
	}

	bid := &Bid{
		InvestmentValue : 450,
		ExpectedProfit : 500,
	}

	err = investor.newBid(invoice, bid)

	if err != nil {
		t.Error(err)
	}

	bids, err := getInvestorBids(investor)

	if err != nil {
		t.Error(err)
	}

	if len(bids) == 0 {
		t.Error("expected at least 1 bid, got none")
	}

	if bids[0].Investor.ID != investor.ID {
		t.Errorf("expected investor id %v, instead got %v", investor.ID, bids[0].Investor.ID)
	}

	if bids[0].Invoice.ID != invoice.ID {
		t.Errorf("expected invoice id %v, instead got %v", invoice.ID, bids[0].Invoice.ID)
	}

	if bids[0].InvestmentValue != bid.InvestmentValue {
		t.Errorf("expected investment value %v, instead got %v", bid.InvestmentValue, bids[0].InvestmentValue)
	}

	if bids[0].ExpectedProfit != bid.ExpectedProfit {
		t.Errorf("expected profit value %v, instead got %v", bid.ExpectedProfit, bids[0].ExpectedProfit)
	}

	if bids[0].Position != 1 {
		t.Errorf("expected position %v, instead got %v", 1, bids[0].Position)
	}

}


func TestInsufficientBalance(t *testing.T) {

}

func TestMismatchedDiscount(t *testing.T) {
	
}

func TestDiscountCalculator(t *testing.T) {
	
}