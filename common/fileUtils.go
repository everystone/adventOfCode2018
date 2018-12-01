package common

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
