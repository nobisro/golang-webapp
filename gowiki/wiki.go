package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

// This is a method named save that takes as its receiver p, a pointer to Page. It takes no parameters, and returns a value of type error. It returns an error value because that is the return type of WriteFile. Octal integer 0600 indicates that the file be created with r/w permissions
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Constructs the file name from title paramater, reads file's contents into body variable, and returns pointer to a Page literal, constructed with proper title and body.
func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
