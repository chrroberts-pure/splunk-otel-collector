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
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"

	"github.com/signalfx/splunk-otel-collector/internal/receiver/databricksreceiver/internal/metadata"
)

const typeStr = "databricks"

func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createReceiverFunc(newDatabricksClient), component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.ReceiverConfig {
	scs := scraperhelper.NewDefaultScraperControllerSettings(typeStr)
	// we set the default collection interval to 30 seconds which is half of the
	// lowest job frequency of 1 minute
	scs.CollectionInterval = time.Second * 30
	return &Config{
		MaxResults:                25, // 25 is the max the API supports
		ScraperControllerSettings: scs,
	}
}

func createReceiverFunc(createDBClientFunc func(baseURL string, tok string, httpClient *http.Client, logger *zap.Logger) databricksClientIntf) func(
	_ context.Context,
	settings component.ReceiverCreateSettings,
	cfg component.ReceiverConfig,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	return func(
		_ context.Context,
		settings component.ReceiverCreateSettings,
		cfg component.ReceiverConfig,
		consumer consumer.Metrics,
	) (component.MetricsReceiver, error) {
		dbcfg := cfg.(*Config)
		dbcfg.resolveDatabricksEndpoint()

		httpClient, err := dbcfg.ToClient(nil, settings.TelemetrySettings)
		if err != nil {
			return nil, fmt.Errorf("%s: createReceiverFunc closure: %w", typeStr, err)
		}
		dbService := newDatabricksService(createDBClientFunc(dbcfg.Endpoint, dbcfg.Token, httpClient, settings.Logger), dbcfg.MaxResults)
		s := scraper{
			resourceOpt: metadata.WithDatabricksInstanceName(dbcfg.InstanceName),
			builder:     metadata.NewMetricsBuilder(dbcfg.Metrics, settings.BuildInfo),
			rmp:         newRunMetricsProvider(dbService),
			mp:          metricsProvider{dbService: dbService},
			ssvc: sparkService{
				dbsvc:      dbService,
				httpClient: httpClient,
				logger:     settings.Logger,
				orgID:      dbcfg.OrgID,
				tok:        dbcfg.Token,
			},
		}
		scrpr, err := scraperhelper.NewScraper(typeStr, s.scrape)
		if err != nil {
			return nil, fmt.Errorf("%s: createReceiverFunc closure: %w", typeStr, err)
		}
		return scraperhelper.NewScraperControllerReceiver(
			&dbcfg.ScraperControllerSettings,
			settings,
			consumer,
			scraperhelper.AddScraper(scrpr),
		)
	}
}
