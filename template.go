package main

const INDEX_TEMPLATE = `
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">

		<title>reveal.js</title>

		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/css/reveal.css">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/css/theme/black.css">

		<!-- Theme used for syntax highlighting of code -->
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/lib/css/zenburn.css">

		<!-- Printing and PDF exports -->
		<script>
			var link = document.createElement( 'link' );
			link.rel = 'stylesheet';
			link.type = 'text/css';
			link.href = window.location.search.match( /print-pdf/gi ) ? 'https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/css/print/pdf.css' : 'https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/css/print/paper.css';
			document.getElementsByTagName( 'head' )[0].appendChild( link );
		</script>
	</head>
	<body>
		<div class="reveal">
			<div class="slides">
				{{range. -}}
				<section data-markdown="{{.Content}}" data-background-color="#{{.Color}}"></section>
				{{end}}
			</div>
		</div>

		<script src="https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/lib/js/head.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/js/reveal.js"></script>

		<script>
			// More info https://github.com/hakimel/reveal.js#configuration
			Reveal.initialize({
			history: true,
			width: 960,
			height: 700,

				// More info https://github.com/hakimel/reveal.js#dependencies
				dependencies: [
					{ src: 'https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/plugin/markdown/marked.js' },
					{ src: 'https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/plugin/markdown/markdown.js' },
					{ src: 'https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/plugin/notes/notes.js', async: true },
					{ src: 'https://cdnjs.cloudflare.com/ajax/libs/reveal.js/3.3.0/plugin/highlight/highlight.js', async: true, callback: function() { hljs.initHighlightingOnLoad(); } }
				]
			});
		</script>
	</body>
</html>


`

const INTRO_SLIDE = `
## {{.Title}}

{{.Author}}

`

const CLOSING_SLIDE = `
## {{.Message}}

`

const LEECH_FILE = `
---
input_file: ./README.md
output_directory: ./slides
output_filename: slide
reveal:
  template:
  template_color: FFFFFF
  intro:
  intro_title: Insert Your Title Here!
  intro_author: Author / Title / twitter
  intro_color: 049FD9
  closing:
  closing_message: Thank you!
  closing_color: 049FD9
`
