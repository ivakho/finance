package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (s *Service) AddTransaction(categoryID string, txType string, amount float64, createdAt time.Time) error {
	catID, err := strconv.Atoi(categoryID)
	if err != nil {
		return fmt.Errorf("invalid categoryID: %w", err)
	}

	reqBody := map[string]interface{}{
		"category_id": catID,
		"type":        txType,
		"amount":      int64(amount),             
		"created_at":  createdAt.Format("2006-01-02"),
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction data: %w", err)
	}

	resp, err := s.client.Client.Post(
		fmt.Sprintf("%s/transactions", s.client.BaseURL),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errResp map[string]interface{}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		return fmt.Errorf("server returned status %d: %v", resp.StatusCode, errResp)
	}

	return nil
}