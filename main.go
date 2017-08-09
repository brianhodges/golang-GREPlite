package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var search = ""
var cnt = 0

func check(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

func examine(path string, f os.FileInfo, err error) error {
	file, err := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), search) {
			fmt.Println(path)
			cnt += 1
		}
	}
	return nil
}

func main() {
	fmt.Println()
	if len(os.Args) <= 2 {
		fmt.Println("You must enter both paramaters.")
		fmt.Println("Ex. go run main.go C:\\path\\to\\directory \"pattern match\"")
		os.Exit(1)
	}

	dir := os.Args[1]
	search = os.Args[2]
	fmt.Println("Results:\n========")
	err := filepath.Walk(dir, examine)
	check(err)
	if cnt == 0 {
		fmt.Println("No matches found.")
	}
}
