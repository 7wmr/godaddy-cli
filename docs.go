package main

import (
	"fmt"
	"github.com/7wmr/godaddy-cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenReSTTree(cmd.GetRootCmd(), "./docs")
	if err != nil {
		fmt.Println(err)
	}
}
