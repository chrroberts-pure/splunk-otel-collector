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

type sparkService struct {
	logger                   *zap.Logger
	dbsvc                    databricksServiceIntf
	httpClient               *http.Client
	sparkAPIURL              string
	sparkUIPort              int
	orgID                    string
	tok                      string
	createSparkClusterClient func(
		logger *zap.Logger,
		httpClient *http.Client,
		sparkProxyURL string,
		orgID string,
		port int,
		token string,
		clusterID string,
	) sparkClusterClientIntf
}

func (s sparkService) getSparkMetricsForAllClusters() (map[string]*sparkClusterMetrics, error) {
	clusterIDs, err := s.dbsvc.runningClusterIDs()
	if err != nil {
		return nil, fmt.Errorf("error getting cluster IDs: %w", err)
	}
	out := map[string]*sparkClusterMetrics{}
	for _, id := range clusterIDs {
		metrics, err := s.getSparkMetricsForCluster(id)
		if err != nil {
			return nil, fmt.Errorf("error getting spark metrics for cluster: %s: %w", id, err)
		}
		out[id] = metrics
	}
	return out, nil
}

func (s sparkService) getSparkMetricsForCluster(clusterID string) (*sparkClusterMetrics, error) {
	scc := s.createSparkClusterClient(
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
