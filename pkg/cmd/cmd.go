package cmd

import (
	"io"
	"log"

	"github.com/spf13/cobra"
)

// Flags
var (
	Url    string
	List   string
	Output string
	Proxy  string
	Silent bool
	/*
		payload        string
		javascript     string
		javascriptFile string
	*/
)

// Sub cmd
var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run the scan with various options",
	Run: func(cmd *cobra.Command, args []string) {
		if Silent {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			cmd.SetOut(io.Discard)
			cmd.SetErr(io.Discard)
			log.SetOutput(io.Discard)
		}
	},
}

// Flag Commend
func init() {
	ScanCmd.Flags().StringVarP(&Url, "url", "u", "", "Input URL")
	ScanCmd.Flags().StringVarP(&List, "list", "l", "", "File containing input URLs")
	ScanCmd.Flags().StringVarP(&Output, "output", "o", "", "File to write output results")
	ScanCmd.Flags().StringVarP(&Proxy, "proxy", "p", "", "Set a proxy server (URL)")
	ScanCmd.Flags().BoolVarP(&Silent, "silent", "s", false, "Silent output. Print only results")
	/*
		scanCmd.Flags().StringVarP(&payload, "payload", "p", "", "Custom payload")
			scanCmd.Flags().StringVarP(&javascript, "javascript", "js", "", "Run custom Javascript on target")
			scanCmd.Flags().StringVarP(&javascriptFile, "javascript-file", "jsf", "", "File containing custom Javascript to run on target")
	*/
}
