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
$ go install
```

# Configuration

Create a `.leech.yml` file:

```
---
input_file: ./README.md
output_directory: ./slides
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


Author: asroach@cisco.com
