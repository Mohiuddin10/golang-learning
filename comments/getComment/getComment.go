package getComment

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserData(prompt string) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	text = strings.TrimSuffix(text, "?")
	return text, nil
}
