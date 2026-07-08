/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"otlp-tester/models"
	"otlp-tester/http"
	"text/template"
	"github.com/spf13/cobra"
	"bytes"
	"io"
	"time"
)

var (
	metricName string
	metricValue int
)

// metricCmd represents the metric command
var metricCmd = &cobra.Command{
	Use:   "metric",
	Short: "Send predefined sample metric to your OTLP endpoint",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		res, err := http.PostRequest(otlpEndpoint + "/v1/metrics", output)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("Couldn't make API call. %w", err)
		}

		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		fmt.Println(res.StatusCode)
		fmt.Println(string(body))
		fmt.Println("Metric Name in Grafana: " + metricName + "_ratio")
		return nil
	},
	SilenceErrors: true,
	SilenceUsage: true,
}

func init() {
	sendCmd.AddCommand(metricCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metricCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	metricCmd.Flags().StringVar(&metricName, "name", "otlp.test.metric", "Metric name")
	metricCmd.Flags().IntVar(&metricValue, "value", 1, "Metric value")
}
