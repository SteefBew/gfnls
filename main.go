package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

const URL = "https://static.nvidiagrid.net/supported-public-game-list/gfnpc.json?JSON"

func main() {
	// Feb 4th 20200204184424
	// Feb 5th 20200205085751
	var source, format string
	flag.StringVar(&source, "source", "", "Source for the data, either filename or archive.org timestamp")
	flag.StringVar(&format, "format", "{{.Title}}", "Format the output using the given Go template")
	flag.Parse()
	g := GFNPC{}
	if len(source) != 0 {
		if _, err := os.Stat(source); os.IsNotExist(err) {
			// Not a valid file so try reading it as an archive.org timestamp
			g.LoadUrl(fmt.Sprintf("https://web.archive.org/web/%sid_/%s", source, URL))
		} else {
			// Valid path so just read from it
			g.LoadFile(source)
		}
	} else {
		// No source provided so default to the main nvidia site
		g.LoadUrl(URL)
	}

	tFormat := template.Must(template.New("format").Parse(fmt.Sprintf("%s\n", format)))

	for _, title := range g {
		tFormat.Execute(os.Stdout, title)
	}
}
