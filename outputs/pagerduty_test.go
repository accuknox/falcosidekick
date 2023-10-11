// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package outputs

import (
	"encoding/json"
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/stretchr/testify/require"

	"github.com/falcosecurity/falcosidekick/types"
)

func TestPagerdutyPayload(t *testing.T) {
	var falcoTestInput = `{"output":"This is a test from falcosidekick","priority":"Debug","rule":"Test rule","hostname":"test-host","time":"2001-01-01T01:10:00Z","output_fields": {"hostname": "test-host", "proc.name":"falcosidekick", "proc.tty": 1234}}`
	var excpectedOutput = pagerduty.V2Event{
		RoutingKey: "",
		Action:     "trigger",
		Payload: &pagerduty.V2Payload{
			Summary:   "This is a test from falcosidekick",
			Source:    "falco",
			Severity:  "critical",
			Timestamp: "2001-01-01T01:10:00Z",
			Component: "",
			Group:     "",
			Class:     "",
			Details: map[string]interface{}{
				"hostname":  "test-host",
				"proc.name": "falcosidekick",
				"proc.tty":  float64(1234),
			},
		},
	}

	var f types.FalcoPayload
	json.Unmarshal([]byte(falcoTestInput), &f)

	event := createPagerdutyEvent(f, types.PagerdutyConfig{})

	require.Equal(t, excpectedOutput, event)
}
