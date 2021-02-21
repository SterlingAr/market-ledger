package ledger

import (
	"testing"
)

func TestNewIssuer(t *testing.T) {
	cleanDB()
	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	iss, err := getIssuer(issuer.Name)
	if err != nil {
		t.Error(err)
	}

	if iss.Name != issuer.Name {
		t.Errorf("expected name =%v, actual name =%v", issuer.Name, iss.Name)
	}
}

func TestIssueInvoice(t *testing.T) {
	cleanDB()
	// An issuer party-A wants to finance an invoice with number invoice-1. The face value
	//of the invoice is €1000.
	//We need to record in the ledger that party-A has a €1000 invoice invoice-1 that
	//should be financed;

	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		FaceValue: 1000,
	}

	err = sellInvoice(issuer, invoice)

	if err != nil {
		t.Errorf("expected to sell invoice, got error %v", err)
	}

	iss, err := getIssuer(issuer.Name)

	if err != nil {
		t.Errorf("expected to find issuer %v, got error %v", issuer.Name, err)
	}

	var (
		invoices []*Invoice
		minValue = 900
		maxValue = 1000
	)

	invoices = getInvoices(iss, minValue, maxValue)

	if len(invoices) == 0 {
		t.Errorf("expected invoices = %v, actual invoices =%v", 1, len(invoices))
	}

	for _, invoice := range invoices {

		if invoice.FaceValue < minValue || invoice.FaceValue > maxValue {
			t.Errorf("unexpected FaceValue %v, minValue = %v, maxValue = %v", invoice.FaceValue, minValue, maxValue)
		}

		//if invoice.Financed {
		//	t.Error("invoice should not be financed")
		//}

		if invoice.IssuerID != issuer.ID {
			t.Errorf("mismatched issuer_id,  expected = %v  actual = %v", issuer.ID, invoice.IssuerID)
		}

	}
}
