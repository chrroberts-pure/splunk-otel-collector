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
	"net/http"

	"go.uber.org/zap"
)

type sparkService struct {
	logger                     *zap.Logger
	dbService                  databricksServiceIntf
	httpClient                 *http.Client
	sparkAPIURL                string
	sparkUIPort                int
	orgID                      string
	tok                        string
	sparkClusterClientProvider sparkClusterClientProvider
}

func newSparkService(logger *zap.Logger, dbService databricksServiceIntf, httpClient *http.Client, sparkAPIURL string, sparkUIPort int, orgID string, tok string) sparkService {
	return sparkService{
		logger:                     logger,
		dbService:                  dbService,
		httpClient:                 httpClient,
		sparkAPIURL:                sparkAPIURL,
		sparkUIPort:                sparkUIPort,
		orgID:                      orgID,
		tok:                        tok,
		sparkClusterClientProvider: newSparkClusterClient,
	}
}

type sparkClusterClientProvider func(
	logger *zap.Logger,
	httpClient *http.Client,
	sparkProxyURL string,
	orgID string,
	port int,
	token string,
	clusterID string,
) sparkClusterClientIntf

func (s sparkService) getSparkMetricsForCluster(clusterID string) (*sparkMetrics, error) {
	scc := s.sparkClusterClientProvider(
		s.logger,
		s.httpClient,
		s.sparkAPIURL,
		s.orgID,
		s.sparkUIPort,
		s.tok,
		clusterID,
	)
	unm := sparkUnmarshaller{scc: scc}
	return unm.metrics()
}
