package config

import (
	"os"
	"testing"
)

// Test: Init
func TestInit(t *testing.T) {
	os.Setenv("SLACK_TOKEN", "OK")

	c := Init()

	if c.Slack.Token != "OK" {
		t.Error("Expected ok")
	}
}
