package main

import (
	"fmt"
	"log"
	"os"

	"chapter3/words"
)

func main() {
	filename := os.Args[1]

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text.\n", count)
}
