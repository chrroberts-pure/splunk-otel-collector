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
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/receiver/scraperhelper"

	"github.com/signalfx/splunk-otel-collector/internal/receiver/databricksreceiver/internal/metadata"
)

type Config struct {
	Metrics                                 metadata.MetricsSettings `mapstructure:"metrics"`
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	confighttp.HTTPClientSettings           `mapstructure:",squash"`
	OrgID                                   string `mapstructure:"org_id"`
	InstanceName                            string `mapstructure:"instance_name"`
	Token                                   string `mapstructure:"token"`
	MaxResults                              int    `mapstructure:"max_results"`
	SparkAPIURL                             string `mapstructure:"spark_api_url"`
	SparkUIPort                             int    `mapstructure:"spark_ui_port"`
}

func (c *Config) Validate() error {
	if c.OrgID == "" && c.Endpoint == "" {
		return errors.New("both `org_id` and `endpoint` are empty, at least one must be specified")
	}
	if c.InstanceName == "" {
		return errors.New("instance_name is empty")
	}
	if c.Token == "" {
		return errors.New("token is empty")
	}
	return nil
}

func (c *Config) resolveDatabricksEndpoint() {
	if c.OrgID == "" {
		return
	}
	if c.Endpoint != "" {
		return
	}
	lastChar := c.OrgID[len(c.OrgID)-1]
	c.Endpoint = fmt.Sprintf("https://adb-%s.%c.azuredatabricks.net", c.OrgID, lastChar)
}
