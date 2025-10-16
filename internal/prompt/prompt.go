package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Confirm prompts the user with a yes/no question and returns true if they answer yes
func Confirm(message string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s (y/n): ", message)

	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

// ConfirmCommand asks the user if they want to run a suggested command
func ConfirmCommand(projectType, commandName, command string) bool {
	fmt.Printf("\nNo mapping found for '%s' in this project.\n", commandName)
	fmt.Printf("Detected: %s project\n\n", projectType)

	return Confirm(fmt.Sprintf("Run \"%s\"?", command))
}
