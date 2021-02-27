package ledger

import "testing"

func TestNewInvestor(t *testing.T) {
	cleanDB()
	investor := &Investor{
		Name:    "investor-1",
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

func TestSuccessfulNewBid(t *testing.T) {
	cleanDB()
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		FaceValue:   1000,
		NeededValue: 900,
	}

	err = sellInvoice(issuer, invoice)

	if err != nil {
		t.Error(err)
	}

	investor := &Investor{
		Name:    "investor-1",
		Balance: 1000,
	}

	err = newInvestor(investor)

	if err != nil {
		t.Error(err)
	}

	bid := &Bid{
		InvestmentValue:  450,
		Discount: 10,
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

	if bids[0].Discount != bid.Discount {
		t.Errorf("expected profit value %v, instead got %v", bid.Discount, bids[0].Discount)
	}

	if bids[0].Position != 1 {
		t.Errorf("expected position %v, instead got %v", 1, bids[0].Position)
	}

	inv, err := getInvestor(investor.Name)

	if err != nil {
		t.Error(err)
	}

	if inv.Balance != 550 {
		t.Errorf("expected balance %v, actual balance %v", 550, inv.Balance)
	}
}

func TestInsufficientBalance(t *testing.T) {
	cleanDB()
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		FaceValue:   1000,
		NeededValue: 900,
	}

	err = sellInvoice(issuer, invoice)

	if err != nil {
		t.Error(err)
	}

	investor := &Investor{
		Name:    "investor-1",
		Balance: 1000,
	}

	err = newInvestor(investor)

	if err != nil {
		t.Error(err)
	}

	bid := &Bid{
		InvestmentValue:  5450,
		Discount: 5,
	}

	err = investor.newBid(invoice, bid)

	expectedErrorString := "insufficient balance"

	if err == nil {
		t.Error("expected error, got none")
	} else {
		if err.Error() != expectedErrorString {
			t.Errorf("expected error string: %v", expectedErrorString)
		}
	}
}

func TestExcessProfit(t *testing.T) {
	cleanDB()
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		FaceValue:   1000,
		NeededValue: 900,
	}

	err = sellInvoice(issuer, invoice)

	if err != nil {
		t.Error(err)
	}

	investor := &Investor{
		Name:    "investor-1",
		Balance: 1000,
	}

	err = newInvestor(investor)

	if err != nil {
		t.Error(err)
	}

	bid := &Bid{
		InvestmentValue:  450,
		Discount: 15,
	}

	err = investor.newBid(invoice, bid)

	expectedErrorString := "bid discount is bigger than invoice discount"
	if err == nil {
		t.Error("expected error, got none")
	} else {
		if err.Error() != expectedErrorString {
			t.Errorf("expected error string: %v", expectedErrorString)
		}
	}
}

func TestDiscountCalculator(t *testing.T) {
	discount := calcDiscount(900, 1000)

	if discount != 10 {
		t.Errorf("expected discount value %v, actual discount value %v", 10, discount)
	}

	discount = calcDiscount(1000, 900)

	if discount != 10 {
		t.Errorf("expected discount value %v, actual discount value %v", 10, discount)
	}

	discount = calcDiscount(889, 1000)

	if discount != 11.1 {
		t.Errorf("expected discount value %v, actual discount value %v", 10, discount)
	}

	discount = calcDiscount(1000, 1000)

	if discount != 0 {
		t.Errorf("expected discount value %v, actual discount value %v", 0, discount)
	}
}
