package cmd

import (
	"github.com/spf13/cobra"
)

var addDomain string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a record for a domain.",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		// TBD
	},
}

func init() {
	domainCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&addDomain, "domain", "d", "", "Instance domain")
	setCmd.MarkFlagRequired("domain")
}
