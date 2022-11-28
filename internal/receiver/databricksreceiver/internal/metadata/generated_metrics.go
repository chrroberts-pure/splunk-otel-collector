// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledProvidedByUser bool
}

// IsEnabledProvidedByUser returns true if `enabled` option is explicitly set in user settings to any value.
func (ms *MetricSettings) IsEnabledProvidedByUser() bool {
	return ms.enabledProvidedByUser
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledProvidedByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for databricksreceiver metrics.
type MetricsSettings struct {
	DatabricksJobsActiveTotal     MetricSettings `mapstructure:"databricks.jobs.active.total"`
	DatabricksJobsRunDuration     MetricSettings `mapstructure:"databricks.jobs.run.duration"`
	DatabricksJobsScheduleStatus  MetricSettings `mapstructure:"databricks.jobs.schedule.status"`
	DatabricksJobsTotal           MetricSettings `mapstructure:"databricks.jobs.total"`
	DatabricksTasksRunDuration    MetricSettings `mapstructure:"databricks.tasks.run.duration"`
	DatabricksTasksScheduleStatus MetricSettings `mapstructure:"databricks.tasks.schedule.status"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		DatabricksJobsActiveTotal: MetricSettings{
			Enabled: true,
		},
		DatabricksJobsRunDuration: MetricSettings{
			Enabled: true,
		},
		DatabricksJobsScheduleStatus: MetricSettings{
			Enabled: true,
		},
		DatabricksJobsTotal: MetricSettings{
			Enabled: true,
		},
		DatabricksTasksRunDuration: MetricSettings{
			Enabled: true,
		},
		DatabricksTasksScheduleStatus: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeTaskType specifies the a value task_type attribute.
type AttributeTaskType int

const (
	_ AttributeTaskType = iota
	AttributeTaskTypeNotebookTask
	AttributeTaskTypeSparkJarTask
	AttributeTaskTypeSparkPythonTask
	AttributeTaskTypePipelineTask
	AttributeTaskTypePythonWheelTask
	AttributeTaskTypeSparkSubmitTask
)

// String returns the string representation of the AttributeTaskType.
func (av AttributeTaskType) String() string {
	switch av {
	case AttributeTaskTypeNotebookTask:
		return "NotebookTask"
	case AttributeTaskTypeSparkJarTask:
		return "SparkJarTask"
	case AttributeTaskTypeSparkPythonTask:
		return "SparkPythonTask"
	case AttributeTaskTypePipelineTask:
		return "PipelineTask"
	case AttributeTaskTypePythonWheelTask:
		return "PythonWheelTask"
	case AttributeTaskTypeSparkSubmitTask:
		return "SparkSubmitTask"
	}
	return ""
}

// MapAttributeTaskType is a helper map of string to AttributeTaskType attribute value.
var MapAttributeTaskType = map[string]AttributeTaskType{
	"NotebookTask":    AttributeTaskTypeNotebookTask,
	"SparkJarTask":    AttributeTaskTypeSparkJarTask,
	"SparkPythonTask": AttributeTaskTypeSparkPythonTask,
	"PipelineTask":    AttributeTaskTypePipelineTask,
	"PythonWheelTask": AttributeTaskTypePythonWheelTask,
	"SparkSubmitTask": AttributeTaskTypeSparkSubmitTask,
}

type metricDatabricksJobsActiveTotal struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills databricks.jobs.active.total metric with initial data.
func (m *metricDatabricksJobsActiveTotal) init() {
	m.data.SetName("databricks.jobs.active.total")
	m.data.SetDescription("A snapshot of the number of active jobs taken at each scrape")
	m.data.SetUnit("{jobs}")
	m.data.SetEmptyGauge()
}

func (m *metricDatabricksJobsActiveTotal) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDatabricksJobsActiveTotal) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDatabricksJobsActiveTotal) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDatabricksJobsActiveTotal(settings MetricSettings) metricDatabricksJobsActiveTotal {
	m := metricDatabricksJobsActiveTotal{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDatabricksJobsRunDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills databricks.jobs.run.duration metric with initial data.
func (m *metricDatabricksJobsRunDuration) init() {
	m.data.SetName("databricks.jobs.run.duration")
	m.data.SetDescription("The execution duration in milliseconds per completed job")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDatabricksJobsRunDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, jobIDAttributeValue int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutInt("job_id", jobIDAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDatabricksJobsRunDuration) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDatabricksJobsRunDuration) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDatabricksJobsRunDuration(settings MetricSettings) metricDatabricksJobsRunDuration {
	m := metricDatabricksJobsRunDuration{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDatabricksJobsScheduleStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills databricks.jobs.schedule.status metric with initial data.
func (m *metricDatabricksJobsScheduleStatus) init() {
	m.data.SetName("databricks.jobs.schedule.status")
	m.data.SetDescription("A snapshot of the pause/run status per job taken at each scrape")
	m.data.SetUnit("{status}")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDatabricksJobsScheduleStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, jobIDAttributeValue int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutInt("job_id", jobIDAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDatabricksJobsScheduleStatus) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDatabricksJobsScheduleStatus) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDatabricksJobsScheduleStatus(settings MetricSettings) metricDatabricksJobsScheduleStatus {
	m := metricDatabricksJobsScheduleStatus{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDatabricksJobsTotal struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills databricks.jobs.total metric with initial data.
func (m *metricDatabricksJobsTotal) init() {
	m.data.SetName("databricks.jobs.total")
	m.data.SetDescription("A snapshot of the total number of jobs registered in the Databricks instance taken at each scrape")
	m.data.SetUnit("{jobs}")
	m.data.SetEmptyGauge()
}

func (m *metricDatabricksJobsTotal) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDatabricksJobsTotal) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDatabricksJobsTotal) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDatabricksJobsTotal(settings MetricSettings) metricDatabricksJobsTotal {
	m := metricDatabricksJobsTotal{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDatabricksTasksRunDuration struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills databricks.tasks.run.duration metric with initial data.
func (m *metricDatabricksTasksRunDuration) init() {
	m.data.SetName("databricks.tasks.run.duration")
	m.data.SetDescription("The execution duration in milliseconds per completed task")
	m.data.SetUnit("ms")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDatabricksTasksRunDuration) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, jobIDAttributeValue int64, taskIDAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutInt("job_id", jobIDAttributeValue)
	dp.Attributes().PutStr("task_id", taskIDAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDatabricksTasksRunDuration) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDatabricksTasksRunDuration) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDatabricksTasksRunDuration(settings MetricSettings) metricDatabricksTasksRunDuration {
	m := metricDatabricksTasksRunDuration{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricDatabricksTasksScheduleStatus struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills databricks.tasks.schedule.status metric with initial data.
func (m *metricDatabricksTasksScheduleStatus) init() {
	m.data.SetName("databricks.tasks.schedule.status")
	m.data.SetDescription("A snapshot of the pause/run status per task taken at each scrape")
	m.data.SetUnit("{status}")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricDatabricksTasksScheduleStatus) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, jobIDAttributeValue int64, taskIDAttributeValue string, taskTypeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutInt("job_id", jobIDAttributeValue)
	dp.Attributes().PutStr("task_id", taskIDAttributeValue)
	dp.Attributes().PutStr("task_type", taskTypeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricDatabricksTasksScheduleStatus) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricDatabricksTasksScheduleStatus) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricDatabricksTasksScheduleStatus(settings MetricSettings) metricDatabricksTasksScheduleStatus {
	m := metricDatabricksTasksScheduleStatus{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                           pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                     int                 // maximum observed number of metrics per resource.
	resourceCapacity                    int                 // maximum observed number of resource attributes.
	metricsBuffer                       pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                           component.BuildInfo // contains version information
	metricDatabricksJobsActiveTotal     metricDatabricksJobsActiveTotal
	metricDatabricksJobsRunDuration     metricDatabricksJobsRunDuration
	metricDatabricksJobsScheduleStatus  metricDatabricksJobsScheduleStatus
	metricDatabricksJobsTotal           metricDatabricksJobsTotal
	metricDatabricksTasksRunDuration    metricDatabricksTasksRunDuration
	metricDatabricksTasksScheduleStatus metricDatabricksTasksScheduleStatus
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, buildInfo component.BuildInfo, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                           pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                       pmetric.NewMetrics(),
		buildInfo:                           buildInfo,
		metricDatabricksJobsActiveTotal:     newMetricDatabricksJobsActiveTotal(settings.DatabricksJobsActiveTotal),
		metricDatabricksJobsRunDuration:     newMetricDatabricksJobsRunDuration(settings.DatabricksJobsRunDuration),
		metricDatabricksJobsScheduleStatus:  newMetricDatabricksJobsScheduleStatus(settings.DatabricksJobsScheduleStatus),
		metricDatabricksJobsTotal:           newMetricDatabricksJobsTotal(settings.DatabricksJobsTotal),
		metricDatabricksTasksRunDuration:    newMetricDatabricksTasksRunDuration(settings.DatabricksTasksRunDuration),
		metricDatabricksTasksScheduleStatus: newMetricDatabricksTasksScheduleStatus(settings.DatabricksTasksScheduleStatus),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithDatabricksInstanceName sets provided value as "databricks.instance.name" attribute for current resource.
func WithDatabricksInstanceName(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("databricks.instance.name", val)
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/databricksreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricDatabricksJobsActiveTotal.emit(ils.Metrics())
	mb.metricDatabricksJobsRunDuration.emit(ils.Metrics())
	mb.metricDatabricksJobsScheduleStatus.emit(ils.Metrics())
	mb.metricDatabricksJobsTotal.emit(ils.Metrics())
	mb.metricDatabricksTasksRunDuration.emit(ils.Metrics())
	mb.metricDatabricksTasksScheduleStatus.emit(ils.Metrics())
	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordDatabricksJobsActiveTotalDataPoint adds a data point to databricks.jobs.active.total metric.
func (mb *MetricsBuilder) RecordDatabricksJobsActiveTotalDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDatabricksJobsActiveTotal.recordDataPoint(mb.startTime, ts, val)
}

// RecordDatabricksJobsRunDurationDataPoint adds a data point to databricks.jobs.run.duration metric.
func (mb *MetricsBuilder) RecordDatabricksJobsRunDurationDataPoint(ts pcommon.Timestamp, val int64, jobIDAttributeValue int64) {
	mb.metricDatabricksJobsRunDuration.recordDataPoint(mb.startTime, ts, val, jobIDAttributeValue)
}

// RecordDatabricksJobsScheduleStatusDataPoint adds a data point to databricks.jobs.schedule.status metric.
func (mb *MetricsBuilder) RecordDatabricksJobsScheduleStatusDataPoint(ts pcommon.Timestamp, val int64, jobIDAttributeValue int64) {
	mb.metricDatabricksJobsScheduleStatus.recordDataPoint(mb.startTime, ts, val, jobIDAttributeValue)
}

// RecordDatabricksJobsTotalDataPoint adds a data point to databricks.jobs.total metric.
func (mb *MetricsBuilder) RecordDatabricksJobsTotalDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricDatabricksJobsTotal.recordDataPoint(mb.startTime, ts, val)
}

// RecordDatabricksTasksRunDurationDataPoint adds a data point to databricks.tasks.run.duration metric.
func (mb *MetricsBuilder) RecordDatabricksTasksRunDurationDataPoint(ts pcommon.Timestamp, val int64, jobIDAttributeValue int64, taskIDAttributeValue string) {
	mb.metricDatabricksTasksRunDuration.recordDataPoint(mb.startTime, ts, val, jobIDAttributeValue, taskIDAttributeValue)
}

// RecordDatabricksTasksScheduleStatusDataPoint adds a data point to databricks.tasks.schedule.status metric.
func (mb *MetricsBuilder) RecordDatabricksTasksScheduleStatusDataPoint(ts pcommon.Timestamp, val int64, jobIDAttributeValue int64, taskIDAttributeValue string, taskTypeAttributeValue AttributeTaskType) {
	mb.metricDatabricksTasksScheduleStatus.recordDataPoint(mb.startTime, ts, val, jobIDAttributeValue, taskIDAttributeValue, taskTypeAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
