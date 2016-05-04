package main

import (
	"log"
	"flag"
	"os"
)

func main() {
	flag.Parse()

	lineCount, errorCount := indexValidator(flag.Arg(0))

	log.Printf("\n1..%d\n# Tests %d\n# Pass  %d", lineCount, lineCount, lineCount-errorCount);

	if errorCount != 0 {
		log.Println("\nnot ok");
		os.Exit(1)
	}
	log.Println("\nok");
}
