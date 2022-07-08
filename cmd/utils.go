package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
Get Config path for current directory
*/
func getConfigPath() (string, error) {
	if cfgFile != "" {
		return cfgFile, nil
	}

	localPath, err := os.Getwd()

	if err != nil {
		return "", err
	}

	// Locally
	if _, err := os.Stat(fmt.Sprintf("%s/.gitter.yaml", localPath)); err == nil {
		return localPath, nil
	}

	homePath, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	// Globally
	if _, err := os.Stat(fmt.Sprintf("%s/.gitter.yaml", homePath)); err == nil {
		return homePath, nil
	}

	return "", errors.New("unable to find config file")
}

/**
Gets a ticket from either an argument or user input
*/
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

/**
Get description from either an argument or user input
*/
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
