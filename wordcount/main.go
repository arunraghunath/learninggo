package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Inside main function - word count")
	filename := os.Args[1]
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	contents := string(text)
	lsub := strings.Fields(contents)
	fmt.Printf("Number of words is %d", len(lsub))

}
