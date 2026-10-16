# otlp-tester

CLI tool to test OTLP endpoints. If you are using OpenTelemetry Collectors, Mimir, Prometheus, Loki or Tempo and need to verify your data pipeline, this tool will help you by sending sample signals so you can search them in Grafana on using HTTP APIs.

## Auth
It is possible to use the tool with Grafana Cloud. In Grafana Cloud you need to get the otlp collector token and use base64 enccoded string:

```bash
Authorization: Basic <token>
```

If there is env variable GRAFANA_CLOUD_TOKEN with your base64 encoded auth string, tool will add the Authorization header into each HTTP request. You should also define OTLP_EXPORTER_ENDPOINT according to your Grafana cloud stack.

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


## Notes

Tool is using OTLP/HTTP protocol to send sample telemetry signals.
Some backends like Mimir or Prometheus may rename your metrics by adding _ratio suffix to the name you give.

