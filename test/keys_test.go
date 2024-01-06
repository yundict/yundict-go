package test

import (
	"os"
	"testing"

	"github.com/yundict/yundict-go"
)

var (
	client *yundict.Client
)

func init() {
	token := os.Getenv("YUNDICT_API_TOKEN")
	client = yundict.NewClient(token)
}

func TestKeysExport(t *testing.T) {
	res, err := client.Keys.Export("VCj0-P0d", "GtiQs", "json")
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
