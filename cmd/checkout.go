/*
Copyright © 2022 Jack Allen <me@jackallen.me>

*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func getTicket(args []string) string {
	var ticket string
	project := viper.GetString("project")

	if len(args) > 0 {
		ticket = args[0]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("ticket: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		ticket = text
	}

	_, err := strconv.Atoi(ticket)

	if err == nil && project != "" {
		ticket = fmt.Sprintf("%s-%s", project, ticket)
	}

	return strings.ToUpper(ticket)

}

func getDescription(args []string) string {
	var description string

	if len(args) > 1 {
		description = args[1]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("description: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		description = text
	}

	sp := regexp.MustCompile(" ")

	description = sp.ReplaceAllString(description, "-")

	lr := regexp.MustCompile("[^A-Za-z\\d-_]")

	description = lr.ReplaceAllString(description, "")

	return strings.ToLower(description)
}

func runCheckout(cmd *cobra.Command, args []string) error {
	prefix := viper.GetString("prefix")

	var segments []string

	if prefix != "" {
		segments = append(segments, prefix)
	}

	ticket := getTicket(args)
	segments = append(segments, ticket)

	description := getDescription(args)

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
