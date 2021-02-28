package v1

import "testing"

func TestNewSellOrder(t *testing.T) {
	cleanDB()
	// an existing invoice, make it available so that bids can be placed
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		Name: "invoice-1",
		Issuer:      nil,
		Bids:        nil,
	}

	err = sellInvoice(issuer, invoice)
	if err != nil {
		t.Error(err)
	}

	_, err = newSellOrder(invoice)

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
