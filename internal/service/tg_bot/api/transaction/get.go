package transaction

import (
	"encoding/json"
	"fmt"
	"net/url"

	"finance/internal/service/tg_bot/model"
)

func (s *Service) GetTransactions(
	txType string,
	categoryID string,
	dateFrom string,
	dateTo string,
) ([]model.Transaction, error) {

	var endpoint string

	switch txType {
	case "Income":
		endpoint = "getIncome"
	case "Expense":
		endpoint = "getExpense"
	default:
		return nil, fmt.Errorf("unknown transaction type")
	}

	u := fmt.Sprintf(
		"%s/transactions/%s?category_id=%s&date_from=%s&date_to=%s",
		s.client.BaseURL,
		endpoint,
		url.QueryEscape(categoryID),
		url.QueryEscape(dateFrom),
		url.QueryEscape(dateTo),
	)

	resp, err := s.client.Client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Value []model.Transaction `json:"value"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Value, nil
}
