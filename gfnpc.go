package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type GFNPC []struct {
	ID                    int      `json:"id"`
	Title                 string   `json:"title"`
	IsFullyOptimized      bool     `json:"isFullyOptimized"`
	IsHighlightsSupported bool     `json:"isHighlightsSupported"`
	SteamURL              string   `json:"steamUrl"`
	Publisher             string   `json:"publisher"`
	Genres                []string `json:"genres"`
	Status                string   `json:"status"`
}

func (g *GFNPC) load(r io.Reader) {
	dec := json.NewDecoder(r)
	for {
		if err := dec.Decode(&g); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *GFNPC) LoadUrl(url string) {
	resp, _ := http.Get(url)
	g.load(resp.Body)
}

func (g *GFNPC) LoadFile(filen string) {
	reader, _ := os.Open(filen)
	g.load(reader)
}
