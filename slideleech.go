package main

import (
	"errors"
	"fmt"
	"bufio"
	"os"
	"flag"
	"html/template"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var outputDir string
var inputFile string
var templateFile string
var revealTemplColor string
var revealIntro string
var revealIntroAuthor string
var revealIntroTitle string
var revealIntroColor string
var revealClosing string
var revealClosingMsg string
var revealClosingColor string

type SlideEntry struct {
	Content string
	Color string
}

type Config struct {
	InputFile string `yaml:"input_file"`
	OutputDir string `yaml:"output_directory"`
	Reveal struct {
		Templ string `yaml:"template"`
		TemplColor string `yaml:"template_color"`
		Intro string `yaml:"intro"`
		IntroTitle string `yaml:"intro_title"`
		IntroAuthor string `yaml:"intro_author"`
		IntroColor string `yaml:"intro_color"`
		Closing string `yaml:"closing"`
		ClosingMsg string `yaml:"closing_message"`
		ClosingColor string `yaml:"closing_color"`
	}
}

func (c *Config) Parse(data []byte) error {
	err := yaml.Unmarshal(data, c)
	check(err)

	if c.InputFile == "" {
		return errors.New("Slideleech config: invalid `input_file`")
	}
	if c.OutputDir == "" {
		return errors.New("Slideleech config: invalid `output_directory`")
	}

	inputFile = c.InputFile
	outputDir = c.OutputDir
	templateFile = c.Reveal.Templ
	revealTemplColor = c.Reveal.TemplColor
	revealIntro = c.Reveal.Intro
	revealIntroTitle = c.Reveal.IntroTitle
	revealIntroAuthor = c.Reveal.IntroAuthor
	revealIntroColor = c.Reveal.IntroColor
	revealClosing = c.Reveal.Closing
	revealClosingMsg = c.Reveal.ClosingMsg
	revealClosingColor = c.Reveal.ClosingColor

	return nil
}

func init() {
	var config Config

	yml, err := ioutil.ReadFile(".leech.yml")
	check(err)
	err2 := config.Parse(yml)
	check(err2)

	// flag.StringVar(&inputFile, "i", "./README.md", "input filename")
	// flag.StringVar(&outputDir, "o", "./slides", "output directory")
	// flag.StringVar(&templateFile, "t", "", "full path to RevealJS template")
	flag.Usage = func() {
	fmt.Fprintf(os.Stderr, "\n*****************************************\n"+
		"This is the slideleech.  It will extract "+
		"your slide text/bullets contained in a markdown file.\n\n"+
		"Enclose your slide text/bullets in "+
		"`[item]: # (slide)` and `[item]: # (/slide)`.\n"+
		"Any content between those tags will be added to your slide file.\n"+
		"Include as many opening and closing tag pairs as you like "+
		"in your Markdown.\n\n"+
		"Edit a `.leech.yml` to configure your project\n\n",
		os.Args[0])
	flag.PrintDefaults()
  }
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CreateIntroSlide() {
	type IntroSlide struct {
		Title string
		Author string
	}
	slide := IntroSlide{revealIntroTitle, revealIntroAuthor}
	templ, err := template.New("slide0.md").Parse(INTRO_SLIDE)
	introFile, err := os.Create(outputDir + "/slide0.md")
	fmt.Printf("Creating intro slide...\n")
	err = templ.Execute(introFile, slide)
	check(err)
	introFile.Close()
}

func CreateClosingSlide(slideCount int) {
	type ClosingSlide struct {
		Message string
	}
	slideName := fmt.Sprintf("slide%d.md", slideCount)
	slide := ClosingSlide{revealClosingMsg}
	templ, err := template.New("slideEnd.md").Parse(CLOSING_SLIDE)
	introFile, err := os.Create(outputDir + "/" + slideName)
	fmt.Printf("Creating closing slide... %s\n", slideName)
	err = templ.Execute(introFile, slide)
	check(err)
	introFile.Close()
}


func CreateSite(slideCount int) {

    var slides []SlideEntry

		// add intro slide to slice
		var intro SlideEntry
		intro.Content = "slide0.md"
		intro.Color = revealIntroColor
		slides = append(slides, intro)

		// Add content slides
    for i := 1; i <= slideCount; i++ {
      var slideName SlideEntry
      slideName.Content = fmt.Sprintf("slide%d.md", i)
			slideName.Color = revealTemplColor
      slides = append(slides, slideName)
    }

		// add closing slide to slice
		var closing SlideEntry
		closing.Content = fmt.Sprintf("slide%d.md", slideCount + 1)
		closing.Color = revealClosingColor
		slides = append(slides, closing)

    // Include custom template based on a flag
    templ := template.New("index.html")
    if templateFile != "" {
      fmt.Println("Creating RevealJS index.html from EXTERNAL template...")
      templ, _ = templ.ParseFiles(templateFile)
    } else {
      fmt.Println("Creating RevealJS index.html from INTERNAL template...")
      templ, _ = templ.Parse(INDEX_TEMPLATE)
    }
    // Save the template to the output directory
    indexFile, err := os.Create(outputDir + "/index.html")
    err = templ.Execute(indexFile, slides)
    indexFile.Close()
    check(err)

}

func CreateSlides(inputFile string) int {

	file, err := os.Open(inputFile)

	check(err)

	scanner := bufio.NewScanner(file)

	var matching = false
	var slideFile *os.File
	var slideNum int

	fmt.Println("Creating slides...")

	for slideNum = 1; scanner.Scan();  {

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
	return slideNum
}

func main() {
	flag.Parse();

	fmt.Println("Output Directory:", outputDir)
	fmt.Println("Input Filename:", inputFile)
	fmt.Println("Template File:", templateFile)

	// Make a directory for the slideshow
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}

	CreateIntroSlide()
	slideNum := CreateSlides(inputFile)
	CreateSite(slideNum - 1)
	CreateClosingSlide(slideNum)

}
