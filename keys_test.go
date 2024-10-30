package yundict

import (
	"os"
	"testing"
)

var (
	client *Client
)

func init() {
	token := os.Getenv("YUNDICT_API_TOKEN")
	client = NewClient(token)
}

func TestKeysExport(t *testing.T) {
	res, err := client.Keys.Export("test-team", "test-project", "json")
	if err != nil {
		t.Errorf("Keys.Export returned error: %v", err)
	}

	if !res.Success {
		t.Errorf("Keys.Export returned success: %v", res.Success)
	}

	if res.Data == "" {
		t.Errorf("Keys.Export returned empty data")
	}

	// data starts with "http"
	if res.Data[:4] != "http" {
		t.Errorf("Keys.Export returned invalid data: %v", res.Data)
	}
}
