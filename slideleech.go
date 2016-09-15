package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	file, err := os.Open("./mocks/test.md")

	if err != nil {
		log.Fatal(err)
	}
	
	scanner := bufio.NewScanner(file)

	var matching = false
	
	for scanner.Scan() {

		value := scanner.Text()
		
		if value == "-startpreso-" {
			matching = true
		} else if value == "-endpreso-" {
			matching = false
		}

		if matching && value != "-startpreso-" {
			
			fmt.Println(value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

