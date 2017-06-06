package slack

import (
	"fmt"

	"github.com/bureaux-a-partager/lucca-leaves-slack-reminder/lib/config"
	"github.com/nlopes/slack"
)

// Send text to a Slack chan
func Send(config config.Config, text string) {
	api := slack.New(config.Slack.Token)
	params := slack.PostMessageParameters{}
	params.IconEmoji = config.Slack.IconEmoji
	params.Username = config.Slack.Username
	channelID, timestamp, err := api.PostMessage(config.Slack.Channel, text, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
