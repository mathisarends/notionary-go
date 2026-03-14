package main

import (
	"fmt"

	"github.com/mathisbot/notionary-go/blocks/markdown"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:   "markdown [lines...]",
	Short: "Markdown-Zeilen parsen",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parser := markdown.NewLineParser()
		for _, line := range args {
			block, ok := parser.Parse(line)
			if ok {
				fmt.Printf("%-30s → %T\n", line, block)
			} else {
				fmt.Printf("%-30s → kein match\n", line)
			}
		}
	},
}