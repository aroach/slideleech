package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
)

var outputDir string
var inputFile string

func init() {
	flag.StringVar(&outputDir, "o", "./output", "output directory")
	flag.StringVar(&inputFile, "i", "./README.md", "input filename")
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	flag.Parse();
	
	fmt.Println("Output Directory:", outputDir)
	fmt.Println("Input Filename:", inputFile)

	file, err := os.Open(inputFile)

	check(err)
	
	scanner := bufio.NewScanner(file)

	var matching = false
	var slideFile *os.File

	fmt.Println("Creating slides...")
	
	for slideNum := 1; scanner.Scan();  {

		value := scanner.Text()
		
		if value == "-startpreso-" {

			matching = true
			slideFileName := fmt.Sprintf(outputDir + "/slide%d.md", slideNum)
			fmt.Println(slideFileName)

			var err error
			slideFile, err = os.Create(slideFileName)
			check(err)
			continue
			
		} else if value == "-endpreso-" {
			matching = false
			
			slideFile.Close()
			slideNum++
		}

		if matching {

			_, err := slideFile.Write([]byte(scanner.Text() + "\n"))
			check(err)

		}
		
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}

