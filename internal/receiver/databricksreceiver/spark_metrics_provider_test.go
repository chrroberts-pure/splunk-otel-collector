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
	assert.Equal(t, 2, emitted.MetricCount())
	assert.Equal(t, 2, emitted.DataPointCount())
	ms := emitted.ResourceMetrics().At(0).ScopeMetrics().At(0).Metrics()
	assert.Equal(t, 1, ms.Len())
	metricMap := metricsByName(ms)
	diskSpaceUsed := metricMap["databricks.spark.blockmanager.memory.diskspaceused"]
	assert.EqualValues(t, 42, diskSpaceUsed.Gauge().DataPoints().At(0).IntValue())
}
