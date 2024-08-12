package main

import (
	"fmt"
	"sync"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/falcosecurity/falcosidekick/outputs"
	"github.com/falcosecurity/falcosidekick/types"
)

var globalMap map[string]*outputs.Client
var AlertLock *sync.RWMutex
var AlertBufferChannel chan []byte
var slackClient *outputs.Client
var statsdClient, dogstatsdClient *statsd.Client

func main() {
	var t1 types.SlackOutputConfig
	t1.WebhookURL = "https://hooks.slack.com/services/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"
	t1.Channel = "integration-alerts"
	t1.Footer = "filters"
	t1.Icon = "https://help.accuknox.com/assets/images/logo.png"
	t1.Username = "accuknox"
	cf1 := types.Configuration{
		Slack: t1,
	}
	stats := &types.Statistics{}
	promStats := &types.PromStatistics{}
	initClientArgs := &types.InitClientArgs{
		Config:          &cf1,
		Stats:           stats,
		DogstatsdClient: dogstatsdClient,
		PromStats:       promStats,
	}
	c1, err := outputs.NewClient("Slack", cf1.Slack.WebhookURL, cf1.Slack.MutualTLS, cf1.Slack.CheckCert, *initClientArgs)
	if err != nil {
		fmt.Println("error---")
	}
	// var t2 types.SlackOutputConfig
	// t2.WebhookURL = ""
	// cf2 := types.Configuration{
	// 	Slack: t2,
	// }

	// c2 := outputs.Client{
	// 	Config: &cf2,
	// }
	// outputs.AlertLock = &sync.RWMutex{}
	// outputs.AlertRunning = true
	fmt.Println(c1)
	outputs.InitSidekick()

	go c1.SendAlerts()
	go c1.AddAlertFromBuffChan()
	go c1.WatchSlackAlerts()

	select {}
}
