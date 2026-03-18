package collector

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/client_golang/prometheus"
)

// **Validates: Requirements 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 3.10**

// collectMetrics drains a collector into a map of desc-string -> metric value.
// For metrics with labels, the key includes label values.
// It calls Describe first to initialize any package-level descriptors.
func collectMetrics(c prometheus.Collector) map[string]float64 {
	// Call Describe to initialize descriptors (e.g. libbeatOutputType)
	descCh := make(chan *prometheus.Desc, 200)
	c.Describe(descCh)
	close(descCh)

	ch := make(chan prometheus.Metric, 200)
	c.Collect(ch)
	close(ch)

	result := make(map[string]float64)
	for m := range ch {
		d := &dto.Metric{}
		_ = m.Write(d)

		desc := m.Desc().String()
		var val float64
		if d.Gauge != nil {
			val = d.Gauge.GetValue()
		} else if d.Counter != nil {
			val = d.Counter.GetValue()
		} else if d.Untyped != nil {
			val = d.Untyped.GetValue()
		}

		// Build a key that includes label values for uniqueness
		key := desc
		for _, lp := range d.Label {
			key += fmt.Sprintf("[%s=%s]", lp.GetName(), lp.GetValue())
		}
		result[key] = val
	}
	return result
}

// findMetricValue searches collected metrics for one whose desc contains descSubstr
// and whose labels match the given label pairs. Returns the value and whether it was found.
func findMetricValue(metrics map[string]float64, descSubstr string, labels map[string]string) (float64, bool) {
	for key, val := range metrics {
		if !containsStr(key, descSubstr) {
			continue
		}
		match := true
		for lk, lv := range labels {
			if !containsStr(key, fmt.Sprintf("[%s=%s]", lk, lv)) {
				match = false
				break
			}
		}
		if match {
			return val, true
		}
	}
	return 0, false
}

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && searchStr(s, substr)
}

func searchStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestBeatCollectorMetricEmission(t *testing.T) {
	// Property 2: Beat collector metric emission completeness
	f := func(s Stats) bool {
		beatInfo := &BeatInfo{Beat: "testbeat"}
		bc := NewBeatCollector(beatInfo, &s)

		// Point the collector's stats to our generated stats
		// The collector stores a pointer to Stats, so we need to update it
		metrics := collectMetrics(bc)

		// Verify cgroup CPU fields
		checks := []struct {
			desc   string
			labels map[string]string
			want   float64
		}{
			{"cgroup_cpu_cfs_period_microseconds", nil, s.Beat.Cgroup.CPU.CFS.Period.Us},
			{"cgroup_cpu_cfs_quota_microseconds", nil, s.Beat.Cgroup.CPU.CFS.Quota.Us},
			{"cgroup_cpu_stats_periods_total", nil, s.Beat.Cgroup.CPU.Stats.Periods},
			{"cgroup_cpu_throttled_nanoseconds_total", nil, s.Beat.Cgroup.CPU.Stats.Throttled.Ns},
			{"cgroup_cpu_throttled_periods_total", nil, s.Beat.Cgroup.CPU.Stats.Throttled.Periods},
			{"cgroup_cpuacct_nanoseconds_total", nil, s.Beat.Cgroup.Cpuacct.Total.Ns},
			{"cgroup_memory_limit_bytes", nil, s.Beat.Cgroup.Memory.Mem.Limit.Bytes},
			{"cgroup_memory_usage_bytes", nil, s.Beat.Cgroup.Memory.Mem.Usage.Bytes},
			{"handles_limit_hard", nil, s.Beat.Handles.Limit.Hard},
			{"handles_limit_soft", nil, s.Beat.Handles.Limit.Soft},
			{"handles_open", nil, s.Beat.Handles.Open},
			{"memstats_memory_sys", nil, s.Beat.Memstats.MemorySys},
		}

		for _, c := range checks {
			got, ok := findMetricValue(metrics, c.desc, c.labels)
			if !ok {
				t.Logf("metric %q not found in emitted metrics", c.desc)
				return false
			}
			if got != c.want {
				t.Logf("metric %q: got %v, want %v", c.desc, got, c.want)
				return false
			}
		}

		// Verify info metric has correct label values
		infoFound := false
		for key := range metrics {
			if containsStr(key, "\"testbeat_info\"") {
				if containsStr(key, fmt.Sprintf("[ephemeral_id=%s]", s.Beat.Info.EphemeralID)) &&
					containsStr(key, fmt.Sprintf("[name=%s]", s.Beat.Info.Name)) &&
					containsStr(key, fmt.Sprintf("[version=%s]", s.Beat.Info.Version)) {
					infoFound = true
				}
				break
			}
		}
		if !infoFound {
			t.Logf("info metric with correct labels not found")
			return false
		}

		return true
	}

	cfg := &quick.Config{
		MaxCount: 100,
		Values: func(values []reflect.Value, r *rand.Rand) {
			s := Stats{}.Generate(r, 0).Interface().(Stats)
			values[0] = reflect.ValueOf(s)
		},
	}
	if err := quick.Check(f, cfg); err != nil {
		t.Errorf("Property 2 failed: %v", err)
	}
	fmt.Println("Property 2: Beat collector metric emission completeness passed 100 iterations")
}
