{
  "resourceLogs": [
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
      "scopeLogs": [
        {
          "scope": {
            "name": "my.library",
            "version": "1.0.0",
            "attributes": [
              {
                "key": "otlp.tester.attribute",
                "value": {
                  "stringValue": "some otlp-tester attribute"
                }
              }
            ]
          },
          "logRecords": [
            {
              "timeUnixNano": "{{ .CurrentTime }}",
              "observedTimeUnixNano": "{{ .CurrentTime }}",
              "severityNumber": 10,
              "severityText": "Information",
              "traceId": "5B8EFFF798038103D269B633813FC60C",
              "spanId": "EEE19B7EC3C1B174",
              "body": {
                "stringValue": "Example log record from otlp-tester"
              },
              "attributes": []
            }
          ]
        }
      ]
    }
  ]
}
