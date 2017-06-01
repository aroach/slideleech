# Introduction

Slideleech will extract sections of a Markdown file into new files that can be
added to a [reveal.js](https://github.com/hakimel/reveal.js) slideshow.

# Features

* Creates a full RevealJS presentation in your output directory.  Read below for template customization.
* Delimit slides in you Markdown document with the opening `[item]: # (slide)` and closing `[item]: # (/slide)` tags.

Review the [CHANGELOG](CHANGELOG.md)

# Template customization

Within a reveal.js slideshow, customizations can be performed via the [index.html](https://github.com/hakimel/reveal.js/blob/master/index.html) file.  slideleech will automatically add your generated slides to the correct section using the following template markup.  If you want to provide your own `index.html`, just make sure you replace the pertinent part with the snippet below.

```
<div class="reveal">
  <div class="slides">
    {{range. -}}
    <section data-markdown="{{.Content}}"  data-background-color="#{{.Color}}"></section>
    {{end}}
  </div>
</div>
```

# Installation

Currently, the project is hosted on Bitbucket.  As a result, the standard `go get` method won't work.

```
$ git clone ssh://git@bitbucket-eng-sjc1.cisco.com:7999/dll/slideleech.git
$ cd slideleech
$ go get gopkg.in/yaml.v2
$ go install
```

# Configuration

Create a `.leech.yml` file:

```
---
input_file: ./README.md
output_directory: ./slides
output_filename: slide
reveal:
  template:
  template_color: FFFFFF
  intro:
  intro_title: Your Fancy Title
  intro_author: John Doe / ACME, Inc. / @jdoe
  intro_color: 049FD9
  closing:
  closing_message: Thank you!
  closing_color: 049FD9
```

| Key           | Value       | Required  |
| ------------- | ----------- | --------- |
| input_file    | The markdown file that you would like to generate slides from. | Yes |
| output_directory      | The directory name where to place the slides.      |   Yes |
| output_filename | The string to use as the filename. (E.g., by entering `slide` as a value would produce slide.html, slide0.html, etc.)     |    Yes |
| reveal > template | An existing reveal template to use for slide generation.      |    No |
| reveal > template_color | Background color for main slides.     |    Yes |
| reveal > intro | Not used      |    No |
| reveal > intro_title | The presentation title |    Yes |
| reveal > intro_author | The presentation author |    Yes |
| reveal > intro_color | Background color for intro slide |    Yes |
| reveal > closing | Not used      |    No |
| reveal > closing_message | Closing slide message     |    Yes |
| reveal > closing_color | Closing slide background color      |    Yes |


# Usage

From within the directory that you are building your slides:

If you're starting from an existing Markdown file, insert the relevant sections in your Markdown file that are bracketed by the slideleech tags.

If you're starting a new project:

```
$ mkdir awesome-project
$ cd awesome-project
```

Create your Markdown file with the `[item]: # (slide)` and `[item]: # (/slide)`.  See [example](mocks/test.md) for a brief example.

Start leeching!

```
$ slideleech

```

To serve the slides on your local machine use the `-serve` flag.  This will build your project and serve up on port 3000.

```
$ slideleech -serve
2017/06/01 17:19:36 Listening...
2017/06/01 17:19:36 Open your browser to http://localhost:3000
```


Author: asroach@cisco.com
