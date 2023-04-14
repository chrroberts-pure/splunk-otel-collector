// Code generated by monitor-code-gen. DO NOT EDIT.

package traefik

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "traefik"

var groupSet = map[string]bool{}

const (
	goGcDurationSeconds                           = "go_gc_duration_seconds"
	goGcDurationSecondsCount                      = "go_gc_duration_seconds_count"
	goGcDurationSecondsSum                        = "go_gc_duration_seconds_sum"
	goGoroutines                                  = "go_goroutines"
	goMemstatsAllocBytes                          = "go_memstats_alloc_bytes"
	goMemstatsAllocBytesTotal                     = "go_memstats_alloc_bytes_total"
	goMemstatsBuckHashSysBytes                    = "go_memstats_buck_hash_sys_bytes"
	goMemstatsFreesTotal                          = "go_memstats_frees_total"
	goMemstatsGcCPUFraction                       = "go_memstats_gc_cpu_fraction"
	goMemstatsGcSysBytes                          = "go_memstats_gc_sys_bytes"
	goMemstatsHeapAllocBytes                      = "go_memstats_heap_alloc_bytes"
	goMemstatsHeapIdleBytes                       = "go_memstats_heap_idle_bytes"
	goMemstatsHeapInuseBytes                      = "go_memstats_heap_inuse_bytes"
	goMemstatsHeapObjects                         = "go_memstats_heap_objects"
	goMemstatsHeapReleasedBytes                   = "go_memstats_heap_released_bytes"
	goMemstatsHeapSysBytes                        = "go_memstats_heap_sys_bytes"
	goMemstatsLastGcTimeSeconds                   = "go_memstats_last_gc_time_seconds"
	goMemstatsLookupsTotal                        = "go_memstats_lookups_total"
	goMemstatsMallocsTotal                        = "go_memstats_mallocs_total"
	goMemstatsMcacheInuseBytes                    = "go_memstats_mcache_inuse_bytes"
	goMemstatsMcacheSysBytes                      = "go_memstats_mcache_sys_bytes"
	goMemstatsMspanInuseBytes                     = "go_memstats_mspan_inuse_bytes"
	goMemstatsMspanSysBytes                       = "go_memstats_mspan_sys_bytes"
	goMemstatsNextGcBytes                         = "go_memstats_next_gc_bytes"
	goMemstatsOtherSysBytes                       = "go_memstats_other_sys_bytes"
	goMemstatsStackInuseBytes                     = "go_memstats_stack_inuse_bytes"
	goMemstatsStackSysBytes                       = "go_memstats_stack_sys_bytes"
	goMemstatsSysBytes                            = "go_memstats_sys_bytes"
	goThreads                                     = "go_threads"
	processCPUSecondsTotal                        = "process_cpu_seconds_total"
	processMaxFds                                 = "process_max_fds"
	processOpenFds                                = "process_open_fds"
	processResidentMemoryBytes                    = "process_resident_memory_bytes"
	processStartTimeSeconds                       = "process_start_time_seconds"
	processVirtualMemoryBytes                     = "process_virtual_memory_bytes"
	traefikBackendOpenConnections                 = "traefik_backend_open_connections"
	traefikBackendRequestDurationSecondsBucket    = "traefik_backend_request_duration_seconds_bucket"
	traefikBackendRequestDurationSecondsCount     = "traefik_backend_request_duration_seconds_count"
	traefikBackendRequestDurationSecondsSum       = "traefik_backend_request_duration_seconds_sum"
	traefikBackendRequestsTotal                   = "traefik_backend_requests_total"
	traefikBackendServerUp                        = "traefik_backend_server_up"
	traefikConfigLastReloadFailure                = "traefik_config_last_reload_failure"
	traefikConfigLastReloadSuccess                = "traefik_config_last_reload_success"
	traefikConfigReloadsFailureTotal              = "traefik_config_reloads_failure_total"
	traefikConfigReloadsTotal                     = "traefik_config_reloads_total"
	traefikEntrypointOpenConnections              = "traefik_entrypoint_open_connections"
	traefikEntrypointRequestDurationSecondsBucket = "traefik_entrypoint_request_duration_seconds_bucket"
	traefikEntrypointRequestDurationSecondsCount  = "traefik_entrypoint_request_duration_seconds_count"
	traefikEntrypointRequestDurationSecondsSum    = "traefik_entrypoint_request_duration_seconds_sum"
	traefikEntrypointRequestsTotal                = "traefik_entrypoint_requests_total"
)

