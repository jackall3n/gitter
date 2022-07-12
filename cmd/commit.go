/*
Copyright © 2022 Jack Allen <me@jackallen.me>

*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func getMessage(args []string) string {
	var message string

	if len(args) > 0 {
		message = args[0]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("message: ")

		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		message = text
	}

	return message
}

func getBranch() (string, error) {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	if err != nil {
		return "", errors.New("unable to get branch")
	}

	if branch != nil {
		return strings.Replace(string(branch), "\n", "", -1), nil
	}

	return "", errors.New("unable to get branch")
}

func getTicketFromBranch(branch string) string {
	r := regexp.MustCompile(`^([A-Za-z]+/)?(?P<ticket>[A-Za-z]{2,5}-\d+)`)

	matches := r.FindStringSubmatch(branch)

	for i, name := range r.SubexpNames() {
		if name == "ticket" {
			if len(matches) < i {
				return ""
			}

			return matches[i]
		}
	}

	return ""
}

func runCommit(cmd *cobra.Command, args []string) {
	message := getMessage(args)

	branch, err := getBranch()

	if err != nil {
		fmt.Println(err)
		return
	}

	ticket := getTicketFromBranch(branch)

	var commit string

	if ticket == "" {
		commit = message
	} else {
		commit = fmt.Sprintf("[%s]: %s", ticket, message)
	}

	coCmd := exec.Command("git", "commit", "-m", commit)

	out, err := coCmd.CombinedOutput()

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit [message]",
	Short: "Just commit",
	Run:   runCommit,
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
