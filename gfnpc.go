package main

import (
	"encoding/json"
	"io"
	"log"
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

func (g *GFNPC) Load(r io.Reader) {
	dec := json.NewDecoder(r)
	for {
		if err := dec.Decode(&g); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
