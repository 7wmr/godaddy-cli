package cmd

import (
	"github.com/spf13/cobra/doc"
)

func init() {
	doc.GenReSTTree(cmd.rootCmd, "./docs")
}
