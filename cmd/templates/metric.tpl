{
  "resource_metrics": [
    {
      "resource": {
        "attributes": [
          { "key": "service.name", "value": { "string_value": "otlp-tester" } }
        ]
      },
      "scope_metrics": [
        {
          "metrics": [
            {
              "name": "{{ .MetricName }}",
              "description": "A simple test metric",
              "unit": "1",
              "gauge": {
                "dataPoints": [
                  {
			"asDouble": {{ .MetricValue }},
                    	"timeUnixNano": "{{ .CurrentTime }}",
			"attributes": [
			{
				"key": "host.name",
				"value": {
					"string_value": "otlp.test.host"
				}
			},
			{
				"key": "service.name",
				"value": {
					"string_value": "otlp-tester"
				}
			}
			]
                  }
                ]
              }
            }
          ]
        }
      ]
    }
  ]
}
