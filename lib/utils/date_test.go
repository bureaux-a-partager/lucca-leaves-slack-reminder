package utils

import "testing"

// Test: DateIsoToFr
func TestDateIsoToFr(t *testing.T) {
	v := DateIsoToFr("2017-06-06")

	if v != "06/06/2017" {
		t.Error("Expected ok")
	}
}
