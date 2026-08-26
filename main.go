//file organizer

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(" Please provide a directory path as an argument.")
		os.Exit(1)
	}

	entries, err := os.ReadDir(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		fmt.Print(entry.Name())

		if entry.IsDir() {
			fmt.Printf("  -   Directory\n")
		} else {
			fmt.Printf("  -   %s\n", extIndex(ext))
		}
		fmt.Println("")
	}
}

func extIndex(ext string) string {
	switch ext {
	case ".txt":
		return "Text File"
	case ".pdf":
		return "PDF Document"
	case ".jpg", ".jpeg", ".png", ".gif":
		return "Image File"
	case ".mod":
		return "Go Module File"
	default:
		return "Unknown File Type"
	}
}
