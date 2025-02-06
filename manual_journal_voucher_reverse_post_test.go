package poweroffice_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestManualJournalVoucherReversePost(t *testing.T) {
	req := client.NewManualJournalVoucherReversePost()
	req.PathParams().VoucherID = "b73d9219-f693-40d9-89ab-75aecd6f4399"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
