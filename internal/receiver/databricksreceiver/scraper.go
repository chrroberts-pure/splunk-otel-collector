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
	"fmt"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/signalfx/splunk-otel-collector/internal/receiver/databricksreceiver/internal/metadata"
)

// scraper provides a scrape method to a scraper controller receiver. The scrape
// method is the entry point into this receiver's functionality, running on a
// timer, and building metrics from metrics providers.
type scraper struct {
	instanceName string
	rmp          runMetricsProvider
	mp           metricsProvider
	smp          sparkService
	builder      *metadata.MetricsBuilder
}

func (s scraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	const errfmt = "scraper.scrape(): %w"
	var err error

	ts := pcommon.NewTimestampFromTime(time.Now())

	jobIDs, err := s.mp.addJobStatusMetrics(s.builder, ts)
	if err != nil {
		return pmetric.Metrics{}, fmt.Errorf(errfmt, err)
	}

	err = s.mp.addNumActiveRunsMetric(s.builder, ts)
	if err != nil {
		return pmetric.Metrics{}, fmt.Errorf(errfmt, err)
	}

	err = s.rmp.addMultiJobRunMetrics(jobIDs, s.builder, ts)
	if err != nil {
		return pmetric.Metrics{}, fmt.Errorf(errfmt, err)
	}

	return s.builder.Emit(), err
}
