package ledger

type Invoice struct {
	tableName struct{} `pg:"ledger.invoices"`
	ID        uint64
	FaceValue int
	IssuerID  uint64
	Issuer    *Issuer `pg:"rel:has-one"`
	Financed  bool
}
