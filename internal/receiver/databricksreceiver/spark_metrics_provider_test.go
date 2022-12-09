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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestStripSparkMetricKey(t *testing.T) {
	key := "app-20221117221047-0000.driver.BlockManager.memory.diskSpaceUsed_MB"
	appID, stripped := stripSparkMetricKey(key)
	assert.Equal(t, "app-20221117221047-0000", appID)
	assert.Equal(t, "blockmanager.memory.diskspaceused", stripped)
}

func TestSparkMetricsProvider(t *testing.T) {
	ssvc := newTestSparkService()
	mp := sparkMetricsProvider{ssvc}
	builder := newTestMetricsBuilder()
	err := mp.addSparkMetrics(builder, 0)
	require.NoError(t, err)
	emitted := builder.Emit()
	const numMetrics = 41
	assert.Equal(t, numMetrics, emitted.MetricCount())
	assert.Equal(t, numMetrics, emitted.DataPointCount())
	ms := emitted.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics()
	assert.Equal(t, numMetrics, ms.Len())
	metricMap := metricsByName(ms)
	assertMetricEq(t, metricMap, "blockmanager.memory.diskspaceused", 42)
	assertMetricEq(t, metricMap, "blockmanager.memory.maxmem", 123)
	assertMetricEq(t, metricMap, "blockmanager.memory.maxoffheapmem", 111)
	assertMetricEq(t, metricMap, "blockmanager.memory.maxonheapmem", 222)
	assertMetricEq(t, metricMap, "blockmanager.memory.memused", 333)
	assertMetricEq(t, metricMap, "blockmanager.memory.offheapmemused", 444)
	assertMetricEq(t, metricMap, "blockmanager.memory.onheapmemused", 555)
	assertMetricEq(t, metricMap, "blockmanager.memory.remainingmem", 666)
	assertMetricEq(t, metricMap, "blockmanager.memory.remainingoffheapmem", 777)
	assertMetricEq(t, metricMap, "blockmanager.memory.remainingonheapmem", 888)
	assertMetricEq(t, metricMap, "dagscheduler.job.activejobs", 999)
	assertMetricEq(t, metricMap, "dagscheduler.job.alljobs", 1111)
	assertMetricEq(t, metricMap, "dagscheduler.stage.failedstages", 2222)
	assertMetricEq(t, metricMap, "dagscheduler.stage.runningstages", 3333)
	assertMetricEq(t, metricMap, "dagscheduler.stage.waitingstages", 4444)
	assertMetricEq(t, metricMap, "executormetrics.directpoolmemory", 591058)
	assertMetricEq(t, metricMap, "executormetrics.jvmheapmemory", 1748700144)
	assertMetricEq(t, metricMap, "executormetrics.jvmoffheapmemory", 269709952)
	assertMetricEq(t, metricMap, "executormetrics.majorgccount", 5)
	assertMetricEq(t, metricMap, "executormetrics.majorgctime", 748)
	assertMetricEq(t, metricMap, "executormetrics.mappedpoolmemory", 5555)
	assertMetricEq(t, metricMap, "executormetrics.minorgccount", 5)
	assertMetricEq(t, metricMap, "executormetrics.minorgctime", 200)
	assertMetricEq(t, metricMap, "executormetrics.offheapexecutionmemory", 6666)
	assertMetricEq(t, metricMap, "executormetrics.offheapstoragememory", 7777)
	assertMetricEq(t, metricMap, "executormetrics.offheapunifiedmemory", 8888)
	assertMetricEq(t, metricMap, "executormetrics.onheapexecutionmemory", 9999)
	assertMetricEq(t, metricMap, "executormetrics.onheapstoragememory", 11111)
	assertMetricEq(t, metricMap, "executormetrics.onheapunifiedmemory", 22222)
	assertMetricEq(t, metricMap, "executormetrics.processtreejvmrssmemory", 33333)
	assertMetricEq(t, metricMap, "executormetrics.processtreejvmvmemory", 44444)
	assertMetricEq(t, metricMap, "executormetrics.processtreeotherrssmemory", 55555)
	assertMetricEq(t, metricMap, "executormetrics.processtreeothervmemory", 66666)
	assertMetricEq(t, metricMap, "executormetrics.processtreepythonrssmemory", 77777)
	assertMetricEq(t, metricMap, "executormetrics.processtreepythonvmemory", 88888)
	assertMetricEq(t, metricMap, "jvmcpu.jvmcputime", 57690000000)
	assertMetricEq(t, metricMap, "livelistenerbus.queue.appstatus.size", 99999)
	assertMetricEq(t, metricMap, "livelistenerbus.queue.executormanagement.size", 111111)
	assertMetricEq(t, metricMap, "livelistenerbus.queue.shared.size", 222222)
	assertMetricEq(t, metricMap, "livelistenerbus.queue.streams.size", 333333)
	assertMetricEq(t, metricMap, "sparksqloperationmanager.numhiveoperations", 444444)
}

func assertMetricEq(t *testing.T, metricMap map[string]pmetric.Metric, metricName string, expected int) bool {
	return assert.EqualValues(t, expected, metricMap["databricks.spark."+metricName].Gauge().DataPoints().At(0).IntValue())
}
