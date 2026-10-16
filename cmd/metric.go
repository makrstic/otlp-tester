package cmd

import (
	"bytes"
	"fmt"
	"io"
	"otlp-tester/http"
	"otlp-tester/models"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

var (
	metricName  string
	metricValue int
)

// metricCmd represents the metric command
var metricCmd = &cobra.Command{
	Use:   "metric",
	Short: "Send a sample metric to your OTLP endpoint",
	Long:  `If you are sending sample metrics to Grafana Cloud, you must use auth token and put it in GRAFANA_CLOUD_TOKEN environment variable.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := time.Now().UnixNano()
		m := models.Metric{metricName, metricValue, t}

		tmpl, err := template.ParseFS(virtualFS, "templates/metric.tpl")
		if err != nil {
			return fmt.Errorf("Couldn't open metric template. %w", err)
		}

		var output bytes.Buffer
		tmpl.Execute(&output, m)
		fmt.Println(output.String())

		res, err := http.PostRequest(otlpEndpoint+"/v1/metrics", output, auth)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("Couldn't make API call. %w", err)
		}

		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		fmt.Println(res.StatusCode)
		fmt.Println(string(body))
		fmt.Println("Metric Name in Grafana: " + metricName)
		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	sendCmd.AddCommand(metricCmd)

	metricCmd.Flags().StringVar(&metricName, "name", "otlp.test.metric", "Metric name")
	metricCmd.Flags().IntVar(&metricValue, "value", 1, "Metric value")
}
