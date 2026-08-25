//file organizer

package main

import (
	"fmt"
	"log"
	"os"
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
		fmt.Print(entry.Name())

		if entry.IsDir() {
			fmt.Printf("  -   Directory\n")
			// } else {
			// 	info, err := entry.Info()
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	fmt.Printf("\n %v", info)
		}
		fmt.Println("")
	}
}
