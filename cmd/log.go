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

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Send predefined sample log message to your OTLP endpoint",
	Long:  `If you are sending sample logs to Grafana Cloud, you must use auth token and put it in GRAFANA_CLOUD_TOKEN environment variable.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := time.Now().UnixNano()
		log := models.Log{t}

		tmpl, err := template.ParseFS(virtualFS, "templates/log.tpl")
		if err != nil {
			return fmt.Errorf("Couldn't open log template.. %w", err)
		}

		var output bytes.Buffer
		tmpl.Execute(&output, log)
		fmt.Println(output.String())

		res, err := http.PostRequest(otlpEndpoint+"/v1/logs", output, auth)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("Couldn't make API call. %w", err)
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		fmt.Println("HTTP Code: ", res.StatusCode)
		fmt.Println(string(body))
		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	sendCmd.AddCommand(logCmd)
}
