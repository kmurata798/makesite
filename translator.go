package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	_ "cloud.google.com/go/translate/apiv3"
	"golang.org/x/text/language"
)

func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	// fmt.Println(resp[0].Text, nil)
	return resp[0].Text, nil
} 

// func translator(text string) {
// var lang string
// flag.StringVar(&lang, "lang", "es", "This is the language you want to translate, inputting google's language abbreviations.")
// flag.Parse()

// 	fmt.Println("Language:", lang)

// 	ctx := context.Background()

// 	// Creates a client.
// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	// Sets the text to translate.
// 	// text := "Hello, world!"
// 	// Sets the target language.
// 	target, err := language.Parse(lang)
// 	if err != nil {
// 		log.Fatalf("Failed to parse target language: %v", err)
// 	}

// 	// Translates the text into Russian.
// 	translations, err := client.Translate(ctx, []string{text}, target, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to translate text: %v", err)
// 	}

// 	fmt.Printf("Text: %v\n", text)
// 	fmt.Printf("Translation: %v\n", translations[0].Text)
// }
