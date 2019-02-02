package main

import (
	"os"
	"log"
	"html/template"
)

/*
Using HTML Templates from a Folder (Complied)
解析html模板
*/
type Page struct {
	Content string
}

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	p := Page{Content: "content"}
	err := templates.ExecuteTemplate(os.Stdout, "index.html", p)
	if err != nil {
		log.Fatal("Cannot Get View ", err)
	}
}
