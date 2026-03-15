package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ask(question string) string {
	fmt.Print(question + " ")
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	return strings.TrimSpace(answer)
}

func Select(question string, options []string) string {
	fmt.Println(question)
	for i, opt := range options {
		fmt.Printf("  %d) %s\n", i+1, opt)
	}
	fmt.Print("Enter choice: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for i, opt := range options {
		if input == fmt.Sprintf("%d", i+1) {
			return opt
		}
	}

	fmt.Println("invalid choice, please try again")
	return Select(question, options)
}

func Confirm(question string) bool {
	answer := Ask(question + " (y/n):")
	return strings.ToLower(answer) == "y"
}
