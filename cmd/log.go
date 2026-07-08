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

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Send predefined sample log message to your OTLP endpoint",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		res, err := http.PostRequest(otlpEndpoint + "/v1/logs", output)
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
	SilenceUsage: true,
}

func init() {
	sendCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
