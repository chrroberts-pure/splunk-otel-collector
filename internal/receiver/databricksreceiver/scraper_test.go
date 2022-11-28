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
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/signalfx/splunk-otel-collector/internal/receiver/databricksreceiver/internal/metadata"
)

func TestMetricsProvider_Scrape(t *testing.T) {
	const ignored = 25
	c := newDatabricksService(&testdataDBClient{}, ignored)
	var dbClient databricksServiceIntf = c
	scrpr := scraper{
		resourceOpt: metadata.WithDatabricksInstanceName("my-instance"),
		builder:     newTestMetricsBuilder(),
		rmp:         newRunMetricsProvider(c),
		mp:          metricsProvider{dbService: dbClient},
	}
	metrics, err := scrpr.scrape(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 4, metrics.MetricCount())
	attrs := metrics.ResourceMetrics().At(0).Resource().Attributes()
	v, _ := attrs.Get("databricks.instance.name")
	assert.Equal(t, "my-instance", v.Str())
}
