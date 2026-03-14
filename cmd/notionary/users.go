package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func nameStr(s *string) string {
	if s == nil {
		return "<kein Name>"
	}
	return *s
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "User-Operationen",
}

var usersMeCmd = &cobra.Command{
	Use:   "me",
	Short: "Aktuellen Bot-User anzeigen",
	RunE: func(cmd *cobra.Command, args []string) error {
		me, err := client.Users.Me(context.Background())
		if err != nil {
			return err
		}
		fmt.Printf("ID:   %s\nName: %s\n", me.ID, nameStr(me.Name))
		return nil
	},
}

var usersListCmd = &cobra.Command{
	Use:   "list",
	Short: "Alle User im Workspace auflisten",
	RunE: func(cmd *cobra.Command, args []string) error {
		users, err := client.Users.List(context.Background())
		if err != nil {
			return err
		}
		for _, u := range users {
			if u.Person != nil {
				fmt.Printf("[person] %s – %s\n", u.Person.ID, nameStr(u.Person.Name))
			} else if u.Bot != nil {
				fmt.Printf("[bot]    %s – %s\n", u.Bot.ID, nameStr(u.Bot.Name))
			}
		}
		return nil
	},
}

func init() {
	usersCmd.AddCommand(usersMeCmd, usersListCmd)
}