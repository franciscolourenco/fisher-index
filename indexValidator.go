package main

import (
	"os"
	"bufio"
	"log"
)

func indexValidator(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var errorCount, lineNumber int = 0, 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalln(err)
		}

		line := indexLine{lineNumber, indexCode(lineNumber % 6), scanner.Text()}
		if err := line.verify(); err != nil {
			errorCount++;
			log.Println("not ok", err)
		} else {
			log.Println("ok", line.no+1, line.text)
		}

		lineNumber++
	}

	return lineNumber, errorCount;
}
