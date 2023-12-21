package wikipage

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	f := p.Title + ".page"
	return os.WriteFile(f, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	f := title + ".page"
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: data}, nil
}
