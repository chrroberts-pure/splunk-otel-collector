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
	"net/http"

	"go.uber.org/zap"
)

type sparkClusterClient struct {
	logger     *zap.Logger
	authClient authClient
}

const (
	sparkMetricsPath = "/metrics/json/"
)

type sparkClusterClientIntf interface {
	metrics() ([]byte, error)
}

func newSparkClusterClient(logger *zap.Logger, httpClient *http.Client, sparkProxyURL string, orgID string, port int, token string, clusterID string) sparkClusterClientIntf {
	return sparkClusterClient{
		authClient: authClient{
			httpClient: httpClient,
			endpoint:   fmt.Sprintf("%s/driver-proxy-api/o/%s/%s/%d", sparkProxyURL, orgID, clusterID, port),
			tok:        token,
		},
		logger: logger,
	}
}

func (c sparkClusterClient) metrics() ([]byte, error) {
	return c.authClient.get(sparkMetricsPath)
}
