package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

// docsCmd represents the domain command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Commands to administer GoDaddy domains.",
	Long:  `TBC`,
	Run:   runDoc,
}

func runDoc(cmd *cobra.Command, args []string) {
	dir := "./docs"
	err := doc.GenReSTTree(rootCmd, dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
