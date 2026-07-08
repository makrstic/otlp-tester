# otlp-tester

CLI tool to test OTLP endpoints. If you are using OpenTelemetry Collectors, Mimir, Prometheus, Loki or Tempo and need to verify your data pipeline, this tool will help you by sending sample signals so you can search them in Grafana on using HTTP APIs.

## Use
You need to define your OTLP endpoint by setting environment variable OTLP_EXPORTER_ENDPOINT. The tool will add /v1/metrics, /v1/logs or /v1/traces when POSTing sample signal. Default endpoint is: http://localhost:4318. Default metric name is otlp_test_metric.

```bash
export OTLP_EXPORTER_ENDPOINT=https://ingest.your.domain:4318

otlp-tester --help

# Send custom metrics 
otlp-tester send metric --name my.sample.metric --value 1

# Send default metric
otlp-tester send metric

# Send default log
otlp-tester send log

# Send default trace
otlp-tester send trace
```

All metrics sent by the tool may get *_ratio* suffix set by Mimir/Prometheus.

