package collector

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//CPUTimings json structure
type CPUTimings struct {
	Ticks float64 `json:"ticks"`
	Time  struct {
		MS float64 `json:"ms"`
	} `json:"time"`
	Value float64 `json:"value"`
}

//BeatStats json structure
type BeatStats struct {
	CPU struct {
		System CPUTimings `json:"system"`
		Total  CPUTimings `json:"total"`
		User   CPUTimings `json:"user"`
	} `json:"cpu"`
	Cgroup struct {
		CPU struct {
			CFS struct {
				Period struct{ Us float64 `json:"us"` } `json:"period"`
				Quota  struct{ Us float64 `json:"us"` } `json:"quota"`
			} `json:"cfs"`
			Stats struct {
				Periods   float64 `json:"periods"`
				Throttled struct {
					Ns      float64 `json:"ns"`
					Periods float64 `json:"periods"`
				} `json:"throttled"`
			} `json:"stats"`
		} `json:"cpu"`
		Cpuacct struct {
			Total struct{ Ns float64 `json:"ns"` } `json:"total"`
		} `json:"cpuacct"`
		Memory struct {
			Mem struct {
				Limit struct{ Bytes float64 `json:"bytes"` } `json:"limit"`
				Usage struct{ Bytes float64 `json:"bytes"` } `json:"usage"`
			} `json:"mem"`
		} `json:"memory"`
	} `json:"cgroup"`
	Handles struct {
		Limit struct {
			Hard float64 `json:"hard"`
			Soft float64 `json:"soft"`
		} `json:"limit"`
		Open float64 `json:"open"`
	} `json:"handles"`
	Info struct {
		Uptime struct {
			MS float64 `json:"ms"`
		} `json:"uptime"`
		EphemeralID string `json:"ephemeral_id"`
		Name        string `json:"name"`
		Version     string `json:"version"`
	} `json:"info"`
	Memstats struct {
		GCNext      float64 `json:"gc_next"`
		MemoryAlloc float64 `json:"memory_alloc"`
		MemorySys   float64 `json:"memory_sys"`
		MemoryTotal float64 `json:"memory_total"`
		RSS         float64 `json:"rss"`
	} `json:"memstats"`

	Runtime struct {
		Goroutines uint64 `json:"goroutines"`
	} `json:"runtime"`
}

type beatCollector struct {
	beatInfo *BeatInfo
	stats    *Stats
	metrics  exportedMetrics
	infoDesc *prometheus.Desc
}

// NewBeatCollector constructor
func NewBeatCollector(beatInfo *BeatInfo, stats *Stats) prometheus.Collector {
	bc := &beatCollector{
		beatInfo: beatInfo,
		stats:    stats,
		infoDesc: prometheus.NewDesc(
			prometheus.BuildFQName(beatInfo.Beat, "", "info"),
			"beat.info",
			[]string{"ephemeral_id", "name", "version"}, nil,
		),
		metrics: exportedMetrics{
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cpu_time", "seconds_total"),
					"beat.cpu.time",
					nil, prometheus.Labels{"mode": "system"},
				),
				eval: func(stats *Stats) float64 {
					return (time.Duration(stats.Beat.CPU.System.Time.MS) * time.Millisecond).Seconds()
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cpu_time", "seconds_total"),
					"beat.cpu.time",
					nil, prometheus.Labels{"mode": "user"},
				),
				eval: func(stats *Stats) float64 {
					return (time.Duration(stats.Beat.CPU.User.Time.MS) * time.Millisecond).Seconds()
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cpu", "ticks_total"),
					"beat.cpu.ticks",
					nil, prometheus.Labels{"mode": "system"},
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.CPU.System.Ticks },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cpu", "ticks_total"),
					"beat.cpu.ticks",
					nil, prometheus.Labels{"mode": "user"},
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.CPU.User.Ticks },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "uptime", "seconds_total"),
					"beat.info.uptime.ms",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return (time.Duration(stats.Beat.Info.Uptime.MS) * time.Millisecond).Seconds()
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "memstats", "gc_next_total"),
					"beat.memstats.gc_next",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.Beat.Memstats.GCNext
				},
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "memstats", "memory_alloc"),
					"beat.memstats.memory_alloc",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.Beat.Memstats.MemoryAlloc
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "memstats", "memory"),
					"beat.memstats.memory_total",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.Beat.Memstats.MemoryTotal
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "memstats", "rss"),
					"beat.memstats.rss",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return stats.Beat.Memstats.RSS
				},
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "runtime", "goroutines"),
					"beat.runtime.goroutines",
					nil, nil,
				),
				eval: func(stats *Stats) float64 {
					return float64(stats.Beat.Runtime.Goroutines)
				},
				valType: prometheus.GaugeValue,
			},
			// Task 3.1: cgroup CPU fields
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_cpu_cfs", "period_microseconds"),
					"beat.cgroup.cpu.cfs.period.us",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.CPU.CFS.Period.Us },
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_cpu_cfs", "quota_microseconds"),
					"beat.cgroup.cpu.cfs.quota.us",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.CPU.CFS.Quota.Us },
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_cpu_stats", "periods_total"),
					"beat.cgroup.cpu.stats.periods",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.CPU.Stats.Periods },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_cpu", "throttled_nanoseconds_total"),
					"beat.cgroup.cpu.stats.throttled.ns",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.CPU.Stats.Throttled.Ns },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_cpu", "throttled_periods_total"),
					"beat.cgroup.cpu.stats.throttled.periods",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.CPU.Stats.Throttled.Periods },
				valType: prometheus.CounterValue,
			},
			// Task 3.2: cgroup cpuacct and memory
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_cpuacct", "nanoseconds_total"),
					"beat.cgroup.cpuacct.total.ns",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.Cpuacct.Total.Ns },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_memory", "limit_bytes"),
					"beat.cgroup.memory.mem.limit.bytes",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.Memory.Mem.Limit.Bytes },
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "cgroup_memory", "usage_bytes"),
					"beat.cgroup.memory.mem.usage.bytes",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Cgroup.Memory.Mem.Usage.Bytes },
				valType: prometheus.GaugeValue,
			},
			// Task 3.3: handles and memstats.memory_sys
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "handles", "limit_hard"),
					"beat.handles.limit.hard",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Handles.Limit.Hard },
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "handles", "limit_soft"),
					"beat.handles.limit.soft",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Handles.Limit.Soft },
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "handles", "open"),
					"beat.handles.open",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Handles.Open },
				valType: prometheus.GaugeValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "memstats", "memory_sys"),
					"beat.memstats.memory_sys",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Beat.Memstats.MemorySys },
				valType: prometheus.GaugeValue,
			},
		},
	}

	return bc
}

// Describe returns all descriptions of the collector.
func (c *beatCollector) Describe(ch chan<- *prometheus.Desc) {

	for _, metric := range c.metrics {
		ch <- metric.desc
	}
	ch <- c.infoDesc

}

// Collect returns the current state of all metrics of the collector.
func (c *beatCollector) Collect(ch chan<- prometheus.Metric) {

	for _, i := range c.metrics {
		ch <- prometheus.MustNewConstMetric(i.desc, i.valType, i.eval(c.stats))
	}

	ch <- prometheus.MustNewConstMetric(c.infoDesc, prometheus.GaugeValue, 1,
		c.stats.Beat.Info.EphemeralID, c.stats.Beat.Info.Name, c.stats.Beat.Info.Version)

}
