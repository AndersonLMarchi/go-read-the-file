package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	validator "./src/utils"
)

var next bool

func main() {

	scanner := readFile()

	// Default scanner is bufio.ScanLines. Lets get all line data
	scanner.Split(bufio.ScanLines)

	// Scan for next token and start reading from 3rd line because
	// 1st and 2nd lines are not important for the File Reading Service
	scanner.Scan()
	scanner.Scan()
	next = scanner.Scan()

	if next == false {
		// False on error or EOF. Check error
		err := scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	for {

		line := splitLine(scanner.Text())
		printLineResult(line)
		next = scanner.Scan() // Call scanner.Scan() again to find next token
		if !next {
			break
		}
	}
}

func readFile() *bufio.Scanner {
	// Open file and create scanner on top of it
	// file, err := os.Open("./repository/base_teste.txt") // complete file
	file, err := os.Open("./repository/base_teste_short.txt") //short file version for test
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)

}

func splitLine(line string) []string {

	// slice each word into an array
	array := strings.Fields(line)

	return array
}

func printLineResult(line []string) {

	fmt.Println("================================================================================")
	if validator.IsCPF(line[0]) {
		fmt.Println("CPF:", line[0])
		fmt.Println("PRIVATE:", line[1])
		fmt.Println("INCOMPLETO:", line[2])
		fmt.Println("DATA ULTIMA COMPRA:", line[3])
		fmt.Println("COMPRA TICKET MEDIO:", line[4])
		fmt.Println("TICKET DA ULTIMA COMPRA:", line[5])
		fmt.Println("LOJA MAIS FREQUENTE:", line[6])
		fmt.Println("LOJA DA ULTIMA COMPRA:", line[7])
	}
	fmt.Println("")
}
