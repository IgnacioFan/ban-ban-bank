package handler

import (
	"go-bank-express/internal/entity"

	"github.com/shopspring/decimal"
)

type TransferReq struct {
	Type   entity.TransactionType `json:"transaction_type"`
	From   string                 `json:"from"`
	To     string                 `json:"to"`
	Amount decimal.Decimal        `json:"amount"`
}

type ListTransactionsReq struct {
	ID         uint64                   `json:"id"`
	Type       entity.TransactionType   `json:"transaction_type"`
	From       string                   `json:"from"`
	To         string                   `json:"to"`
	CurrencyId uint64                   `json:"currency_id"`
	Status     entity.TransactionStatus `json:"transaction_status"`
	Remark     string                   `json:"remark"`
	Sort       string                   `json:"sort"`
	Pagination string                   `json:"page"`
}
