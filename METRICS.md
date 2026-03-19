# Sample of metrics provisioned by beat-exporter
```txt
# HELP beat_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which beat_exporter was built.
# TYPE beat_exporter_build_info gauge
beat_exporter_build_info{branch="",goversion="go1.26.1",revision="",version=""} 1
# HELP beat_exporter_target_info target information
# TYPE beat_exporter_target_info gauge
beat_exporter_target_info{beat="filebeat",uri="filebeat:5066",version="8.17.4"} 1
# HELP filebeat_auditd_kernel_lost auditd.kernel_lost
# TYPE filebeat_auditd_kernel_lost gauge
filebeat_auditd_kernel_lost 0
# HELP filebeat_auditd_reassembler_seq_gaps auditd.reassembler_seq_gaps
# TYPE filebeat_auditd_reassembler_seq_gaps gauge
filebeat_auditd_reassembler_seq_gaps 0
# HELP filebeat_auditd_received_msgs auditd.received_msgs
# TYPE filebeat_auditd_received_msgs gauge
filebeat_auditd_received_msgs 0
# HELP filebeat_auditd_userspace_lost auditd.userspace_lost
# TYPE filebeat_auditd_userspace_lost gauge
filebeat_auditd_userspace_lost 0
# HELP filebeat_cgroup_cpu_cfs_period_microseconds beat.cgroup.cpu.cfs.period.us
# TYPE filebeat_cgroup_cpu_cfs_period_microseconds gauge
filebeat_cgroup_cpu_cfs_period_microseconds 0
# HELP filebeat_cgroup_cpu_cfs_quota_microseconds beat.cgroup.cpu.cfs.quota.us
# TYPE filebeat_cgroup_cpu_cfs_quota_microseconds gauge
filebeat_cgroup_cpu_cfs_quota_microseconds 0
# HELP filebeat_cgroup_cpu_stats_periods_total beat.cgroup.cpu.stats.periods
# TYPE filebeat_cgroup_cpu_stats_periods_total counter
filebeat_cgroup_cpu_stats_periods_total 0
# HELP filebeat_cgroup_cpu_throttled_nanoseconds_total beat.cgroup.cpu.stats.throttled.ns
# TYPE filebeat_cgroup_cpu_throttled_nanoseconds_total counter
filebeat_cgroup_cpu_throttled_nanoseconds_total 0
# HELP filebeat_cgroup_cpu_throttled_periods_total beat.cgroup.cpu.stats.throttled.periods
# TYPE filebeat_cgroup_cpu_throttled_periods_total counter
filebeat_cgroup_cpu_throttled_periods_total 0
# HELP filebeat_cgroup_cpuacct_nanoseconds_total beat.cgroup.cpuacct.total.ns
# TYPE filebeat_cgroup_cpuacct_nanoseconds_total counter
filebeat_cgroup_cpuacct_nanoseconds_total 0
# HELP filebeat_cgroup_memory_limit_bytes beat.cgroup.memory.mem.limit.bytes
# TYPE filebeat_cgroup_memory_limit_bytes gauge
filebeat_cgroup_memory_limit_bytes 0
# HELP filebeat_cgroup_memory_usage_bytes beat.cgroup.memory.mem.usage.bytes
# TYPE filebeat_cgroup_memory_usage_bytes gauge
filebeat_cgroup_memory_usage_bytes 1.06156032e+08
# HELP filebeat_cpu_ticks_total beat.cpu.ticks
# TYPE filebeat_cpu_ticks_total counter
filebeat_cpu_ticks_total{mode="system"} 940
filebeat_cpu_ticks_total{mode="user"} 5000
# HELP filebeat_cpu_time_seconds_total beat.cpu.time
# TYPE filebeat_cpu_time_seconds_total counter
filebeat_cpu_time_seconds_total{mode="system"} 0.94
filebeat_cpu_time_seconds_total{mode="user"} 5
# HELP filebeat_filebeat_events filebeat.events
# TYPE filebeat_filebeat_events untyped
filebeat_filebeat_events{event="active"} 0
filebeat_filebeat_events{event="added"} 407119
filebeat_filebeat_events{event="done"} 407119
# HELP filebeat_filebeat_harvester filebeat.harvester
# TYPE filebeat_filebeat_harvester untyped
filebeat_filebeat_harvester{harvester="closed"} 0
filebeat_filebeat_harvester{harvester="open_files"} 4
filebeat_filebeat_harvester{harvester="running"} 4
filebeat_filebeat_harvester{harvester="skipped"} 0
filebeat_filebeat_harvester{harvester="started"} 4
# HELP filebeat_filebeat_input_log filebeat.input_log
# TYPE filebeat_filebeat_input_log untyped
filebeat_filebeat_input_log{files="renamed"} 0
filebeat_filebeat_input_log{files="truncated"} 0
# HELP filebeat_handles_limit_hard beat.handles.limit.hard
# TYPE filebeat_handles_limit_hard gauge
filebeat_handles_limit_hard 1.048576e+06
# HELP filebeat_handles_limit_soft beat.handles.limit.soft
# TYPE filebeat_handles_limit_soft gauge
filebeat_handles_limit_soft 1.048576e+06
# HELP filebeat_handles_open beat.handles.open
# TYPE filebeat_handles_open gauge
filebeat_handles_open 14
# HELP filebeat_info beat.info
# TYPE filebeat_info gauge
filebeat_info{ephemeral_id="f354c371-980f-4707-958a-4174fc53122c",name="filebeat",version="8.17.4"} 1
# HELP filebeat_libbeat_config libbeat.config.module
# TYPE filebeat_libbeat_config gauge
filebeat_libbeat_config{module="running"} 0
filebeat_libbeat_config{module="starts"} 0
filebeat_libbeat_config{module="stops"} 0
# HELP filebeat_libbeat_config_reloads_total libbeat.config.reloads
# TYPE filebeat_libbeat_config_reloads_total counter
filebeat_libbeat_config_reloads_total 0
# HELP filebeat_libbeat_config_scans_total libbeat.config.scans
# TYPE filebeat_libbeat_config_scans_total counter
filebeat_libbeat_config_scans_total 0
# HELP filebeat_libbeat_output_batches_split_total libbeat.output.batches.split
# TYPE filebeat_libbeat_output_batches_split_total counter
filebeat_libbeat_output_batches_split_total 0
# HELP filebeat_libbeat_output_errors_total libbeat.output.errors
# TYPE filebeat_libbeat_output_errors_total counter
filebeat_libbeat_output_errors_total 0
# HELP filebeat_libbeat_output_events libbeat.output.events
# TYPE filebeat_libbeat_output_events untyped
filebeat_libbeat_output_events{type="acked"} 389848
filebeat_libbeat_output_events{type="active"} 0
filebeat_libbeat_output_events{type="batches"} 244
filebeat_libbeat_output_events{type="dead_letter"} 0
filebeat_libbeat_output_events{type="dropped"} 0
filebeat_libbeat_output_events{type="duplicates"} 0
filebeat_libbeat_output_events{type="failed"} 0
filebeat_libbeat_output_events{type="toomany"} 0
filebeat_libbeat_output_events{type="total"} 389848
# HELP filebeat_libbeat_output_read_bytes_total libbeat.output.read.bytes
# TYPE filebeat_libbeat_output_read_bytes_total counter
filebeat_libbeat_output_read_bytes_total 0
# HELP filebeat_libbeat_output_read_errors_total libbeat.output.read.errors
# TYPE filebeat_libbeat_output_read_errors_total counter
filebeat_libbeat_output_read_errors_total 0
# HELP filebeat_libbeat_output_total libbeat.output.type
# TYPE filebeat_libbeat_output_total counter
filebeat_libbeat_output_total{type="console"} 1
# HELP filebeat_libbeat_output_write_bytes_total libbeat.output.write.bytes
# TYPE filebeat_libbeat_output_write_bytes_total counter
filebeat_libbeat_output_write_bytes_total 2.58844686e+08
# HELP filebeat_libbeat_output_write_errors_total libbeat.output.write.errors
# TYPE filebeat_libbeat_output_write_errors_total counter
filebeat_libbeat_output_write_errors_total 0
# HELP filebeat_libbeat_output_write_latency libbeat.output.write.latency.histogram quantile
# TYPE filebeat_libbeat_output_write_latency gauge
filebeat_libbeat_output_write_latency{quantile="0.75"} 0
filebeat_libbeat_output_write_latency{quantile="0.95"} 0
filebeat_libbeat_output_write_latency{quantile="0.99"} 0
filebeat_libbeat_output_write_latency{quantile="0.999"} 0
# HELP filebeat_libbeat_output_write_latency_count libbeat.output.write.latency.histogram.count
# TYPE filebeat_libbeat_output_write_latency_count gauge
filebeat_libbeat_output_write_latency_count 0
# HELP filebeat_libbeat_output_write_latency_max libbeat.output.write.latency.histogram.max
# TYPE filebeat_libbeat_output_write_latency_max gauge
filebeat_libbeat_output_write_latency_max 0
# HELP filebeat_libbeat_output_write_latency_mean libbeat.output.write.latency.histogram.mean
# TYPE filebeat_libbeat_output_write_latency_mean gauge
filebeat_libbeat_output_write_latency_mean 0
# HELP filebeat_libbeat_output_write_latency_median libbeat.output.write.latency.histogram.median
# TYPE filebeat_libbeat_output_write_latency_median gauge
filebeat_libbeat_output_write_latency_median 0
# HELP filebeat_libbeat_output_write_latency_min libbeat.output.write.latency.histogram.min
# TYPE filebeat_libbeat_output_write_latency_min gauge
filebeat_libbeat_output_write_latency_min 0
# HELP filebeat_libbeat_output_write_latency_stddev libbeat.output.write.latency.histogram.stddev
# TYPE filebeat_libbeat_output_write_latency_stddev gauge
filebeat_libbeat_output_write_latency_stddev 0
# HELP filebeat_libbeat_pipeline_clients libbeat.pipeline.clients
# TYPE filebeat_libbeat_pipeline_clients gauge
filebeat_libbeat_pipeline_clients 1
# HELP filebeat_libbeat_pipeline_events libbeat.pipeline.events
# TYPE filebeat_libbeat_pipeline_events untyped
filebeat_libbeat_pipeline_events{type="active"} 0
filebeat_libbeat_pipeline_events{type="dropped"} 0
filebeat_libbeat_pipeline_events{type="failed"} 0
filebeat_libbeat_pipeline_events{type="filtered"} 17271
filebeat_libbeat_pipeline_events{type="published"} 389848
filebeat_libbeat_pipeline_events{type="retry"} 0
filebeat_libbeat_pipeline_events{type="total"} 407119
# HELP filebeat_libbeat_pipeline_queue libbeat.pipeline.queue
# TYPE filebeat_libbeat_pipeline_queue untyped
filebeat_libbeat_pipeline_queue{type="acked"} 389848
# HELP filebeat_libbeat_pipeline_queue_added_bytes_total libbeat.pipeline.queue.added.bytes
# TYPE filebeat_libbeat_pipeline_queue_added_bytes_total counter
filebeat_libbeat_pipeline_queue_added_bytes_total 0
# HELP filebeat_libbeat_pipeline_queue_added_events_total libbeat.pipeline.queue.added.events
# TYPE filebeat_libbeat_pipeline_queue_added_events_total counter
filebeat_libbeat_pipeline_queue_added_events_total 389848
# HELP filebeat_libbeat_pipeline_queue_consumed_bytes_total libbeat.pipeline.queue.consumed.bytes
# TYPE filebeat_libbeat_pipeline_queue_consumed_bytes_total counter
filebeat_libbeat_pipeline_queue_consumed_bytes_total 0
# HELP filebeat_libbeat_pipeline_queue_consumed_events_total libbeat.pipeline.queue.consumed.events
# TYPE filebeat_libbeat_pipeline_queue_consumed_events_total counter
filebeat_libbeat_pipeline_queue_consumed_events_total 389848
# HELP filebeat_libbeat_pipeline_queue_filled_bytes libbeat.pipeline.queue.filled.bytes
# TYPE filebeat_libbeat_pipeline_queue_filled_bytes gauge
filebeat_libbeat_pipeline_queue_filled_bytes 0
# HELP filebeat_libbeat_pipeline_queue_filled_events libbeat.pipeline.queue.filled.events
# TYPE filebeat_libbeat_pipeline_queue_filled_events gauge
filebeat_libbeat_pipeline_queue_filled_events 0
# HELP filebeat_libbeat_pipeline_queue_filled_pct libbeat.pipeline.queue.filled.pct
# TYPE filebeat_libbeat_pipeline_queue_filled_pct gauge
filebeat_libbeat_pipeline_queue_filled_pct 0
# HELP filebeat_libbeat_pipeline_queue_max_bytes libbeat.pipeline.queue.max_bytes
# TYPE filebeat_libbeat_pipeline_queue_max_bytes gauge
filebeat_libbeat_pipeline_queue_max_bytes 0
# HELP filebeat_libbeat_pipeline_queue_max_events libbeat.pipeline.queue.max_events
# TYPE filebeat_libbeat_pipeline_queue_max_events gauge
filebeat_libbeat_pipeline_queue_max_events 3200
# HELP filebeat_libbeat_pipeline_queue_removed_bytes_total libbeat.pipeline.queue.removed.bytes
# TYPE filebeat_libbeat_pipeline_queue_removed_bytes_total counter
filebeat_libbeat_pipeline_queue_removed_bytes_total 0
# HELP filebeat_libbeat_pipeline_queue_removed_events_total libbeat.pipeline.queue.removed.events
# TYPE filebeat_libbeat_pipeline_queue_removed_events_total counter
filebeat_libbeat_pipeline_queue_removed_events_total 389848
# HELP filebeat_memstats_gc_next_total beat.memstats.gc_next
# TYPE filebeat_memstats_gc_next_total counter
filebeat_memstats_gc_next_total 4.5624752e+07
# HELP filebeat_memstats_memory beat.memstats.memory_total
# TYPE filebeat_memstats_memory gauge
filebeat_memstats_memory 1.755386064e+09
# HELP filebeat_memstats_memory_alloc beat.memstats.memory_alloc
# TYPE filebeat_memstats_memory_alloc gauge
filebeat_memstats_memory_alloc 3.8430976e+07
# HELP filebeat_memstats_memory_sys beat.memstats.memory_sys
# TYPE filebeat_memstats_memory_sys gauge
filebeat_memstats_memory_sys 6.667188e+07
# HELP filebeat_memstats_rss beat.memstats.rss
# TYPE filebeat_memstats_rss gauge
filebeat_memstats_rss 1.35458816e+08
# HELP filebeat_registrar_states registrar.states
# TYPE filebeat_registrar_states gauge
filebeat_registrar_states{state="cleanup"} 0
filebeat_registrar_states{state="current"} 11
filebeat_registrar_states{state="update"} 407119
# HELP filebeat_registrar_writes registrar.writes
# TYPE filebeat_registrar_writes gauge
filebeat_registrar_writes{writes="fail"} 0
filebeat_registrar_writes{writes="success"} 11
filebeat_registrar_writes{writes="total"} 11
# HELP filebeat_runtime_goroutines beat.runtime.goroutines
# TYPE filebeat_runtime_goroutines gauge
filebeat_runtime_goroutines 52
# HELP filebeat_up Target up
# TYPE filebeat_up gauge
filebeat_up 1
# HELP filebeat_uptime_seconds_total beat.info.uptime.ms
# TYPE filebeat_uptime_seconds_total counter
filebeat_uptime_seconds_total 30.256
```