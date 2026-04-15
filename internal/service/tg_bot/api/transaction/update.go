package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Service) UpdateTransaction(id int, amount int64) error {
	body, err := json.Marshal(map[string]interface{}{
		"id":     id,
		"amount": amount,
	})
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/transactions", s.client.BaseURL),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("update failed, status: %d", resp.StatusCode)
	}

	return nil
}
