package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	s := "I felt so good!"

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}
}
