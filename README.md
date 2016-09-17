# Introduction

Slideleech will extract sections of a Markdown file into new files that can be
added to a [reveal.js](https://github.com/hakimel/reveal.js) slideshow.

# Features

* Creates a full RevealJS presentation in your output directory.  Read below for template customization.
* Delimit slides in you Markdown document with the opening `[item]: # (slide)` and closing `[item]: # (/slide)` tags.

# Template customization

Within a reveal.js slideshow, customizations can be performed via the [index.html](https://github.com/hakimel/reveal.js/blob/master/index.html) file.  slideleech will automatically add your generated slides to the correct section using the following template markup.  If you want to provide your own `index.html`, just make sure you replace the pertinent part with the snippet below.

```
<div class="reveal">
  <div class="slides">
    {{range. -}}
    <section data-markdown="{{.Content}}"></section>
    {{end}}
  </div>
</div>
```

# Installation

Currently, the project is hosted on Bitbucket.  As a result, the standard `go get` method won't work.

```
$ git clone ssh://git@bitbucket-eng-sjc1.cisco.com:7999/dll/slideleech.git
$ cd slideleech
$ go install
```

# Usage

From within the directory that you are building your slides:

If you're starting from an existing Markdown file, insert the relevant sections in your Markdown file that are bracketed by the slideleech tags.

If you're starting a new project:

```
$ mkdir awesome-project
$ cd awesome-project
```
Create a markdown file, and include the right content and slideleech tags.

```
$ slideleech -i="myMarkdownFile.md" -t="<path to my custom reveal index>/index.html"

```
See below for defaults.

```
*****************************************
This is the slideleech.  It will extract your slide text/bullets contained in a markdown file.

Enclose your slide text/bullets in `[item]: # (slide)` and `[item]: # (/slide)`.
Any content between those tags will be added to your slide file.
Include as many opening and closing tag pairs as you like in your Markdown.

Usage:
  $ slideleech [options] [inputfile [outputfile]]

  -i string
    	input filename (default "./README.md")
  -o string
    	output directory (default "./output")
  -t string
    	full path to RevealJS template
```

Author: asroach@cisco.com