var metricSet = map[string]monitors.MetricInfo{
	goGcDurationSeconds:                           {Type: datapoint.Counter},
	goGcDurationSecondsCount:                      {Type: datapoint.Counter},
	goGcDurationSecondsSum:                        {Type: datapoint.Counter},
	goGoroutines:                                  {Type: datapoint.Counter},
	goMemstatsAllocBytes:                          {Type: datapoint.Gauge},
	goMemstatsAllocBytesTotal:                     {Type: datapoint.Counter},
	goMemstatsBuckHashSysBytes:                    {Type: datapoint.Gauge},
	goMemstatsFreesTotal:                          {Type: datapoint.Counter},
	goMemstatsGcCPUFraction:                       {Type: datapoint.Gauge},
	goMemstatsGcSysBytes:                          {Type: datapoint.Gauge},
	goMemstatsHeapAllocBytes:                      {Type: datapoint.Gauge},
	goMemstatsHeapIdleBytes:                       {Type: datapoint.Gauge},
	goMemstatsHeapInuseBytes:                      {Type: datapoint.Gauge},
	goMemstatsHeapObjects:                         {Type: datapoint.Gauge},
	goMemstatsHeapReleasedBytes:                   {Type: datapoint.Gauge},
	goMemstatsHeapSysBytes:                        {Type: datapoint.Gauge},
	goMemstatsLastGcTimeSeconds:                   {Type: datapoint.Gauge},
	goMemstatsLookupsTotal:                        {Type: datapoint.Counter},
	goMemstatsMallocsTotal:                        {Type: datapoint.Counter},
	goMemstatsMcacheInuseBytes:                    {Type: datapoint.Gauge},
	goMemstatsMcacheSysBytes:                      {Type: datapoint.Gauge},
	goMemstatsMspanInuseBytes:                     {Type: datapoint.Gauge},
	goMemstatsMspanSysBytes:                       {Type: datapoint.Gauge},
	goMemstatsNextGcBytes:                         {Type: datapoint.Gauge},
	goMemstatsOtherSysBytes:                       {Type: datapoint.Gauge},
	goMemstatsStackInuseBytes:                     {Type: datapoint.Gauge},
	goMemstatsStackSysBytes:                       {Type: datapoint.Gauge},
	goMemstatsSysBytes:                            {Type: datapoint.Gauge},
	goThreads:                                     {Type: datapoint.Gauge},
	processCPUSecondsTotal:                        {Type: datapoint.Counter},
	processMaxFds:                                 {Type: datapoint.Gauge},
	processOpenFds:                                {Type: datapoint.Gauge},
	processResidentMemoryBytes:                    {Type: datapoint.Gauge},
	processStartTimeSeconds:                       {Type: datapoint.Gauge},
	processVirtualMemoryBytes:                     {Type: datapoint.Gauge},
	traefikBackendOpenConnections:                 {Type: datapoint.Gauge},
	traefikBackendRequestDurationSecondsBucket:    {Type: datapoint.Counter},
	traefikBackendRequestDurationSecondsCount:     {Type: datapoint.Counter},
	traefikBackendRequestDurationSecondsSum:       {Type: datapoint.Counter},
	traefikBackendRequestsTotal:                   {Type: datapoint.Counter},
	traefikBackendServerUp:                        {Type: datapoint.Gauge},
	traefikConfigLastReloadFailure:                {Type: datapoint.Gauge},
	traefikConfigLastReloadSuccess:                {Type: datapoint.Gauge},
	traefikConfigReloadsFailureTotal:              {Type: datapoint.Counter},
	traefikConfigReloadsTotal:                     {Type: datapoint.Counter},
	traefikEntrypointOpenConnections:              {Type: datapoint.Gauge},
	traefikEntrypointRequestDurationSecondsBucket: {Type: datapoint.Counter},
	traefikEntrypointRequestDurationSecondsCount:  {Type: datapoint.Counter},
	traefikEntrypointRequestDurationSecondsSum:    {Type: datapoint.Counter},
	traefikEntrypointRequestsTotal:                {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	traefikBackendOpenConnections:                true,
	traefikBackendRequestDurationSecondsSum:      true,
	traefikBackendRequestsTotal:                  true,
	traefikBackendServerUp:                       true,
	traefikEntrypointOpenConnections:             true,
	traefikEntrypointRequestDurationSecondsCount: true,
	traefikEntrypointRequestDurationSecondsSum:   true,
	traefikEntrypointRequestsTotal:               true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "traefik",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}