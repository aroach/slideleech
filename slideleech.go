package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
)

var outputDir string

func init() {
	flag.StringVar(&outputDir, "o", "./output", "output directory")
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	flag.Parse();
	
	fmt.Println(outputDir)

	file, err := os.Open("./mocks/test.md")

	check(err)
	
	scanner := bufio.NewScanner(file)

	var matching = false
	var slideFile *os.File
	
	for slideNum := 1; scanner.Scan();  {

		value := scanner.Text()
		
		if value == "-startpreso-" {

			matching = true
			slideFileName := fmt.Sprintf("./mocks/slide%d.md", slideNum)
			fmt.Println(slideFileName)

			var err error
			slideFile, err = os.Create(slideFileName)
			check(err)
			continue
			
		} else if value == "-endpreso-" {
			matching = false
			
			fmt.Println("save");
			// Create new object?
			slideFile.Close()
			slideNum++
		}

		if matching {

			_, err := slideFile.Write([]byte(scanner.Text() + "\n"))
			check(err)
			// fmt.Println(result)
			
			// fmt.Println(value)
		}
		
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("Save created objects here?")
}

