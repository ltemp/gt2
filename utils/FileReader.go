package utils

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

type Parser interface {
	parseRow()
}



func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
