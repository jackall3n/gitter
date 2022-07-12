/*
Copyright © 2022 Jack Allen <me@jackallen.me>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
)

func runCheckout(cmd *cobra.Command, args []string) error {
	prefix := viper.GetString("prefix")

	var segments []string

	if prefix != "" {
		segments = append(segments, prefix)
	}

	ticket := getTicket(args)

	segments = append(segments, ticket)

	description := getDescription(args)

	if description == "" {
		description = getJiraDescription(ticket)
	}

	if description != "" {
		segments = append(segments, description)
	}

	branch := strings.Join(segments[:], "/")

	coCmd := exec.Command("git", "checkout", "-b", branch)

	out, err := coCmd.CombinedOutput()

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

var checkoutCmd = &cobra.Command{
	Use:     "checkout [ticket] [description]",
	Aliases: []string{"co"},
	Short:   "Checkout a branch using a ticket number and optional description",
	RunE:    runCheckout,
}

func init() {

	rootCmd.AddCommand(checkoutCmd)

	checkoutCmd.PersistentFlags().String("prefix", "", "An optional prefix")
	viper.BindPFlag("prefix", checkoutCmd.PersistentFlags().Lookup("prefix"))

	checkoutCmd.PersistentFlags().String("project", "", "An optional project")
	viper.BindPFlag("project", checkoutCmd.PersistentFlags().Lookup("project"))
}
