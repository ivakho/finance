package category

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (s *Service) AddCategory(name string) error {
	body, err := json.Marshal(map[string]string{"name": name})
	if err != nil {
		return fmt.Errorf("failed to marshal category data: %w", err)
	}

	resp, err := s.client.Client.Post(
		fmt.Sprintf("%s/category", s.client.BaseURL),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
