package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type GFNPC struct {
	ID                    int      `json:"id"`
	Title                 string   `json:"title"`
	IsFullyOptimized      bool     `json:"isFullyOptimized"`
	IsHighlightsSupported bool     `json:"isHighlightsSupported"`
	SteamURL              string   `json:"steamUrl"`
	Publisher             string   `json:"publisher"`
	Genres                []string `json:"genres"`
	Status                string   `json:"status"`
}

func LoadJSON(r io.Reader) []GFNPC {
	var g []GFNPC
	dec := json.NewDecoder(r)
	for {
		if err := dec.Decode(&g); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return g
}

func LoadUrl(url string) []GFNPC {
	resp, _ := http.Get(url)
	return LoadJSON(resp.Body)
}

func LoadFile(filen string) []GFNPC {
	reader, _ := os.Open(filen)
	return LoadJSON(reader)
}
