package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
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
func filterInput(input string) string {
	ext := ".html"
	s := strings.Split(input, ".")[0] + ext
	return s
	// char := input
	// new_input := ""
	// for i := 0; i <= len(input); i++ {
	// 	if char[i] != "." {
	// 		new_input += char[i]
	// 	}
	// 	else {
	// 		return
	// 	}
	// }
}

func writeTemplateToFile(templateName string, data string) {
	c := content{Description: readFile(data)}
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))

	filter := filterInput(data)
	// f, err := os.Create(arg[:len(arg)-4] + ".html")
	f, err := os.Create(filter)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}

}

func main() {
	arg := os.Args[1]
	renderTemplate("template.tmpl", readFile(arg))
	writeTemplateToFile("template.tmpl", arg)
}
