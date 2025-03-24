package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pages := []struct {
		Name string
		Content string
	}{
		{"index.html", "<html><body><h1>Index</h1></body></html>"},
		{"about.html", "<html><body><h1>About</h1></body></html>"},
		{"contact.html", "<html><body><h1>Contact</h1></body></html>"},
	}

	outputDir := "docs"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pages := range pages {
		filePath := filepath.Join(outputDir, pages.Name)
		err := os.WriteFile(filePath, []byte(pages.Content), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("Created:", filePath)
	}
	fmt.Println("Done")
}