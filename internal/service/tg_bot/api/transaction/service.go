package transaction

import "finance/internal/service/tg_bot/api"

type Service struct {
	client *api.Client
}

func New(client *api.Client) *Service {
	return &Service{
		client: client,
	}
}
