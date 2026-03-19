# Beat Exporter Alerting Rules

Prometheus alerting rules for Filebeat metrics exposed by beat_exporter.

## Error Alerts

### FilebeatOutputErrors

Output destination (Elasticsearch, Logstash, etc.) is rejecting or failing to receive events.

```yaml
- alert: FilebeatOutputErrors
  expr: rate(filebeat_libbeat_output_errors_total[5m]) > 0
  for: 5m
  labels:
    severity: critical
  annotations:
    summary: "Filebeat output errors detected ({{ $labels.instance }})"
    description: "Output error rate is {{ $value | printf \"%.2f\" }}/s. Check connectivity and health of the output destination."
```

### FilebeatOutputWriteErrors

Network-level write failures to the output destination.

```yaml
- alert: FilebeatOutputWriteErrors
  expr: rate(filebeat_libbeat_output_write_errors_total[5m]) > 0
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat output write errors ({{ $labels.instance }})"
    description: "Write error rate is {{ $value | printf \"%.2f\" }}/s. Likely network issue or output destination overloaded."
```

### FilebeatOutputReadErrors

Failures reading responses from the output destination.

```yaml
- alert: FilebeatOutputReadErrors
  expr: rate(filebeat_libbeat_output_read_errors_total[5m]) > 0
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat output read errors ({{ $labels.instance }})"
    description: "Read error rate is {{ $value | printf \"%.2f\" }}/s. Output destination may be returning malformed responses or timing out."
```

### FilebeatPipelineEventsFailed

Events failing in the internal pipeline before reaching the output.

```yaml
- alert: FilebeatPipelineEventsFailed
  expr: rate(filebeat_libbeat_pipeline_events{type="failed"}[5m]) > 0
  for: 5m
  labels:
    severity: critical
  annotations:
    summary: "Filebeat pipeline events failing ({{ $labels.instance }})"
    description: "Pipeline failure rate is {{ $value | printf \"%.2f\" }}/s. Check processor configuration and event formatting."
```

### FilebeatDeadLetterEvents

Events that exhausted all retry attempts and can never be delivered.

```yaml
- alert: FilebeatDeadLetterEvents
  expr: filebeat_libbeat_output_events{type="dead_letter"} > 0
  for: 1m
  labels:
    severity: critical
  annotations:
    summary: "Filebeat dead letter events detected ({{ $labels.instance }})"
    description: "{{ $value }} events in dead letter queue. These events are permanently undeliverable. Investigate output rejection reasons."
```

### FilebeatRegistrarWriteFailures

Filebeat cannot persist its harvester state to disk. Risk of duplicate or lost events on restart.

```yaml
- alert: FilebeatRegistrarWriteFailures
  expr: filebeat_registrar_writes{writes="fail"} > 0
  for: 2m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat registrar write failures ({{ $labels.instance }})"
    description: "{{ $value }} registrar write failures. Check disk space and permissions on the Filebeat data directory."
```

## Performance Alerts

### FilebeatQueueSaturated

Internal event queue is nearly full. Backpressure will cause harvesters to stall and new events to be delayed.

```yaml
- alert: FilebeatQueueSaturated
  expr: filebeat_libbeat_pipeline_queue_filled_pct >= 0.9
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat queue saturated at {{ $value | printf \"%.0f\" }}% ({{ $labels.instance }})"
    description: "Queue has been >= 90% full for 5 minutes. Output cannot keep up with input rate. Consider scaling output or increasing queue size."
```

### FilebeatHighWriteLatency

p99 write latency to the output destination is elevated. Threshold set at 5000ms — tune to your SLA.

```yaml
- alert: FilebeatHighWriteLatency
  expr: filebeat_libbeat_output_write_latency{quantile="0.99"} > 5000
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat p99 write latency is {{ $value | printf \"%.0f\" }}ms ({{ $labels.instance }})"
    description: "Output write latency p99 exceeds 5s. Check output destination performance and network latency."
```

### FilebeatMemoryNearCgroupLimit

Memory usage approaching the cgroup limit. OOM kill is imminent if this continues.

```yaml
- alert: FilebeatMemoryNearCgroupLimit
  expr: filebeat_cgroup_memory_usage_bytes / filebeat_cgroup_memory_limit_bytes > 0.85 and filebeat_cgroup_memory_limit_bytes > 0
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat memory at {{ $value | printf \"%.0f\" }}% of cgroup limit ({{ $labels.instance }})"
    description: "Memory usage is above 85% of the cgroup limit. Risk of OOM kill. Consider increasing memory limit or reducing queue/bulk size."
```

### FilebeatCPUThrottling

Filebeat is being CPU-throttled by cgroup limits. Processing lag will increase.

```yaml
- alert: FilebeatCPUThrottling
  expr: rate(filebeat_cgroup_cpu_throttled_periods_total[5m]) / rate(filebeat_cgroup_cpu_stats_periods_total[5m]) > 0.25 and rate(filebeat_cgroup_cpu_stats_periods_total[5m]) > 0
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat CPU throttled {{ $value | printf \"%.0f\" }}% of periods ({{ $labels.instance }})"
    description: "More than 25% of CPU periods are being throttled. Increase CPU limit or reduce processing load."
```

### FilebeatFileHandleExhaustion

Open file handles approaching the soft limit. Filebeat may fail to open new files for harvesting.

```yaml
- alert: FilebeatFileHandleExhaustion
  expr: filebeat_handles_open / filebeat_handles_limit_soft > 0.8
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Filebeat using {{ $value | printf \"%.0f\" }}% of file handle limit ({{ $labels.instance }})"
    description: "Open handles above 80% of soft limit. Check for file handle leaks or increase ulimit."
```
