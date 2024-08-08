package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadCredFile() map[string]string {
	file, err := os.Open(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string) // File does not exist, return empty map
		}
		return nil
	}
	defer file.Close()

	cred := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			cred[key] = value
		}
	}
	fmt.Println(cred)
	return cred
}
