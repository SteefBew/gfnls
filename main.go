package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"
)

var (
	source, format string
)

const CDXSearchURL = "http://web.archive.org/cdx/search/cdx?output=json"

const URL = "https://static.nvidiagrid.net/supported-public-game-list/gfnpc.json?JSON"

func main() {
	// Feb 4th 20200204184424
	// Feb 5th 20200205085751
	var cmd string
	flag.StringVar(&cmd, "cmd", "ls", "Command to run, options are 'ls' and 'wbls'. 'wbls' provides a list of valid timestamps on archive.org")
	flag.StringVar(&source, "source", "", "Source for the data, either filename or archive.org timestamp")
	flag.StringVar(&format, "format", "{{.Title}}", "Format the output using the given Go template")
	flag.Parse()
	switch cmd {
	case "ls":
		doLs()
	case "wbls":
		doArchiveLs()
	}
}

func doArchiveLs() {
	var g [][]string

	u, _ := url.Parse(CDXSearchURL)
	q, _ := url.ParseQuery(u.RawQuery)

	q.Add("url", URL)

	u.RawQuery = q.Encode()

	resp, _ := http.Get(u.String())
	dec := json.NewDecoder(resp.Body)
	for {
		if err := dec.Decode(&g); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	first := true
	for _, result := range g {
		if first {
			// Skip first
			first = false
			continue
		}
		fmt.Printf("%+v\n", result[1])
	}
}

func doLs() {
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
