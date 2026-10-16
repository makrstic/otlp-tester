/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send metric, trace or log to the OTLP endpoint.",
	Long:  `Send metric, trace or log to the OTLP endpoint using pre-defined JSON template.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send called")
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
