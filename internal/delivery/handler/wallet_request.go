package handler

import (
	"go-bank-express/internal/entity"

	"github.com/shopspring/decimal"
)

type CreateWalletReq struct {
	Type       entity.WalletType `json:"type"`
	CurrencyId uint64            `json:"currency_id"`
}

type ListWalletReq struct {
	Type       entity.WalletType `json:"type"`
	Account    string            `json:"account"`
	CurrencyId uint64            `json:"currency_id"`
	Sort       string            `json:"sort"`
	Pagination string            `json:"page"`
}

type UpdateWalletBalanceReq struct {
	Type    entity.WalletType `json:"type"`
	Account string            `json:"account"`
	Amount  decimal.Decimal   `json:"amount"`
}
