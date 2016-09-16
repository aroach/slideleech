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
	flag.StringVar(&inputFile, "i", "./README.md", "input filename")
	flag.StringVar(&outputDir, "o", "./output", "output directory")
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "\n*****************************************\n"+
      "This is the slideleech.  It will extract "+
      "your slide text/bullets contained in a markdown file.\n\n"+
      "Enclose your slide text/bullets in "+
      "`[item]: # (slide)` and `[item]: # (/slide)`.\n"+
      "Any content between those tags will be added to your slide file.\n"+
      "Include as many opening and closing tag pairs as you like "+
      "in your Markdown.\n\n"+
      "Usage:\n"+
      "  %s [options] [inputfile [outputfile]]\n\n",
      os.Args[0])
    flag.PrintDefaults()
  }
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	flag.Parse();
  // fmt.Println(inputFile)
  // if len(inputFile) == 0 {
  //   flag.Usage()
  //   os.Exit(-1)
  // }

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

		if value == "[item]: # (slide)" {

			matching = true
			slideFileName := fmt.Sprintf(outputDir + "/slide%d.md", slideNum)
			fmt.Println(slideFileName)

			var err error
			slideFile, err = os.Create(slideFileName)
			check(err)
			continue

		} else if value == "[item]: # (/slide)" {
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
