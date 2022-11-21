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
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestSparkService(t *testing.T) {
	nopLogger := zap.New(zapcore.NewNopCore())
	dbsvc := newDatabricksService(&testdataDBClient{}, 25)
	ssvc := sparkService{
		logger:    nopLogger,
		dbService: dbsvc,
		sparkClusterClientProvider: func(*zap.Logger, *http.Client, string, string, int, string, string) sparkClusterClientIntf {
			return testdataSparkClusterClient{}
		},
	}
	clusterIDs, err := dbsvc.runningClusterIDs()
	require.NoError(t, err)
	for _, clusterID := range clusterIDs {
		metrics, err := ssvc.getSparkMetricsForCluster(clusterID)
		require.NoError(t, err)
		fmt.Printf("clusterID: %s, metrics: ->%v<-\n", clusterID, metrics)
	}
}

type testdataSparkClusterClient struct{}

func (c testdataSparkClusterClient) metrics() ([]byte, error) {
	return os.ReadFile(filepath.Join("testdata", "spark-metrics.json"))
}
