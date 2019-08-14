package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var godaddyKey string
var godaddySecret string

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
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Display version.")
}

// initConfig loads settings from environment variables.
func initConfig() {
	godaddyKey = viper.GetString("GODADDY_KEY")
	godaddySecret = viper.GetString("GODADDY_SECRET")
}
