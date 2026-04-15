package category

import (
	"fmt"
	"net/http"
)

func (s *Service) DeleteCategory(id int) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/category/%d", s.client.BaseURL, id),
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	resp, err := s.client.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("delete category failed: status %d", resp.StatusCode)
	}

	return nil
}
