# ADR-020 — Introduce Prometheus Observability Layer

## Status
Accepted

## Context
The system now includes validation, use cases, adapters and Clean Architecture.
However, it lacks observability:

- no latency measurement,
- no counting of API calls,
- no error ratio tracking,
- no endpoint-level visibility.

Prometheus is the industry standard for Go services, containers, and microservices.

Prometheus works through:
- counters,
- histograms,
- gauges,
- summaries.

Gin supports middleware for Prometheus using:
- custom middleware
- `promhttp.Handler()`

## Decision
Introduce a dedicated observability layer.

Components:
- middleware recording:
  - request count,
  - request latency histogram,
  - error count.
- `/metrics` endpoint for scraping.

Metrics added:
- `http_requests_total`
- `http_request_duration_seconds`
- `http_error_total`

## Consequences

### Positive
- Production-grade monitoring.
- Compatible with Grafana dashboards.
- Detects regressions, latency spikes, and load.

### Negative
- Slight performance overhead (minimal).
- Requires scraping infrastructure.

## Alternatives
A) OpenTelemetry  
Rejected for now — too heavy for this stage.

B) No observability  
Rejected — unacceptable for production-targeted architecture.
