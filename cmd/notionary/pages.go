package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var pagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "Page-Operationen",
}

var pagesFindCmd = &cobra.Command{
	Use:   "find [title]",
	Short: "Seite nach Titel suchen",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		page, err := client.Pages.FindByTitle(context.Background(), args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", page)
		return nil
	},
}

func init() {
	pagesCmd.AddCommand(pagesFindCmd)
}