package yundict

import (
	"encoding/json"
	"fmt"
)

type KeysService service

type KeysExportResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

// Export all keys
func (s *KeysService) Export(teamName string, projectName string, exportType string) (*KeysExportResponse, error) {
	path := fmt.Sprintf("/teams/%s/projects/%s/keys/export?type=%s", teamName, projectName, exportType)
	body, err := s.client.Get(path)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var result KeysExportResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
