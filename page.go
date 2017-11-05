package qwertycore

import (
	"io/ioutil"
)

//
type Page struct {
	Filename string
	Title    string
	Body     []byte
}

// LoadPage
func LoadPage(title string) (*Page, error) {
	filename := a + "/views/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Filename: filename, Title: title, Body: body}, nil
}
