package category

import (
	"encoding/json"
	"fmt"

	"finance/internal/service/tg_bot/model"
)

func (s *Service) GetCategories() ([]model.Category, error) {
	resp, err := s.client.Client.Get(fmt.Sprintf("%s/category", s.client.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Value []model.Category `json:"value"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Value, nil
}
