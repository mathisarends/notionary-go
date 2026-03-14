package main

import (
	"fmt"

	"github.com/mathisbot/notionary-go/blocks/markdown"
	"github.com/spf13/cobra"
)

var markdownSyntaxCmd = &cobra.Command{
	Use:   "syntax",
	Short: "Syntax-Referenz aller unterstützten Block-Typen ausgeben",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(markdown.SyntaxGuide())
	},
}
