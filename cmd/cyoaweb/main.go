package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}
	story, err := cyoa.JsonStory(f)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v \n", story)
}
