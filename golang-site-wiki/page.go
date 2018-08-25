package main

import (
	"io/ioutil"
	"path/filepath"
)

// Page - A struct that represents and wiki page
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := filepath.Join("data", p.Title+".txt")
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := filepath.Join("data", title+".txt")
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
