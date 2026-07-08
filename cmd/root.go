/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"embed"
)

var otlpEndpoint string

//go:embed templates/*.tpl
var virtualFS embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "otlp-tester",
	Short: "Test your OTLP endpoint by sending metrics, logs, traces",
	Long: `
Use otlp-tester to test your OTLP endpoint by sending metrics, logs, traces. It will use 
OTLP_EXPORTER_ENDPOINT defined in .env or ENV variable if found.

By default, http://localhost:4318 will be used to send all signals.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		otlpEndpoint = os.Getenv("OTLP_EXPORTER_ENDPOINT")
                if otlpEndpoint == "" {
			fmt.Println("\nMissing OTLP_EXPORTER_ENDPOINT. Using default: http://127.0.0.1:4318")
			otlpEndpoint = "http://127.0.0.1:4318"
                } else {
			fmt.Println("OTLP_EXPORTER_ENDPOINT = ", otlpEndpoint, "\n")
		}
                return nil
        },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

