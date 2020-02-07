package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type content struct {
	Description string
}

func readFile(name string) string {
	/*
		Makesite MVP

		collects data from file
	*/
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(fileContents)

}

func writeFile(name string, data string) {
	/*
		Makesite MVP

		Writes data onto file
	*/
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func renderTemplate(filename string, data string) {
	/*
		Makesite MVP

		Print out .html template onto the terminal to check
	*/
	c := content{Description: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	err := t.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}

func writeTemplateToFile(templateName string, fileName string) {
	/*
		MakeSite MVP

		Creates new template with the filename given
	*/
	c := content{Description: readFile(fileName)}
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))

	filter := filterInput(fileName) //option 1
	// f, err := os.Create(arg[:len(arg)-4] + ".html") //option 2
	f, err := os.Create(filter)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}
}

func filterInput(input string) string {
	/*
		Makesite v1.1

			Traverse through input until you reach '.', then add '.html' to the end.
			return s
	*/
	ext := ".html"
	s := strings.Split(input, ".")[0] + ext
	return s
}

func parser() {
	/*
		Makesite v1.1

			Collects files in given directory,
			checks if file includes '.txt',
			creates a '.html' file for the .txt files
	*/
	var dir string
	flag.StringVar(&dir, "dir", "", "this is the directory")
	flag.Parse()

	fmt.Println("Directory:", dir)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filenameCheck(file.Name()) == true {
			fmt.Println(file.Name())
			writeTemplateToFile("template.tmpl", file.Name())
		}
	}
}

func filenameCheck(filename string) bool {
	/*
		makesite v1.1

		checks if filename includes .txt,
		if so, returns True
		else, returns false
	*/
	tail := "txt"
	for i := range filename {
		if filename[i] == '.' {
			s := strings.Split(filename, ".")[1]
			// fmt.Println(s)
			if s == tail {
				return true
			}
		}
	}
	return false
}

func main() {
	// arg := os.Args[1] // Makesite MVP
	parser() //Makesite v1.1
	// renderTemplate("template.tmpl", readFile(arg)) //makesite MVP
	// writeTemplateToFile("template.tmpl", arg) //makesite MVP
}
