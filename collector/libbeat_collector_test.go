package collector

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// **Validates: Requirements 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9**

func TestLibBeatCollectorMetricEmission(t *testing.T) {
	// Property 3: LibBeat collector metric emission completeness
	f := func(s Stats) bool {
		beatInfo := &BeatInfo{Beat: "testbeat"}
		lc := NewLibBeatCollector(beatInfo, &s)

		metrics := collectMetrics(lc)

		checks := []struct {
			desc   string
			labels map[string]string
			want   float64
		}{
			// config.scans
			{"libbeat_config_scans_total", nil, s.LibBeat.Config.Scans},
			// output.batches.split
			{"libbeat_output_batches_split_total", nil, s.LibBeat.Output.Batches.Split},
			// output.errors
			{"libbeat_output_errors_total", nil, s.LibBeat.Output.Errors},
			// output.events new fields
			{"output_events", map[string]string{"type": "dead_letter"}, s.LibBeat.Output.Events.DeadLetter},
			{"output_events", map[string]string{"type": "toomany"}, s.LibBeat.Output.Events.Toomany},
			{"output_events", map[string]string{"type": "total"}, s.LibBeat.Output.Events.Total},
			// latency histogram simple gauges
			{"output_write_latency_count", nil, s.LibBeat.Output.Write.Latency.Histogram.Count},
			{"output_write_latency_max", nil, s.LibBeat.Output.Write.Latency.Histogram.Max},
			{"output_write_latency_mean", nil, s.LibBeat.Output.Write.Latency.Histogram.Mean},
			{"output_write_latency_median", nil, s.LibBeat.Output.Write.Latency.Histogram.Median},
			{"output_write_latency_min", nil, s.LibBeat.Output.Write.Latency.Histogram.Min},
			{"output_write_latency_stddev", nil, s.LibBeat.Output.Write.Latency.Histogram.Stddev},
			// latency histogram percentile gauges with quantile label
			{"output_write_latency\"", map[string]string{"quantile": "0.75"}, s.LibBeat.Output.Write.Latency.Histogram.P75},
			{"output_write_latency\"", map[string]string{"quantile": "0.95"}, s.LibBeat.Output.Write.Latency.Histogram.P95},
			{"output_write_latency\"", map[string]string{"quantile": "0.99"}, s.LibBeat.Output.Write.Latency.Histogram.P99},
			{"output_write_latency\"", map[string]string{"quantile": "0.999"}, s.LibBeat.Output.Write.Latency.Histogram.P999},
			// pipeline.events.total
			{"pipeline_events", map[string]string{"type": "total"}, s.LibBeat.Pipeline.Events.Total},
			// pipeline.queue expanded fields
			{"pipeline_queue_added_bytes_total", nil, s.LibBeat.Pipeline.Queue.Added.Bytes},
			{"pipeline_queue_added_events_total", nil, s.LibBeat.Pipeline.Queue.Added.Events},
			{"pipeline_queue_consumed_bytes_total", nil, s.LibBeat.Pipeline.Queue.Consumed.Bytes},
			{"pipeline_queue_consumed_events_total", nil, s.LibBeat.Pipeline.Queue.Consumed.Events},
			{"pipeline_queue_filled_bytes", nil, s.LibBeat.Pipeline.Queue.Filled.Bytes},
			{"pipeline_queue_filled_events", nil, s.LibBeat.Pipeline.Queue.Filled.Events},
			{"pipeline_queue_filled_pct", nil, s.LibBeat.Pipeline.Queue.Filled.Pct},
			{"pipeline_queue_max_bytes", nil, s.LibBeat.Pipeline.Queue.MaxBytes},
			{"pipeline_queue_max_events", nil, s.LibBeat.Pipeline.Queue.MaxEvents},
			{"pipeline_queue_removed_bytes_total", nil, s.LibBeat.Pipeline.Queue.Removed.Bytes},
			{"pipeline_queue_removed_events_total", nil, s.LibBeat.Pipeline.Queue.Removed.Events},
		}

		for _, c := range checks {
			got, ok := findMetricValue(metrics, c.desc, c.labels)
			if !ok {
				t.Logf("metric %q (labels: %v) not found in emitted metrics", c.desc, c.labels)
				return false
			}
			if got != c.want {
				t.Logf("metric %q (labels: %v): got %v, want %v", c.desc, c.labels, got, c.want)
				return false
			}
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
		t.Errorf("Property 3 failed: %v", err)
	}
	fmt.Println("Property 3: LibBeat collector metric emission completeness passed 100 iterations")
}
