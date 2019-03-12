package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileCount := 0

	file, err := os.Open("../../../../wiki/content/query-language/index.md")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inQuery := false

	f, err := os.Create(fmt.Sprintf("query-%02d", fileCount))
	check(err)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "runnable") {
			inQuery = !inQuery
			if inQuery {
				f.Close()
				fileCount++
				if fileCount >= 87 {
					// last three queries require GraphQL Variables
					break
				}
				f, err = os.Create(fmt.Sprintf("query-%02d", fileCount))
				check(err)
			}

			continue
		}
		if inQuery {
			f.WriteString(line + "\n")
		}
	}

	f.Close()
	check(scanner.Err())
}
