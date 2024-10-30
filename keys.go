package yundict

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type KeysService service

type KeysExportResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

// Export all keys
func (s *KeysService) Export(teamName string, projectName string, exportType string, languages []string) (*KeysExportResponse, error) {
	queryParams := url.Values{}
	queryParams.Add("type", exportType)
	if len(languages) > 0 {
		queryParams.Add("languages", strings.Join(languages, ","))
	}

	path := fmt.Sprintf("/teams/%s/projects/%s/keys/export", teamName, projectName)
	fullPath := path + "?" + queryParams.Encode()

	body, err := s.client.Get(fullPath)
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
