package internal

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func readWordfile(fileName string) []string {
	wordFile, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	items := strings.Split(string(wordFile), "\n")
	fmt.Printf("Loaded %d words from %s\n", len(items), fileName)

	return items
}

func generateAndWriteWordfile(wordlistFileName, generatedFileName, listName string) {
	tmpl, err := template.ParseFiles("internal/templates/wordlist.tmpl")
	if err != nil {
		panic(err)
	}

	items := readWordfile(wordlistFileName)

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, map[string]any{
		"ListName": listName,
		"Words":    items,
	}); err != nil {
		panic(err)
	}

	if err := os.WriteFile(generatedFileName, buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}
}

func ProcessWordList() {
	generateAndWriteWordfile(
		"internal/lists/adjectives.txt",
		"adjectives.go",
		"Adjectives",
	)
	generateAndWriteWordfile(
		"internal/lists/nouns.txt",
		"nouns.go",
		"Nouns",
	)
}
