package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

// LibBeat json structure
type LibBeat struct {
	Config struct {
		Module struct {
			Running float64 `json:"running"`
			Starts  float64 `json:"starts"`
			Stops   float64 `json:"stops"`
		} `json:"module"`
		Reloads float64 `json:"reloads"`
		Scans   float64 `json:"scans"`
	} `json:"config"`
	Output   LibBeatOutput   `json:"output"`
	Pipeline LibBeatPipeline `json:"pipeline"`
}

// LibBeatEvents json structure (used by pipeline)
type LibBeatEvents struct {
	Acked      float64 `json:"acked"`
	Active     float64 `json:"active"`
	Batches    float64 `json:"batches"`
	Dropped    float64 `json:"dropped"`
	Duplicates float64 `json:"duplicates"`
	Failed     float64 `json:"failed"`
	Filtered   float64 `json:"filtered"`
	Published  float64 `json:"published"`
	Retry      float64 `json:"retry"`
	Total      float64 `json:"total"`
}

// LibBeatOutputEvents json structure (used by output, extends LibBeatEvents fields)
type LibBeatOutputEvents struct {
	Acked      float64 `json:"acked"`
	Active     float64 `json:"active"`
	Batches    float64 `json:"batches"`
	DeadLetter float64 `json:"dead_letter"`
	Dropped    float64 `json:"dropped"`
	Duplicates float64 `json:"duplicates"`
	Failed     float64 `json:"failed"`
	Filtered   float64 `json:"filtered"`
	Published  float64 `json:"published"`
	Retry      float64 `json:"retry"`
	Toomany    float64 `json:"toomany"`
	Total      float64 `json:"total"`
}

// LibBeatOutputBytesErrors json structure
type LibBeatOutputBytesErrors struct {
	Bytes  float64 `json:"bytes"`
	Errors float64 `json:"errors"`
}

