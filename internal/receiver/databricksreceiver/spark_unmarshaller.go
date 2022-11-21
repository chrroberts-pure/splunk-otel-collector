// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package databricksreceiver

import (
	"encoding/json"
	"fmt"
)

type sparkUnmarshaller struct {
	scc sparkClusterClientIntf
}

func (u sparkUnmarshaller) metrics() (*sparkMetrics, error) {
	bytes, err := u.scc.metrics()
	if err != nil {
		return nil, fmt.Errorf("failed to get metrics from spark cluster client: %w", err)
	}
	sm := sparkMetrics{}
	err = json.Unmarshal(bytes, &sm)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spark metrics: %w", err)
	}
	return &sm, nil
}

type sparkMetrics struct {
	Gauges     map[string]sparkGauge
	Counters   map[string]sparkCounter
	Histograms map[string]sparkHistogram
	Meters     map[string]any
	Timers     map[string]sparkTimer
}

type sparkGauge struct {
	Value int `json:"value"`
}

type sparkCounter struct {
	Count int `json:"count"`
}

type sparkHistogram struct {
	Count  int     `json:"count"`
	Max    int     `json:"max"`
	Mean   float64 `json:"mean"`
	Min    int     `json:"min"`
	P50    float64 `json:"p50"`
	P75    float64 `json:"p75"`
	P95    float64 `json:"p95"`
	P98    float64 `json:"p98"`
	P99    float64 `json:"p99"`
	P999   float64 `json:"p999"`
	Stddev float64 `json:"stddev"`
}

type sparkTimer struct {
	Count         int     `json:"count"`
	Max           float64 `json:"max"`
	Mean          float64 `json:"mean"`
	Min           float64 `json:"min"`
	P50           float64 `json:"p50"`
	P75           float64 `json:"p75"`
	P95           float64 `json:"p95"`
	P98           float64 `json:"p98"`
	P99           float64 `json:"p99"`
	P999          float64 `json:"p999"`
	Stddev        float64 `json:"stddev"`
	M15Rate       float64 `json:"m15_rate"`
	M1Rate        float64 `json:"m1_rate"`
	M5Rate        float64 `json:"m5_rate"`
	MeanRate      float64 `json:"mean_rate"`
	DurationUnits string  `json:"duration_units"`
	RateUnits     string  `json:"rate_units"`
}
