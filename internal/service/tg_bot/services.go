package tg_bot

import (
	"finance/internal/service/tg_bot/api"
	"finance/internal/service/tg_bot/api/category"
	"finance/internal/service/tg_bot/api/transaction"
)

type Services struct {
	Category    *category.Service
	Transaction *transaction.Service
}

func NewServices(client *api.Client) *Services {
	return &Services{
		Category:    category.New(client),
		Transaction: transaction.New(client),
	}
}
