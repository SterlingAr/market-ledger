package v1

import (
	"context"
	v1 "github.com/SterlingAr/market-ledger/api/proto/v1"
)

func (m MarketLedgerServer) Sell(ctx context.Context, req *v1.NewInvoiceRequest) (*v1.NewInvoiceResponse, error) {
	var (
		res    *v1.NewInvoiceResponse
		issuer Issuer
	)

	tx, err := db.Begin()

	if err != nil {
		return res, err
	}

	defer txCloseLog(tx)

	err = tx.Model(&issuer).Where("id = ?", req.IssuerId).Select()

	if err != nil {
		return res, err
	}

	invoice := &Invoice{
		Name:        req.Name,
		FaceValue:   req.FaceValue,
		NeededValue: req.NeededValue,
		IssuerID:    req.IssuerId,
		Issuer:      &issuer,
	}

	_, err = tx.Model(invoice).Returning("*").Insert()

	if err != nil {
		return res, err
	}

	res = &v1.NewInvoiceResponse{
		Id: issuer.ID,
	}

	return res, tx.Commit()
}

func (m MarketLedgerServer) NewSellOrder(ctx context.Context, req *v1.NewSellOrderRequest) (*v1.NewSellOrderResponse, error) {
	var (
		res * v1.NewSellOrderResponse
		invoice Invoice
	)

	tx, err := db.Begin()

	if err != nil {
		return res, err
	}

	defer txCloseLog(tx)

	err = tx.Model(&invoice).Where("id = ?", req.InvoiceId).Select()

	if err != nil {
		return res, err
	}

	sellOrder := &SellOrder{
		InvoiceID: invoice.ID,
		Invoice:   &invoice,
		Financed:  false,
	}

	_, err = tx.Model(sellOrder).Returning("*").Insert()

	res = &v1.NewSellOrderResponse{
		SellOrderId: sellOrder.ID,
	}

	return res, tx.Commit()
}

func (m MarketLedgerServer) NewBid(ctx context.Context, req *v1.NewBidRequest) (*v1.NewBidResponse, error) {
	var (
		res * v1.NewBidResponse
		sellOrder SellOrder
		investor Investor
		reservedValue float64
	)

	tx, err := db.Begin()

	if err != nil {
		return res, err
	}

	defer txCloseLog(tx)

	err = tx.Model(&sellOrder).Where("invoice_id = ?", req.SellOrderId).Relation("Invoice").Select()

	if err != nil {
		return res, err
	}

	err = tx.Model(&investor).Where("id = ?", req.InvestorId).Select()

	if err != nil {
		return res, err
	}

	bid := &Bid{
		InvestmentValue: req.InvestedValue,
		Discount: req.Discount,
	}
	
	err = investor.newBid(sellOrder.Invoice, bid)

	if err != nil {
		return res, err
	}

	reservedValue = investmentDiscount(bid.InvestmentValue, bid.Discount)

	res = &v1.NewBidResponse{
		Position:       uint64(bid.Position),
		ReservedValue:  reservedValue,
		ExpectedProfit: bid.InvestmentValue - reservedValue,
	}

	return res, tx.Commit()
}

func (m MarketLedgerServer) MatchingAlgorithm(ctx context.Context, req *v1.MatchingAlgorithmRequest) (*v1.MatchingAlgorithmResponse, error) {
	err := matchingAlgorithm()
	return &v1.MatchingAlgorithmResponse{}, err
}

func (m MarketLedgerServer) GetLedger(ctx context.Context, req *v1.Empty) (*v1.LedgerResponse, error) {

	ledger := getLedger()

	res := &v1.LedgerResponse{
		Entries:  make([]*v1.LedgerEntry, len(ledger.Entries)),
	}

	for _, e := range ledger.Entries {
		res.Entries = append(res.Entries, &v1.LedgerEntry{
			InvestorName:    e.InvestorName,
			InvoiceName:     e.InvoiceName,
			InvestedBalance: e.InvestedBalance,
			ReservedBalance: e.ReservedBalance,
			ExpectedProfit:  e.ExpectedProfit,
		})
	}

	return res, nil
}

