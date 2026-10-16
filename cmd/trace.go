package cmd

import (
	"bytes"
	"fmt"
	"io"
	"otlp-tester/http"
	"otlp-tester/models"
	"strconv"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Send predefined sample trace to your OTLP endpoint",
	Long:  `If you are sending sample traces to Grafana Cloud, you must use auth token and put it in GRAFANA_CLOUD_TOKEN environment variable.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		t := time.Now().UnixNano()
		uidString := strconv.FormatInt(t, 10) + strconv.FormatInt(t, 10)
		fmt.Println(uidString)
		uid := uidString[0:32]
		trace := models.Trace{t, uid}

		tmpl, err := template.ParseFS(virtualFS, "templates/trace.tpl")
		if err != nil {
			return fmt.Errorf("Couldn't open template. %w", err)
		}

		var output bytes.Buffer
		tmpl.Execute(&output, trace)
		fmt.Println(output.String())

		res, err := http.PostRequest(otlpEndpoint+"/v1/traces", output, auth)
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
	sendCmd.AddCommand(traceCmd)
}
