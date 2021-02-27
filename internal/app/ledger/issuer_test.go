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

	issuer := &Issuer{
		Name: "party-A",
	}

	err := newIssuer(issuer)
	if err != nil {
		t.Error(err)
	}

	invoice := &Invoice{
		Name: "invoice-1",
		FaceValue: 1000,
	}

	err = sellInvoice(issuer, invoice)

	if err != nil {
		t.Errorf("expected to sell invoice, got error %v", err)
	}

	var (
		minValue float64 = 900
		maxValue float64 = 1000
	)

	storedInvoice, err := getInvoice(invoice.Name)

	if err != nil {
		t.Errorf("expected to find invoice %v, got error %v", issuer.Name, err)
	}

	if storedInvoice.FaceValue < minValue || storedInvoice.FaceValue > maxValue {
		t.Errorf("unexpected FaceValue %v, minValue = %v, maxValue = %v", storedInvoice.FaceValue, minValue, maxValue)
	}

	if storedInvoice.Name != invoice.Name {
		t.Errorf("mismatched invoice name,  expected = %v  actual = %v", invoice.Name, storedInvoice.Name)
	}

	if storedInvoice.IssuerID != issuer.ID {
		t.Errorf("mismatched issuer_id,  expected = %v  actual = %v", issuer.ID, storedInvoice.IssuerID)
	}

}

func getInvoice(name string) (*Invoice, error) {
	var invoice Invoice

	err := db.Model(&invoice).Where("name = ?", name).Select()

	return &invoice, err
}
