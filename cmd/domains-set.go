package cmd

import (
	"github.com/7wmr/godaddy-cli/conf"

	"github.com/spf13/cobra"
)

var addName string
var addUsername string
var addDomain string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a record for a domain.",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		auth := conf.Auth{Username: addUsername}
		auth.PromptPassword()
		context := conf.Context{Auth: auth.Encode(), Domain: addDomain, Name: addName}
		config.AddContext(context)
		config.WriteConfig(cfgPath)
	},
}

func init() {
	domainCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&addName, "name", "n", "", "Instance name")
	setCmd.MarkFlagRequired("name")

	setCmd.Flags().StringVarP(&addUsername, "username", "u", "", "Instance username")
	setCmd.MarkFlagRequired("username")

	setCmd.Flags().StringVarP(&addDomain, "domain", "d", "", "Instance domain")
	setCmd.MarkFlagRequired("domain")
}
