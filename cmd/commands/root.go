package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yabf-cli",
	Short: "Yet Another Blockchain Framework CLI",
}

func Execute() error {
	return rootCmd.Execute()
}
