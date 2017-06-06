package config

import (
	"flag"

	"github.com/bureaux-a-partager/lucca-leaves-slack-reminder/lib/utils"
)

type Config struct {
	Slack struct {
		Token     string
		Channel   string
		IconEmoji string
		Username  string
	}
	Lucca struct {
		Token string ""
	}
}

// Init configuration
func Init() Config {
	c := Config{}

	// TODO Convert with reflect
	c.Slack.Token = utils.GetEnv("SLACK_TOKEN", "")
	c.Slack.Channel = utils.GetEnv("SLACK_CHANNEL", "")
	c.Slack.IconEmoji = utils.GetEnv("SLACK_ICON_EMOJI", ":sunglasses:")
	c.Slack.Username = utils.GetEnv("SLACK_USERNAME", "Life team")

	c.Lucca.Token = utils.GetEnv("LUCCA_TOKEN", "")

	flag.Parse()

	return c
}
