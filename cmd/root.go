package cmd

import (
	"fmt"
	"os"

	"github.com/7wmr/godaddy-cli/conf"
	"github.com/spf13/cobra"
)

var config conf.Config

var debugFlag bool
var versionFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godaddy",
	Short: "A command line tool for administering GoDaddy account.",
	Long: `
	This command line tool should be used for 
	the administration of GoDaddy accounts. 
	`,
}

// GetRootCmd to return rootCmd for docs
func GetRootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "D", false, "Enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "V", false, "Display version.")
}

func initConfig() {
	if err := config.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
