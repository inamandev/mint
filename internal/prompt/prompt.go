package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func Ask(question string) string {
	fmt.Print(question + " ")
	return readLine()
}

func AskDefault(question, defaultVal string) string {
	fmt.Printf("%s [%s] ", question, defaultVal)
	input := readLine()
	if input == "" {
		return defaultVal
	}
	return input
}

func Select(question string, options []string) string {
	return SelectDefault(question, options, "")
}

func SelectDefault(question string, options []string, defaultVal string) string {
	fmt.Println(question)
	defaultIdx := -1
	for i, opt := range options {
		marker := "  "
		if opt == defaultVal {
			marker = "* "
			defaultIdx = i
		}
		fmt.Printf("%s%d) %s\n", marker, i+1, opt)
	}

	if defaultIdx >= 0 {
		fmt.Printf("Enter choice [%d]: ", defaultIdx+1)
	} else {
		fmt.Print("Enter choice: ")
	}

	input := readLine()

	if input == "" && defaultIdx >= 0 {
		return defaultVal
	}

	for i, opt := range options {
		if input == fmt.Sprintf("%d", i+1) {
			return opt
		}
	}

	fmt.Println("invalid choice, please try again")
	return SelectDefault(question, options, defaultVal)
}

func Confirm(question string) bool {
	return ConfirmDefault(question, false)
}

func ConfirmDefault(question string, defaultVal bool) bool {
	hint := "y/N"
	if defaultVal {
		hint = "Y/n"
	}
	fmt.Printf("%s (%s): ", question, hint)
	input := strings.ToLower(readLine())
	if input == "" {
		return defaultVal
	}
	return input == "y"
}
