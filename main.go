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
	header()

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

// extIndex returns a description based on file extension
func extIndex(ext string) string {
	imgExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	docExtensions := []string{".pdf", ".doc", ".docx", ".txt", ".xls", ".xlsx"}
	audioExtensions := []string{".mp3", ".wav", ".flac"}
	videoExtensions := []string{".mp4", ".avi", ".mkv"}
	codeExtensions := []string{".go", ".py", ".js", ".java", ".c", ".cpp", ".mod"}

	switch {
	case contains(imgExtensions, ext):
		return "Image File"
	case contains(docExtensions, ext):
		return "Document File"
	case contains(audioExtensions, ext):
		return "Audio File"
	case contains(videoExtensions, ext):
		return "Video File"
	case contains(codeExtensions, ext):
		return "Code File"
	default:
		return "Unknown File Type"
	}
}

// contains checks if a string slice contains a given string
func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func header() {
	fmt.Println("File Organizer")
	fmt.Println("================")
}
