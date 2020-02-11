package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func translator(text string) {
	var lang string
	flag.StringVar(&lang, "lang", "es", "This is the language you want to translate, inputting google's language abbreviations.")
	flag.Parse()

	fmt.Println("Language:", lang)

	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to translate.
	// text := "Hello, world!"
	// Sets the target language.
	target, err := language.Parse(lang)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)
}
