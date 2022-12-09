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
	metricMaxOnHeapMem = "blockmanager.memory.maxonheapmem"
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
		case "blockmanager.memory.diskspaceused":
			builder.RecordDatabricksSparkBlockmanagerMemoryDiskspaceusedDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.maxmem":
			builder.RecordDatabricksSparkBlockmanagerMemoryMaxmemDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.maxoffheapmem":
			builder.RecordDatabricksSparkBlockmanagerMemoryMaxoffheapmemDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.maxonheapmem":
			builder.RecordDatabricksSparkBlockmanagerMemoryMaxonheapmemDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.memused":
			builder.RecordDatabricksSparkBlockmanagerMemoryMemusedDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.offheapmemused":
			builder.RecordDatabricksSparkBlockmanagerMemoryOffheapmemusedDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.onheapmemused":
			builder.RecordDatabricksSparkBlockmanagerMemoryOnheapmemusedDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.remainingmem":
			builder.RecordDatabricksSparkBlockmanagerMemoryRemainingmemDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.remainingoffheapmem":
			builder.RecordDatabricksSparkBlockmanagerMemoryRemainingoffheapmemDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "blockmanager.memory.remainingonheapmem":
			builder.RecordDatabricksSparkBlockmanagerMemoryRemainingonheapmemDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "dagscheduler.job.activejobs":
			builder.RecordDatabricksSparkDagschedulerJobActivejobsDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "dagscheduler.job.alljobs":
			builder.RecordDatabricksSparkDagschedulerJobAlljobsDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "dagscheduler.stage.failedstages":
			builder.RecordDatabricksSparkDagschedulerStageFailedstagesDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "dagscheduler.stage.runningstages":
			builder.RecordDatabricksSparkDagschedulerStageRunningstagesDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "dagscheduler.stage.waitingstages":
			builder.RecordDatabricksSparkDagschedulerStageWaitingstagesDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.directpoolmemory":
			builder.RecordDatabricksSparkExecutormetricsDirectpoolmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.jvmheapmemory":
			builder.RecordDatabricksSparkExecutormetricsJvmheapmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.jvmoffheapmemory":
			builder.RecordDatabricksSparkExecutormetricsJvmoffheapmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.majorgccount":
			builder.RecordDatabricksSparkExecutormetricsMajorgccountDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.majorgctime":
			builder.RecordDatabricksSparkExecutormetricsMajorgctimeDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.mappedpoolmemory":
			builder.RecordDatabricksSparkExecutormetricsMappedpoolmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.minorgccount":
			builder.RecordDatabricksSparkExecutormetricsMinorgccountDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.minorgctime":
			builder.RecordDatabricksSparkExecutormetricsMinorgctimeDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.offheapexecutionmemory":
			builder.RecordDatabricksSparkExecutormetricsOffheapexecutionmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.offheapstoragememory":
			builder.RecordDatabricksSparkExecutormetricsOffheapstoragememoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.offheapunifiedmemory":
			builder.RecordDatabricksSparkExecutormetricsOffheapunifiedmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.onheapexecutionmemory":
			builder.RecordDatabricksSparkExecutormetricsOnheapexecutionmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.onheapstoragememory":
			builder.RecordDatabricksSparkExecutormetricsOnheapstoragememoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.onheapunifiedmemory":
			builder.RecordDatabricksSparkExecutormetricsOnheapunifiedmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.processtreejvmrssmemory":
			builder.RecordDatabricksSparkExecutormetricsProcesstreejvmrssmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.processtreejvmvmemory":
			builder.RecordDatabricksSparkExecutormetricsProcesstreejvmvmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.processtreeotherrssmemory":
			builder.RecordDatabricksSparkExecutormetricsProcesstreeotherrssmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.processtreeothervmemory":
			builder.RecordDatabricksSparkExecutormetricsProcesstreeothervmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.processtreepythonrssmemory":
			builder.RecordDatabricksSparkExecutormetricsProcesstreepythonrssmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "executormetrics.processtreepythonvmemory":
			builder.RecordDatabricksSparkExecutormetricsProcesstreepythonvmemoryDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "jvmcpu.jvmcputime":
			builder.RecordDatabricksSparkJvmcpuJvmcputimeDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "livelistenerbus.queue.appstatus.size":
			builder.RecordDatabricksSparkLivelistenerbusQueueAppstatusSizeDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "livelistenerbus.queue.executormanagement.size":
			builder.RecordDatabricksSparkLivelistenerbusQueueExecutormanagementSizeDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "livelistenerbus.queue.shared.size":
			builder.RecordDatabricksSparkLivelistenerbusQueueSharedSizeDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "livelistenerbus.queue.streams.size":
			builder.RecordDatabricksSparkLivelistenerbusQueueStreamsSizeDataPoint(
				ts,
				int64(gauge.Value),
				clusterID,
				appID,
			)
		case "sparksqloperationmanager.numhiveoperations":
			builder.RecordDatabricksSparkSparksqloperationmanagerNumhiveoperationsDataPoint(
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
