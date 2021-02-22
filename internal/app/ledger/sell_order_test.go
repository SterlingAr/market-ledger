package ledger

import "testing"

func TestNewSellOrder(t *testing.T) {
	// an existing invoice, make it available so that bids can be placed
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		FaceValue:   0,
		NeededValue: 0,
		IssuerID:    0,
		Issuer:      nil,
		Bids:        nil,
	}

	err = sellInvoice(issuer, invoice)
	if err != nil {
		t.Error(err)
	}

	err = newSellOrder(invoice)

	if err != nil {
		t.Error(err)
	}

	so, err := getSellOrder(invoice.ID)
	if err != nil {
		t.Error(err)
	}

	if so.Financed {
		t.Error("sell order should not be financed")
	}
}

func TestPlaceMultipleBids(t *testing.T) {
	
}