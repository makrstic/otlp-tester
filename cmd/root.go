package cmd

import (
	"embed"
	"fmt"
	"os"
	"otlp-tester/models"

	"github.com/spf13/cobra"
)

var otlpEndpoint string
var grafanaCloudToken string
var auth models.Auth

//go:embed templates/*.tpl
var virtualFS embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "otlp-tester",
	Short: "Test your OTLP endpoint by sending metrics, logs, traces",
	Long: `
Use otlp-tester to test your OTLP endpoint by sending metrics, logs, traces. It will use 
OTLP_EXPORTER_ENDPOINT defined in .env or ENV variable if found. If you are using Grafana Cloud, you must use GRAFANA_CLOUD_TOKEN defined in ENV.

By default, http://localhost:4318 will be used to send all signals.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		otlpEndpoint = os.Getenv("OTLP_EXPORTER_ENDPOINT")
		grafanaCloudToken = os.Getenv("GRAFANA_CLOUD_TOKEN")

		if otlpEndpoint == "" {
			fmt.Println("\nMissing OTLP_EXPORTER_ENDPOINT. Using default: http://127.0.0.1:4318")
			otlpEndpoint = "http://127.0.0.1:4318"
		} else {
			fmt.Println("OTLP_EXPORTER_ENDPOINT = ", otlpEndpoint)
		}

		if grafanaCloudToken == "" {
			fmt.Println("\nMissing GRAFANA_CLOUD_TOKEN. Using no auth.")
			auth = models.Auth{AuthType: "none", Token: ""}
		} else {
			fmt.Println("GRAFANA_CLOUD_TOKEN = ", grafanaCloudToken)
			auth = models.Auth{AuthType: "Basic", Token: grafanaCloudToken}
			return nil
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