// LatencyHistogram json structure for output.write.latency.histogram
type LatencyHistogram struct {
	Count  float64 `json:"count"`
	Max    float64 `json:"max"`
	Mean   float64 `json:"mean"`
	Median float64 `json:"median"`
	Min    float64 `json:"min"`
	P75    float64 `json:"p75"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
	P999   float64 `json:"p999"`
	Stddev float64 `json:"stddev"`
}

// LibBeatOutput json structure
type LibBeatOutput struct {
	Batches struct {
		Split float64 `json:"split"`
	} `json:"batches"`
	Errors float64                  `json:"errors"`
	Events LibBeatOutputEvents      `json:"events"`
	Read   LibBeatOutputBytesErrors `json:"read"`
	Write  struct {
		LibBeatOutputBytesErrors
		Latency struct {
			Histogram LatencyHistogram `json:"histogram"`
		} `json:"latency"`
	} `json:"write"`
	Type string `json:"type"`
}

// LibBeatQueue json structure for pipeline.queue
type LibBeatQueue struct {
	Acked float64 `json:"acked"`
	Added struct {
		Bytes  float64 `json:"bytes"`
		Events float64 `json:"events"`
	} `json:"added"`
	Consumed struct {
		Bytes  float64 `json:"bytes"`
		Events float64 `json:"events"`
	} `json:"consumed"`
	Filled struct {
		Bytes  float64 `json:"bytes"`
		Events float64 `json:"events"`
		Pct    float64 `json:"pct"`
	} `json:"filled"`
	MaxBytes  float64 `json:"max_bytes"`
	MaxEvents float64 `json:"max_events"`
	Removed   struct {
		Bytes  float64 `json:"bytes"`
		Events float64 `json:"events"`
	} `json:"removed"`
}

// LibBeatPipeline json structure
type LibBeatPipeline struct {
	Clients float64       `json:"clients"`
	Events  LibBeatEvents `json:"events"`
	Queue   LibBeatQueue  `json:"queue"`
}

type libbeatCollector struct {
	beatInfo *BeatInfo
	stats    *Stats
	metrics  exportedMetrics
}

var libbeatOutputType *prometheus.Desc

// NewLibBeatCollector constructor
func NewLibBeatCollector(beatInfo *BeatInfo, stats *Stats) prometheus.Collector {
	return &libbeatCollector{
		beatInfo: beatInfo,
		stats:    stats,
		metrics: exportedMetrics{
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat_config", "reloads_total"),
					"libbeat.config.reloads",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Config.Reloads
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "config"),
					"libbeat.config.module",
					nil, prometheus.Labels{"module": "running"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Config.Module.Running
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "config"),
					"libbeat.config.module",
					nil, prometheus.Labels{"module": "starts"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Config.Module.Starts
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "config"),
					"libbeat.config.module",
					nil, prometheus.Labels{"module": "stops"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Config.Module.Stops
				},
				valType: prometheus.GaugeValue,
			},
			// Task 4.1: config.scans
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat_config", "scans_total"),
					"libbeat.config.scans",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Config.Scans
				},
				valType: prometheus.CounterValue,
			},
			// Task 4.1: output.batches.split
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat_output", "batches_split_total"),
					"libbeat.output.batches.split",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Batches.Split
				},
				valType: prometheus.CounterValue,
			},
			// Task 4.1: output.errors
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat_output", "errors_total"),
					"libbeat.output.errors",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Errors
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_read_bytes_total"),
					"libbeat.output.read.bytes",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Read.Bytes
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_read_errors_total"),
					"libbeat.output.read.errors",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Read.Errors
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_bytes_total"),
					"libbeat.output.write.bytes",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Bytes
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_errors_total"),
					"libbeat.output.write.errors",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Errors
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "acked"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Acked
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "active"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Active
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "batches"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Batches
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "dropped"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Dropped
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "duplicates"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Duplicates
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "failed"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Failed
				},
				valType: prometheus.UntypedValue,
			},
			// Task 4.2: output.events.dead_letter
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "dead_letter"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.DeadLetter
				},
				valType: prometheus.UntypedValue,
			},
			// Task 4.2: output.events.toomany
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "toomany"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Toomany
				},
				valType: prometheus.UntypedValue,
			},
			// Task 4.2: output.events.total
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_events"),
					"libbeat.output.events",
					nil, prometheus.Labels{"type": "total"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Events.Total
				},
				valType: prometheus.UntypedValue,
			},
			// Task 4.3: output.write.latency.histogram - simple gauges
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency_count"),
					"libbeat.output.write.latency.histogram.count",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.Count
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency_max"),
					"libbeat.output.write.latency.histogram.max",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.Max
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency_mean"),
					"libbeat.output.write.latency.histogram.mean",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.Mean
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency_median"),
					"libbeat.output.write.latency.histogram.median",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.Median
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency_min"),
					"libbeat.output.write.latency.histogram.min",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.Min
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency_stddev"),
					"libbeat.output.write.latency.histogram.stddev",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.Stddev
				},
				valType: prometheus.GaugeValue,
			},
			// Task 4.3: output.write.latency.histogram - percentile gauges with quantile label
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency"),
					"libbeat.output.write.latency.histogram quantile",
					nil, prometheus.Labels{"quantile": "0.75"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.P75
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency"),
					"libbeat.output.write.latency.histogram quantile",
					nil, prometheus.Labels{"quantile": "0.95"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.P95
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency"),
					"libbeat.output.write.latency.histogram quantile",
					nil, prometheus.Labels{"quantile": "0.99"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.P99
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "output_write_latency"),
					"libbeat.output.write.latency.histogram quantile",
					nil, prometheus.Labels{"quantile": "0.999"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Output.Write.Latency.Histogram.P999
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_clients"),
					"libbeat.pipeline.clients",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Clients
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue"),
					"libbeat.pipeline.queue",
					nil, prometheus.Labels{"type": "acked"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Acked
				},
				valType: prometheus.UntypedValue,
			},
			// Task 4.5: pipeline.queue expanded fields
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_added_bytes_total"),
					"libbeat.pipeline.queue.added.bytes",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Added.Bytes
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_added_events_total"),
					"libbeat.pipeline.queue.added.events",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Added.Events
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_consumed_bytes_total"),
					"libbeat.pipeline.queue.consumed.bytes",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Consumed.Bytes
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_consumed_events_total"),
					"libbeat.pipeline.queue.consumed.events",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Consumed.Events
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_filled_bytes"),
					"libbeat.pipeline.queue.filled.bytes",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Filled.Bytes
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_filled_events"),
					"libbeat.pipeline.queue.filled.events",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Filled.Events
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_filled_pct"),
					"libbeat.pipeline.queue.filled.pct",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Filled.Pct
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_max_bytes"),
					"libbeat.pipeline.queue.max_bytes",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.MaxBytes
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_max_events"),
					"libbeat.pipeline.queue.max_events",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.MaxEvents
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_removed_bytes_total"),
					"libbeat.pipeline.queue.removed.bytes",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Removed.Bytes
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_queue_removed_events_total"),
					"libbeat.pipeline.queue.removed.events",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Queue.Removed.Events
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_events"),
					"libbeat.pipeline.events",
					nil, prometheus.Labels{"type": "active"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Events.Active
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_events"),
					"libbeat.pipeline.events",
					nil, prometheus.Labels{"type": "dropped"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Events.Dropped
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_events"),
					"libbeat.pipeline.events",
					nil, prometheus.Labels{"type": "failed"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Events.Failed
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_events"),
					"libbeat.pipeline.events",
					nil, prometheus.Labels{"type": "filtered"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Events.Filtered
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_events"),
					"libbeat.pipeline.events",
					nil, prometheus.Labels{"type": "published"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Events.Published
				},
				valType: prometheus.UntypedValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_events"),
					"libbeat.pipeline.events",
					nil, prometheus.Labels{"type": "retry"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Events.Retry
				},
				valType: prometheus.UntypedValue,
			},
			// Task 4.4: pipeline.events.total
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "libbeat", "pipeline_events"),
					"libbeat.pipeline.events",
					nil, prometheus.Labels{"type": "total"},
				),
				eval: func(stats *Stats) float64 {
					return stats.LibBeat.Pipeline.Events.Total
				},
				valType: prometheus.UntypedValue,
			},
		},
	}
}

// Describe returns all descriptions of the collector.
func (c *libbeatCollector) Describe(ch chan<- *prometheus.Desc) {

	for _, metric := range c.metrics {
		ch <- metric.desc
	}

	libbeatOutputType = prometheus.NewDesc(
		prometheus.BuildFQName(c.beatInfo.Beat, "libbeat", "output_total"),
		"libbeat.output.type",
		[]string{"type"}, nil,
	)

	ch <- libbeatOutputType

}

// Collect returns the current state of all metrics of the collector.
func (c *libbeatCollector) Collect(ch chan<- prometheus.Metric) {

	for _, i := range c.metrics {
		ch <- prometheus.MustNewConstMetric(i.desc, i.valType, i.eval(c.stats))
	}

	// output.type with dynamic label
	ch <- prometheus.MustNewConstMetric(libbeatOutputType, prometheus.CounterValue, float64(1), c.stats.LibBeat.Output.Type)

}
