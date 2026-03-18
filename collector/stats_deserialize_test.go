package collector

import (
	"encoding/json"
	"testing"
)

// TestReferencePayloadDeserialization validates that a reference filebeat 8.17.4
// JSON payload deserializes correctly into the Stats struct with all new fields.
// **Validates: Requirement 9.3**
func TestReferencePayloadDeserialization(t *testing.T) {
	payload := `{
  "beat": {
    "cpu": {
      "system": {"ticks": 1234, "time": {"ms": 1234}, "value": 1234},
      "total":  {"ticks": 5678, "time": {"ms": 5678}, "value": 5678},
      "user":   {"ticks": 4444, "time": {"ms": 4444}, "value": 4444}
    },
    "cgroup": {
      "cpu": {
        "cfs": {
          "period": {"us": 100000},
          "quota":  {"us": 200000}
        },
        "stats": {
          "periods": 5000,
          "throttled": {
            "ns": 123456789,
            "periods": 42
          }
        }
      },
      "cpuacct": {
        "total": {"ns": 987654321}
      },
      "memory": {
        "mem": {
          "limit": {"bytes": 1073741824},
          "usage": {"bytes": 536870912}
        }
      }
    },
    "handles": {
      "limit": {
        "hard": 1048576,
        "soft": 65536
      },
      "open": 42
    },
    "info": {
      "uptime": {"ms": 3600000},
      "ephemeral_id": "abc-123-def",
      "name": "filebeat-node01",
      "version": "8.17.4"
    },
    "memstats": {
      "gc_next": 4194304,
      "memory_alloc": 2097152,
      "memory_sys": 10485760,
      "memory_total": 8388608,
      "rss": 6291456
    },
    "runtime": {
      "goroutines": 25
    }
  },
  "libbeat": {
    "config": {
      "module": {"running": 1, "starts": 2, "stops": 0},
      "reloads": 3,
      "scans": 15
    },
    "output": {
      "batches": {"split": 7},
      "errors": 2,
      "events": {
        "acked": 1000,
        "active": 5,
        "batches": 100,
        "dead_letter": 3,
        "dropped": 1,
        "duplicates": 0,
        "failed": 2,
        "filtered": 10,
        "published": 990,
        "retry": 4,
        "toomany": 1,
        "total": 1011
      },
      "read": {"bytes": 50000, "errors": 1},
      "write": {
        "bytes": 100000,
        "errors": 0,
        "latency": {
          "histogram": {
            "count": 500,
            "max": 150.5,
            "mean": 12.3,
            "median": 10.0,
            "min": 0.5,
            "p75": 15.0,
            "p95": 50.0,
            "p99": 100.0,
            "p999": 140.0,
            "stddev": 8.7
          }
        }
      },
      "type": "elasticsearch"
    },
    "pipeline": {
      "clients": 2,
      "events": {
        "acked": 900,
        "active": 10,
        "batches": 90,
        "dropped": 0,
        "duplicates": 0,
        "failed": 1,
        "filtered": 5,
        "published": 895,
        "retry": 2,
        "total": 908
      },
      "queue": {
        "acked": 800,
        "added":    {"bytes": 200000, "events": 1000},
        "consumed": {"bytes": 180000, "events": 900},
        "filled":   {"bytes": 20000, "events": 100, "pct": 0.1},
        "max_bytes": 1048576,
        "max_events": 4096,
        "removed":  {"bytes": 180000, "events": 900}
      }
    }
  },
  "system": {
    "cpu": {"cores": 8},
    "load": {
      "1": 1.5, "5": 1.2, "15": 0.9,
      "norm": {"1": 0.19, "5": 0.15, "15": 0.11}
    }
  },
  "registrar": {
    "writes": {"fail": 0, "success": 500, "total": 500},
    "states": {"cleanup": 10, "current": 50, "update": 490}
  },
  "filebeat": {
    "events": {"active": 5, "added": 1000, "done": 995},
    "harvester": {"closed": 2, "open_files": 10, "running": 10, "skipped": 0, "started": 12},
    "input": {"log": {"files": {"renamed": 1, "truncated": 0}}}
  },
  "metricbeat": {
    "system": {
      "cpu": {"failures": 0, "success": 100},
      "filesystem": {"failures": 0, "success": 100},
      "fsstat": {"failures": 0, "success": 100},
      "load": {"failures": 0, "success": 100},
      "memory": {"failures": 0, "success": 100},
      "network": {"failures": 0, "success": 100},
      "process": {"failures": 0, "success": 100},
      "process_summary": {"failures": 0, "success": 100},
      "uptime": {"failures": 0, "success": 100}
    }
  },
  "auditd": {
    "kernel_lost": 0,
    "reassembler_seq_gaps": 0,
    "received_msgs": 1000,
    "userspace_lost": 0
  }
}`

	var s Stats
	err := json.Unmarshal([]byte(payload), &s)
	if err != nil {
		t.Fatalf("failed to unmarshal reference payload: %v", err)
	}

	// Beat section assertions
	// CPU
	assertEqual(t, "beat.cpu.system.ticks", s.Beat.CPU.System.Ticks, 1234.0)
	assertEqual(t, "beat.cpu.total.ticks", s.Beat.CPU.Total.Ticks, 5678.0)
	assertEqual(t, "beat.cpu.user.ticks", s.Beat.CPU.User.Ticks, 4444.0)

	// Cgroup
	assertEqual(t, "beat.cgroup.cpu.cfs.period.us", s.Beat.Cgroup.CPU.CFS.Period.Us, 100000.0)
	assertEqual(t, "beat.cgroup.cpu.cfs.quota.us", s.Beat.Cgroup.CPU.CFS.Quota.Us, 200000.0)
	assertEqual(t, "beat.cgroup.cpu.stats.periods", s.Beat.Cgroup.CPU.Stats.Periods, 5000.0)
	assertEqual(t, "beat.cgroup.cpu.stats.throttled.ns", s.Beat.Cgroup.CPU.Stats.Throttled.Ns, 123456789.0)
	assertEqual(t, "beat.cgroup.cpu.stats.throttled.periods", s.Beat.Cgroup.CPU.Stats.Throttled.Periods, 42.0)
	assertEqual(t, "beat.cgroup.cpuacct.total.ns", s.Beat.Cgroup.Cpuacct.Total.Ns, 987654321.0)
	assertEqual(t, "beat.cgroup.memory.mem.limit.bytes", s.Beat.Cgroup.Memory.Mem.Limit.Bytes, 1073741824.0)
	assertEqual(t, "beat.cgroup.memory.mem.usage.bytes", s.Beat.Cgroup.Memory.Mem.Usage.Bytes, 536870912.0)

	// Handles
	assertEqual(t, "beat.handles.limit.hard", s.Beat.Handles.Limit.Hard, 1048576.0)
	assertEqual(t, "beat.handles.limit.soft", s.Beat.Handles.Limit.Soft, 65536.0)
	assertEqual(t, "beat.handles.open", s.Beat.Handles.Open, 42.0)

	// Info
	assertEqual(t, "beat.info.uptime.ms", s.Beat.Info.Uptime.MS, 3600000.0)
	assertEqualStr(t, "beat.info.ephemeral_id", s.Beat.Info.EphemeralID, "abc-123-def")
	assertEqualStr(t, "beat.info.name", s.Beat.Info.Name, "filebeat-node01")
	assertEqualStr(t, "beat.info.version", s.Beat.Info.Version, "8.17.4")

	// Memstats
	assertEqual(t, "beat.memstats.gc_next", s.Beat.Memstats.GCNext, 4194304.0)
	assertEqual(t, "beat.memstats.memory_alloc", s.Beat.Memstats.MemoryAlloc, 2097152.0)
	assertEqual(t, "beat.memstats.memory_sys", s.Beat.Memstats.MemorySys, 10485760.0)
	assertEqual(t, "beat.memstats.memory_total", s.Beat.Memstats.MemoryTotal, 8388608.0)
	assertEqual(t, "beat.memstats.rss", s.Beat.Memstats.RSS, 6291456.0)

	// Runtime
	if s.Beat.Runtime.Goroutines != 25 {
		t.Errorf("beat.runtime.goroutines: got %d, want 25", s.Beat.Runtime.Goroutines)
	}

	// LibBeat section assertions
	assertEqual(t, "libbeat.config.scans", s.LibBeat.Config.Scans, 15.0)
	assertEqual(t, "libbeat.config.reloads", s.LibBeat.Config.Reloads, 3.0)
	assertEqual(t, "libbeat.output.batches.split", s.LibBeat.Output.Batches.Split, 7.0)
	assertEqual(t, "libbeat.output.errors", s.LibBeat.Output.Errors, 2.0)
	assertEqual(t, "libbeat.output.events.dead_letter", s.LibBeat.Output.Events.DeadLetter, 3.0)
	assertEqual(t, "libbeat.output.events.toomany", s.LibBeat.Output.Events.Toomany, 1.0)
	assertEqual(t, "libbeat.output.events.total", s.LibBeat.Output.Events.Total, 1011.0)
	assertEqual(t, "libbeat.output.events.acked", s.LibBeat.Output.Events.Acked, 1000.0)
	assertEqualStr(t, "libbeat.output.type", s.LibBeat.Output.Type, "elasticsearch")

	// Latency histogram
	assertEqual(t, "libbeat.output.write.latency.histogram.count", s.LibBeat.Output.Write.Latency.Histogram.Count, 500.0)
	assertEqual(t, "libbeat.output.write.latency.histogram.max", s.LibBeat.Output.Write.Latency.Histogram.Max, 150.5)
	assertEqual(t, "libbeat.output.write.latency.histogram.mean", s.LibBeat.Output.Write.Latency.Histogram.Mean, 12.3)
	assertEqual(t, "libbeat.output.write.latency.histogram.median", s.LibBeat.Output.Write.Latency.Histogram.Median, 10.0)
	assertEqual(t, "libbeat.output.write.latency.histogram.min", s.LibBeat.Output.Write.Latency.Histogram.Min, 0.5)
	assertEqual(t, "libbeat.output.write.latency.histogram.p75", s.LibBeat.Output.Write.Latency.Histogram.P75, 15.0)
	assertEqual(t, "libbeat.output.write.latency.histogram.p95", s.LibBeat.Output.Write.Latency.Histogram.P95, 50.0)
	assertEqual(t, "libbeat.output.write.latency.histogram.p99", s.LibBeat.Output.Write.Latency.Histogram.P99, 100.0)
	assertEqual(t, "libbeat.output.write.latency.histogram.p999", s.LibBeat.Output.Write.Latency.Histogram.P999, 140.0)
	assertEqual(t, "libbeat.output.write.latency.histogram.stddev", s.LibBeat.Output.Write.Latency.Histogram.Stddev, 8.7)

	// Pipeline
	assertEqual(t, "libbeat.pipeline.clients", s.LibBeat.Pipeline.Clients, 2.0)
	assertEqual(t, "libbeat.pipeline.events.total", s.LibBeat.Pipeline.Events.Total, 908.0)
	assertEqual(t, "libbeat.pipeline.queue.acked", s.LibBeat.Pipeline.Queue.Acked, 800.0)
	assertEqual(t, "libbeat.pipeline.queue.added.bytes", s.LibBeat.Pipeline.Queue.Added.Bytes, 200000.0)
	assertEqual(t, "libbeat.pipeline.queue.added.events", s.LibBeat.Pipeline.Queue.Added.Events, 1000.0)
	assertEqual(t, "libbeat.pipeline.queue.consumed.bytes", s.LibBeat.Pipeline.Queue.Consumed.Bytes, 180000.0)
	assertEqual(t, "libbeat.pipeline.queue.consumed.events", s.LibBeat.Pipeline.Queue.Consumed.Events, 900.0)
	assertEqual(t, "libbeat.pipeline.queue.filled.bytes", s.LibBeat.Pipeline.Queue.Filled.Bytes, 20000.0)
	assertEqual(t, "libbeat.pipeline.queue.filled.events", s.LibBeat.Pipeline.Queue.Filled.Events, 100.0)
	assertEqual(t, "libbeat.pipeline.queue.filled.pct", s.LibBeat.Pipeline.Queue.Filled.Pct, 0.1)
	assertEqual(t, "libbeat.pipeline.queue.max_bytes", s.LibBeat.Pipeline.Queue.MaxBytes, 1048576.0)
	assertEqual(t, "libbeat.pipeline.queue.max_events", s.LibBeat.Pipeline.Queue.MaxEvents, 4096.0)
	assertEqual(t, "libbeat.pipeline.queue.removed.bytes", s.LibBeat.Pipeline.Queue.Removed.Bytes, 180000.0)
	assertEqual(t, "libbeat.pipeline.queue.removed.events", s.LibBeat.Pipeline.Queue.Removed.Events, 900.0)

	// System
	if s.System.CPU.Cores != 8 {
		t.Errorf("system.cpu.cores: got %d, want 8", s.System.CPU.Cores)
	}
	assertEqual(t, "system.load.1", s.System.Load.M1, 1.5)
}

func assertEqual(t *testing.T, field string, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %v, want %v", field, got, want)
	}
}

func assertEqualStr(t *testing.T, field, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %q, want %q", field, got, want)
	}
}

// TestUnknownFieldsIgnored validates that unknown JSON fields do not cause errors.
// **Validates: Requirement 9.2**
func TestUnknownFieldsIgnored(t *testing.T) {
	payload := `{"beat":{"unknown_field":123,"cpu":{"system":{"ticks":1}}},"unknown_top":true,"libbeat":{"extra_section":{"nested":42}}}`
	var s Stats
	err := json.Unmarshal([]byte(payload), &s)
	if err != nil {
		t.Fatalf("expected no error for unknown fields, got: %v", err)
	}
	// Verify the known field was still parsed
	if s.Beat.CPU.System.Ticks != 1 {
		t.Errorf("beat.cpu.system.ticks: got %v, want 1", s.Beat.CPU.System.Ticks)
	}
}
