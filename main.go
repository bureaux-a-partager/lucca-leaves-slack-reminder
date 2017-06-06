package main

import (
	"fmt"

	"github.com/bureaux-a-partager/lucca-leaves-slack-reminder/lib/config"
	"github.com/bureaux-a-partager/lucca-leaves-slack-reminder/lib/lucca"
	"github.com/bureaux-a-partager/lucca-leaves-slack-reminder/lib/slack"
)

// Main function
func main() {
	c := config.Init()

	l := lucca.GetLeaves(c.Lucca.Token)
	e := l.ListEmployees()
	t := lucca.FormatEmployees(e)

	slack.Send(c, t)
	fmt.Println(t)
}
