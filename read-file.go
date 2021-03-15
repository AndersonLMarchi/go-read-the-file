package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	connection "go-reader/src/connection"
	validator "go-reader/src/utils"
)

var fileLines connection.FileLines
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

		line := strings.Fields(scanner.Text())
		insertLine(line)
		next = scanner.Scan() // Call scanner.Scan() again to find next line
		if !next {
			break
		}
	}

	if len(fileLines) > 0 {
		if connection.Insert(&fileLines) {
			fmt.Println("Arquivo importado com Sucesso!")
		}
	}
	findImportedFile()

}

func readFile() *bufio.Scanner {
	// Open file and create scanner on top of it
	// file, err := os.Open("./repository/base_teste_short.txt") //short file version for test
	file, err := os.Open("./repository/base_teste.txt") // complete file
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)

}

func insertLine(line []string) {

	// validates if CPF/CNPJ fields on file is valid from database persistence
	if validator.IsCPF(line[0]) && validator.IsCNPJ(line[6]) && validator.IsCNPJ(line[7]) {

		// convert strings fields
		private, _ := strconv.ParseBool(line[1])
		incompleto, _ := strconv.ParseBool(line[2])
		ticketMedio, _ := strconv.ParseFloat(line[4], 64)
		ticketUltimaCompra, _ := strconv.ParseFloat(line[5], 64)

		fileLine := connection.FileType{}
		fileLine.Cpf = line[0]
		fileLine.Private = private
		fileLine.Incompleto = incompleto
		fileLine.DataUltimaCompra = line[3]
		fileLine.CompraTicketMedio = ticketMedio
		fileLine.TicketUltimaCompra = ticketUltimaCompra
		fileLine.CnpjLojaFrequente = line[6]
		fileLine.CnpjUltimaLoja = line[7]

		fileLines = append(fileLines, &fileLine)

	}
}

func findImportedFile() {

	fileLines = connection.FindAll()
	for _, fileType := range fileLines {
		fmt.Println("==========================================================")
		fmt.Println("CPF:", fileType.Cpf)
		fmt.Println("PRIVATE:", fileType.Private)
		fmt.Println("INCOMPLETO:", fileType.Incompleto)
		fmt.Println("DATA ULTIMA COMPRA:", fileType.DataUltimaCompra)
		fmt.Println("COMPRA TICKET MEDIO:", fileType.CompraTicketMedio)
		fmt.Println("TICKET DA ULTIMA COMPRA:", fileType.TicketUltimaCompra)
		fmt.Println("LOJA MAIS FREQUENTE:", fileType.CnpjLojaFrequente)
		fmt.Println("LOJA DA ULTIMA COMPRA:", fileType.CnpjUltimaLoja)
		fmt.Println("")
	}

}
