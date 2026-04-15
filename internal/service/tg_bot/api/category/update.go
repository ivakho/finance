package category

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Service) UpdateCategory(id int, name string) error {
	body, err := json.Marshal(map[string]interface{}{
		"id":   id,
		"name": name,
	})

	if err != nil {
		return fmt.Errorf("failed to marshal category data: %w", err)
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/category", s.client.BaseURL),
		bytes.NewBuffer(body),
	)

	if err != nil {
		return fmt.Errorf("failed to create update request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
