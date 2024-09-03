// SPDX-License-Identifier: MIT OR Apache-2.0

package outputs

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/falcosecurity/falcosidekick/types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type timescaledbPayload struct {
	SQL    string `json:"sql"`
	Values []any  `json:"values"`
}

func NewTimescaleDBClient(config *types.Configuration, stats *types.Statistics, promStats *types.PromStatistics,
	statsdClient, dogstatsdClient *statsd.Client) (*Client, error) {

	ctx := context.Background()
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.TimescaleDB.User,
		config.TimescaleDB.Password,
		config.TimescaleDB.Host,
		config.TimescaleDB.Port,
		config.TimescaleDB.Database,
	)
	connPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Printf("[ERROR] : TimescaleDB - %v\n", err)
		return nil, ErrClientCreation
	}

	return &Client{
		OutputType:        "TimescaleDB",
		Config:            config,
		TimescaleDBClient: connPool,
		Stats:             stats,
		PromStats:         promStats,
		StatsdClient:      statsdClient,
		DogstatsdClient:   dogstatsdClient,
	}, nil
}

func newTimescaleDBPayload(kubearmorpayload types.KubearmorPayload, config *types.Configuration) timescaledbPayload {
	vals := make(map[string]any, 7+len(config.Customfields)+len(config.Templatedfields))
	vals[Time] = kubearmorpayload.Timestamp
	vals[Priority] = kubearmorpayload.EventType
	vals["Source Pod"] = kubearmorpayload.OutputFields["PodName"].(string)

	if kubearmorpayload.Hostname != "" {
		vals[Hostname] = kubearmorpayload.Hostname
	}

	for i, j := range kubearmorpayload.OutputFields {
		switch v := j.(type) {
		case string:
			for k := range config.Customfields {
				if i == k {
					vals[i] = strings.ReplaceAll(v, "\"", "")
				}
			}
			for k := range config.Templatedfields {
				if i == k {
					vals[i] = strings.ReplaceAll(v, "\"", "")
				}
			}
		default:
			continue
		}
	}

	i := 0
	retVals := make([]any, len(vals))
	var cols strings.Builder
	var args strings.Builder
	for k, v := range vals {
		cols.WriteString(k)
		fmt.Fprintf(&args, "$%d", i+1)
		if i < (len(vals) - 1) {
			cols.WriteString(",")
			args.WriteString(",")
		}

		str, isString := v.(string)
		if isString && (strings.ToLower(str) == "null") {
			retVals[i] = nil
		} else {
			retVals[i] = v
		}
		i++
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		config.TimescaleDB.HypertableName,
		cols.String(),
		args.String())

	return timescaledbPayload{SQL: sql, Values: retVals}
}

func (c *Client) TimescaleDBPost(kubearmorpayload types.KubearmorPayload) {
	c.Stats.TimescaleDB.Add(Total, 1)

	var ctx = context.Background()
	tsdbPayload := newTimescaleDBPayload(kubearmorpayload, c.Config)
	_, err := c.TimescaleDBClient.Exec(ctx, tsdbPayload.SQL, tsdbPayload.Values...)
	if err != nil {
		go c.CountMetric(Outputs, 1, []string{"output:timescaledb", "status:error"})
		c.Stats.TimescaleDB.Add(Error, 1)
		c.PromStats.Outputs.With(map[string]string{"destination": "timescaledb", "status": Error}).Inc()
		log.Printf("[ERROR] : TimescaleDB - %v\n", err)
		return
	}

	go c.CountMetric(Outputs, 1, []string{"output:timescaledb", "status:ok"})
	c.Stats.TimescaleDB.Add(OK, 1)
	c.PromStats.Outputs.With(map[string]string{"destination": "timescaledb", "status": OK}).Inc()

	if c.Config.Debug {
		log.Printf("[DEBUG] : TimescaleDB payload : %v\n", tsdbPayload)
	}
}

func (c *Client) WatchTimescaleDBPostAlerts() error {
	uid := uuid.Must(uuid.NewRandom()).String()

	conn := make(chan types.KubearmorPayload, 1000)
	defer close(conn)
	addAlertStruct(uid, conn)
	defer removeAlertStruct(uid)

	fmt.Println("discord running")
	for AlertRunning {
		select {
		case resp := <-conn:
			fmt.Println("response \n", resp)
			c.TimescaleDBPost(resp)
		}
	}
	fmt.Println("discord stopped")
	return nil
}
