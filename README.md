# Introduction

Slideleech will extract sections of a Markdown file into new files that can be
added to a Revealjs slideshow.

# Features

* Creates a full RevealJS presentation in your output directory.  Currently, this uses an embedded template, which limits the customization that you can perform on the presentation.  By limits, you actually can't customize the template without altering the code.
* Delimit slides in you Markdown document with the opening `[item]: # (slide)` and closing `[item]: # (/slide)` tags.

# Usage

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
```

Author: asroach@cisco.com
