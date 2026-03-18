package collector

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// **Validates: Requirements 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 2.10, 9.1**

// randString generates a random alphanumeric string of length 1-10.
func randString(r *rand.Rand) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	n := r.Intn(10) + 1
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

// Generate implements quick.Generator for Stats so testing/quick can produce random instances.
func (s Stats) Generate(r *rand.Rand, size int) reflect.Value {
	st := Stats{}

	// System
	st.System.CPU.Cores = r.Int63n(128)
	st.System.Load.M1 = r.Float64()
	st.System.Load.M5 = r.Float64()
	st.System.Load.M15 = r.Float64()
	st.System.Load.Norm.M1 = r.Float64()
	st.System.Load.Norm.M5 = r.Float64()
	st.System.Load.Norm.M15 = r.Float64()

	// Beat
	st.Beat.CPU.System.Ticks = r.Float64()
	st.Beat.CPU.System.Time.MS = r.Float64()
	st.Beat.CPU.System.Value = r.Float64()
	st.Beat.CPU.Total.Ticks = r.Float64()
	st.Beat.CPU.Total.Time.MS = r.Float64()
	st.Beat.CPU.Total.Value = r.Float64()
	st.Beat.CPU.User.Ticks = r.Float64()
	st.Beat.CPU.User.Time.MS = r.Float64()
	st.Beat.CPU.User.Value = r.Float64()

	st.Beat.Cgroup.CPU.CFS.Period.Us = r.Float64()
	st.Beat.Cgroup.CPU.CFS.Quota.Us = r.Float64()
	st.Beat.Cgroup.CPU.Stats.Periods = r.Float64()
	st.Beat.Cgroup.CPU.Stats.Throttled.Ns = r.Float64()
	st.Beat.Cgroup.CPU.Stats.Throttled.Periods = r.Float64()
	st.Beat.Cgroup.Cpuacct.Total.Ns = r.Float64()
	st.Beat.Cgroup.Memory.Mem.Limit.Bytes = r.Float64()
	st.Beat.Cgroup.Memory.Mem.Usage.Bytes = r.Float64()

	st.Beat.Handles.Limit.Hard = r.Float64()
	st.Beat.Handles.Limit.Soft = r.Float64()
	st.Beat.Handles.Open = r.Float64()

	st.Beat.Info.Uptime.MS = r.Float64()
	st.Beat.Info.EphemeralID = randString(r)
	st.Beat.Info.Name = randString(r)
	st.Beat.Info.Version = randString(r)

	st.Beat.Memstats.GCNext = r.Float64()
	st.Beat.Memstats.MemoryAlloc = r.Float64()
	st.Beat.Memstats.MemorySys = r.Float64()
	st.Beat.Memstats.MemoryTotal = r.Float64()
	st.Beat.Memstats.RSS = r.Float64()

	st.Beat.Runtime.Goroutines = uint64(r.Int63n(10000))

	// LibBeat - Config
	st.LibBeat.Config.Module.Running = r.Float64()
	st.LibBeat.Config.Module.Starts = r.Float64()
	st.LibBeat.Config.Module.Stops = r.Float64()
	st.LibBeat.Config.Reloads = r.Float64()
	st.LibBeat.Config.Scans = r.Float64()

	// LibBeat - Output
	st.LibBeat.Output.Batches.Split = r.Float64()
	st.LibBeat.Output.Errors = r.Float64()
	st.LibBeat.Output.Events.Acked = r.Float64()
	st.LibBeat.Output.Events.Active = r.Float64()
	st.LibBeat.Output.Events.Batches = r.Float64()
	st.LibBeat.Output.Events.DeadLetter = r.Float64()
	st.LibBeat.Output.Events.Dropped = r.Float64()
	st.LibBeat.Output.Events.Duplicates = r.Float64()
	st.LibBeat.Output.Events.Failed = r.Float64()
	st.LibBeat.Output.Events.Filtered = r.Float64()
	st.LibBeat.Output.Events.Published = r.Float64()
	st.LibBeat.Output.Events.Retry = r.Float64()
	st.LibBeat.Output.Events.Toomany = r.Float64()
	st.LibBeat.Output.Events.Total = r.Float64()
	st.LibBeat.Output.Read.Bytes = r.Float64()
	st.LibBeat.Output.Read.Errors = r.Float64()
	st.LibBeat.Output.Write.Bytes = r.Float64()
	st.LibBeat.Output.Write.Errors = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.Count = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.Max = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.Mean = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.Median = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.Min = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.P75 = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.P95 = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.P99 = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.P999 = r.Float64()
	st.LibBeat.Output.Write.Latency.Histogram.Stddev = r.Float64()
	st.LibBeat.Output.Type = randString(r)

	// LibBeat - Pipeline
	st.LibBeat.Pipeline.Clients = r.Float64()
	st.LibBeat.Pipeline.Events.Acked = r.Float64()
	st.LibBeat.Pipeline.Events.Active = r.Float64()
	st.LibBeat.Pipeline.Events.Batches = r.Float64()
	st.LibBeat.Pipeline.Events.Dropped = r.Float64()
	st.LibBeat.Pipeline.Events.Duplicates = r.Float64()
	st.LibBeat.Pipeline.Events.Failed = r.Float64()
	st.LibBeat.Pipeline.Events.Filtered = r.Float64()
	st.LibBeat.Pipeline.Events.Published = r.Float64()
	st.LibBeat.Pipeline.Events.Retry = r.Float64()
	st.LibBeat.Pipeline.Events.Total = r.Float64()
	st.LibBeat.Pipeline.Queue.Acked = r.Float64()
	st.LibBeat.Pipeline.Queue.Added.Bytes = r.Float64()
	st.LibBeat.Pipeline.Queue.Added.Events = r.Float64()
	st.LibBeat.Pipeline.Queue.Consumed.Bytes = r.Float64()
	st.LibBeat.Pipeline.Queue.Consumed.Events = r.Float64()
	st.LibBeat.Pipeline.Queue.Filled.Bytes = r.Float64()
	st.LibBeat.Pipeline.Queue.Filled.Events = r.Float64()
	st.LibBeat.Pipeline.Queue.Filled.Pct = r.Float64()
	st.LibBeat.Pipeline.Queue.MaxBytes = r.Float64()
	st.LibBeat.Pipeline.Queue.MaxEvents = r.Float64()
	st.LibBeat.Pipeline.Queue.Removed.Bytes = r.Float64()
	st.LibBeat.Pipeline.Queue.Removed.Events = r.Float64()

	// Registrar
	st.Registrar.Writes.Fail = r.Float64()
	st.Registrar.Writes.Success = r.Float64()
	st.Registrar.Writes.Total = r.Float64()
	st.Registrar.States.Cleanup = r.Float64()
	st.Registrar.States.Current = r.Float64()
	st.Registrar.States.Update = r.Float64()

	// Filebeat
	st.Filebeat.Events.Active = r.Float64()
	st.Filebeat.Events.Added = r.Float64()
	st.Filebeat.Events.Done = r.Float64()
	st.Filebeat.Harvester.Closed = r.Float64()
	st.Filebeat.Harvester.OpenFiles = r.Float64()
	st.Filebeat.Harvester.Running = r.Float64()
	st.Filebeat.Harvester.Skipped = r.Float64()
	st.Filebeat.Harvester.Started = r.Float64()
	st.Filebeat.Input.Log.Files.Renamed = r.Float64()
	st.Filebeat.Input.Log.Files.Truncated = r.Float64()

	// Metricbeat
	st.Metricbeat.System.CPU.Failures = r.Float64()
	st.Metricbeat.System.CPU.Success = r.Float64()
	st.Metricbeat.System.Filesystem.Failures = r.Float64()
	st.Metricbeat.System.Filesystem.Success = r.Float64()
	st.Metricbeat.System.Fsstat.Failures = r.Float64()
	st.Metricbeat.System.Fsstat.Success = r.Float64()
	st.Metricbeat.System.Load.Failures = r.Float64()
	st.Metricbeat.System.Load.Success = r.Float64()
	st.Metricbeat.System.Memory.Failures = r.Float64()
	st.Metricbeat.System.Memory.Success = r.Float64()
	st.Metricbeat.System.Network.Failures = r.Float64()
	st.Metricbeat.System.Network.Success = r.Float64()
	st.Metricbeat.System.Process.Failures = r.Float64()
	st.Metricbeat.System.Process.Success = r.Float64()
	st.Metricbeat.System.ProcessSummary.Failures = r.Float64()
	st.Metricbeat.System.ProcessSummary.Success = r.Float64()
	st.Metricbeat.System.Uptime.Failures = r.Float64()
	st.Metricbeat.System.Uptime.Success = r.Float64()

	// Auditd
	st.Auditd.KernelLost = r.Float64()
	st.Auditd.ReassemblerSeqGaps = r.Float64()
	st.Auditd.ReceivedMsgs = r.Float64()
	st.Auditd.UserspaceLost = r.Float64()

	return reflect.ValueOf(st)
}

func TestStatsJSONRoundTrip(t *testing.T) {
	// Property 1: Stats struct JSON round-trip
	f := func(s Stats) bool {
		data, err := json.Marshal(s)
		if err != nil {
			t.Logf("Marshal error: %v", err)
			return false
		}
		var s2 Stats
		if err := json.Unmarshal(data, &s2); err != nil {
			t.Logf("Unmarshal error: %v", err)
			return false
		}
		if !reflect.DeepEqual(s, s2) {
			t.Logf("Round-trip mismatch:\nOriginal: %+v\nDecoded:  %+v", s, s2)
			return false
		}
		return true
	}
	cfg := &quick.Config{MaxCount: 100}
	if err := quick.Check(f, cfg); err != nil {
		t.Errorf("Property 1 failed: %v", err)
	}
	fmt.Println("Property 1: Stats JSON round-trip passed 100 iterations")
}
