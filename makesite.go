package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type content struct {
	Description string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(fileContents)

}

func writeFile(name string, data string) {
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func renderTemplate(filename string, data string) {
	c := content{Description: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}

func writeTemplateToFile(filename string, data string) {
	c := content{Description: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	f, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}

}

func main() {
	renderTemplate("template.tmpl", readFile("first-post.txt"))
	writeTemplateToFile("template.tmpl", readFile("first-post.txt"))
}
