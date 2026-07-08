{
  "resourceSpans": [
    {
      "resource": {
        "attributes": [
          {
            "key": "service.name",
            "value": {
              "stringValue": "otlp-tester"
            }
          }
        ]
      },
      "scopeSpans": [
        {
          "scope": {
            "name": "my.library",
            "version": "1.0.0",
            "attributes": [
              {
                "key": "otlp.service.name",
                "value": {
                  "stringValue": "otlp-tester"
                }
              }
            ]
          },
          "spans": [
            {
              "traceId": "{{ .TraceId }}",
              "spanId": "EEE19B7EC3C1B174",
              "name": "root-span",
              "startTimeUnixNano": "{{ .CurrentTime }}",
              "endTimeUnixNano": "{{ .CurrentTime }}",
              "kind": 2,
              "attributes": [
                {
                  "key": "otlp.service.name",
                  "value": {
                    "stringValue": "otlp-tester"
                  }
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
