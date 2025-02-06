package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestManualJournalVouchersGet(t *testing.T) {
	req := client.NewManualJournalVouchersGet()
	req.QueryParams().Filter.Set("ExternalImportReference eq 'Mews 2025-01-09'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
