package v1

import (
	"context"
	v1 "github.com/SterlingAr/market-ledger/api/proto/v1"
)

func Sell(ctx context.Context, req * v1.Invoice) (*v1.SellResponse, error) {
	var (
		res * v1.SellResponse
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

	res = &v1.SellResponse{
		Id: issuer.ID,
	}

	return res, tx.Commit()
}