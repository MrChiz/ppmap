package cmd

import (
	"github.com/spf13/cobra"
)

// Flags
var (
	Url string
	/*
		list           string
		proxy          string
		payload        string
		javascript     string
		javascriptFile string
		output         string
		silent         bool
	*/
)

// Sub cmd
var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run the scan with various options",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Flag Commend
func init() {
	ScanCmd.Flags().StringVarP(&Url, "url", "u", "", "Input URL")
	/*scanCmd.Flags().StringVarP(&list, "list", "l", "", "File containing input URLs")
	scanCmd.Flags().StringVarP(&proxy, "proxy", "px", "", "Set a proxy server (URL)")
	scanCmd.Flags().StringVarP(&payload, "payload", "p", "", "Custom payload")
	scanCmd.Flags().StringVarP(&javascript, "javascript", "js", "", "Run custom Javascript on target")
	scanCmd.Flags().StringVarP(&javascriptFile, "javascript-file", "jsf", "", "File containing custom Javascript to run on target")
	scanCmd.Flags().StringVarP(&output, "output", "o", "", "File to write output results")
	scanCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Silent output. Print only results")*/
}
