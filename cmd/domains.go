package cmd

import (
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Commands to administrer GoDaddy domains.",
	Long:  `TBC`,
}

func init() {
	rootCmd.AddCommand(domainCmd)
}
