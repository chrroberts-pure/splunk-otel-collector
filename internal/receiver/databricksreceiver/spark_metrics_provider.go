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
	"fmt"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/signalfx/splunk-otel-collector/internal/receiver/databricksreceiver/internal/metadata"
)

const (
	metricDiskSpaceUsed = "blockmanager.memory.diskspaceused"
	metricMaxmem        = "blockmanager.memory.maxmem"
)

type sparkMetricsProvider struct {
	ssvc sparkService
}

func (p sparkMetricsProvider) addSparkMetrics(builder *metadata.MetricsBuilder, ts pcommon.Timestamp) error {
	clusterMetrics, err := p.ssvc.getSparkMetricsForAllClusters()
	if err != nil {
		return fmt.Errorf("error getting spark metrics for all clusters: %w", err)
	}
	for clusterID, clusterMetric := range clusterMetrics {
		p.addClusterMetric(builder, clusterMetric, ts, clusterID)
	}
	return nil
}

func (p sparkMetricsProvider) addClusterMetric(builder *metadata.MetricsBuilder, m *sparkClusterMetrics, ts pcommon.Timestamp, clusterID string) {
	for key, gauge := range m.Gauges {
		appID, stripped := stripSparkMetricKey(key)
		switch stripped {
		case metricDiskSpaceUsed:
			builder.RecordDatabricksSparkBlockmanagerMemoryDiskspaceusedDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case metricMaxmem:
			builder.RecordDatabricksSparkBlockmanagerMemoryMaxmemDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		}
	}
}

func stripSparkMetricKey(s string) (string, string) {
	parts := strings.Split(s, ".")
	metricParts := parts[2:]
	lastPart := metricParts[len(metricParts)-1]
	if strings.HasSuffix(lastPart, "_MB") {
		metricParts[len(metricParts)-1] = lastPart[:len(lastPart)-3]
	}
	joined := strings.Join(metricParts, ".")
	return parts[0], strings.ToLower(joined)
}
