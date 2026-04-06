package transaction

import (
	"fmt"
	"net/http"
)

func (s *Service) DeleteTransaction(id int) error {

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/transactions/%d", s.client.BaseURL, id),
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

	return nil
}
