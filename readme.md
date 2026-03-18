# beat-exporter for Prometheus

> **Fork notice:** This is a fork of [trustpilot/beat-exporter](https://github.com/trustpilot/beat-exporter), modernized and updated for filebeat 8.17.4. This project is maintained on a best-effort basis for my own filebeat monitoring needs. It may not receive regular updates beyond that scope. If you're using this fork, feel free to open issues — I'll do my best to help.

## What it does

Scrapes the filebeat HTTP stats endpoint and exposes metrics in Prometheus format.

## What changed from upstream

- Full metric coverage for the filebeat 8.17.4 `/stats` JSON payload (cgroup, handles, latency histograms, queue depth, etc.)
- Go 1.26, updated dependencies (prometheus/client_golang, logrus)
- Multi-stage Dockerfile, non-root runtime
- arm64 build targets (linux, darwin, windows)
- Modernized CI (GitHub Actions v4/v5, ghcr.io releases)
- Removed deprecated APIs (`io/ioutil`, `// +build` tags, `GO111MODULE`)

## Supported beats

- filebeat (primary focus)
- metricbeat
- auditbeat (partial)

## Quick start

### Docker Compose

```bash
docker compose up --build
```

This starts filebeat, the exporter, and Prometheus. Metrics at `http://localhost:9479/metrics`, Prometheus UI at `http://localhost:9090`.

### Standalone

Enable the HTTP endpoint in your filebeat config:

```yaml
http:
  enabled: true
  host: localhost
  port: 5066
```

Run the exporter:

```bash
./beat-exporter
```

Metrics available at `http://localhost:9479/metrics`.

## Configuration

```
./beat-exporter -help
  -beat.uri string       HTTP API address of beat (default "http://localhost:5066")
  -beat.timeout duration Timeout for beat stats (default 10s)
  -beat.system           Expose system stats
  -web.listen-address    Listen address (default ":9479")
  -web.telemetry-path    Metrics path (default "/metrics")
  -tls.certfile string   TLS cert file
  -tls.keyfile string    TLS key file
  -version               Show version
```

## Docker image

```bash
docker pull ghcr.io/karabiy/beat-exporter:latest
```

## Contributing

Issues and PRs welcome.
