package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/MrChiz/ppmap/pkg/cmd"

	"github.com/MrChiz/ppmap/pkg/runner"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ppmap",
		Short: "A tool 4 finding pp(prototype pollution) bug.",
	}
	rootCmd.AddCommand(cmd.ScanCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	//run tool
	runner.Pmap()
}
