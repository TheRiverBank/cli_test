package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "Name",
	Version: "0.1.0",
	Short:   "How engineers learn about CLIs",
}

func main() {
	rootCmd.AddCommand(getCmd)

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
