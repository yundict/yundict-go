package yundict

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	client *Client
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	token := os.Getenv("YUNDICT_API_TOKEN")
	client = NewClient(token)
	client.Endpoint = os.Getenv("YUNDICT_API_ENDPOINT")
}

func TestKeysExport(t *testing.T) {
	res, err := client.Keys.Export(
		os.Getenv("YUNDICT_TEST_TEAM_NAME"),
		os.Getenv("YUNDICT_TEST_PROJECT_NAME"),
		"key-value-json",
		[]string{"en"},
	)
	if err != nil {
		t.Errorf("Keys.Export returned error: %v", err)
	}

	if !res.Success {
		t.Errorf("Keys.Export returned success: %v", res.Success)
	}

	fileURL := res.Data
	if fileURL == "" {
		t.Errorf("Keys.Export returned empty data")
	}

	// data starts with "http"
	if fileURL[:4] != "http" {
		t.Errorf("Keys.Export returned invalid data: %v", fileURL)
	}

	// fetch the json file
	jsonData, err := fetchJsonFile(fileURL)
	if err != nil {
		t.Errorf("Failed to fetch JSON file: %v", err)
	}

	// check the json data key-value pairs length
	if len(jsonData) == 0 {
		t.Errorf("JSON data is empty")
	}
}

func fetchJsonFile(url string) (map[string]interface{}, error) {
	// http download the file
	req, _ := http.NewRequest("GET", url, nil)

	// Send the request via a client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return jsonData, nil
}
