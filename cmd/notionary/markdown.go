package main

import (
	"fmt"

	"github.com/mathisbot/notionary-go/blocks/markdown"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Markdown tools",
}

var markdownSyntaxCmd = &cobra.Command{
	Use:   "syntax",
	Short: "Syntax-Referenz aller unterstützten Block-Typen ausgeben",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(markdown.SyntaxGuide())
	},
}

func init() {
	markdownCmd.AddCommand(markdownSyntaxCmd)
}
