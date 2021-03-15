package connection

import (
	"fmt"
	"log"
	"strings"
)

var err error
var db = Connect()

func Insert(lines *FileLines) bool {

	if createTicketLojaTable() {
		var values []string

		for _, line := range *lines {
			str := "("
			str += fmt.Sprintf("%s,", line.Cpf)
			if line.Private {
				str += "true"
			} else {
				str += "false"
			}
			if line.Incompleto {
				str += "true"
			} else {
				str += "false"
			}
			str += fmt.Sprintf("%s,", line.DataUltimaCompra)
			str += fmt.Sprintf("%.2f,", line.CompraTicketMedio)
			str += fmt.Sprintf("%.2f,", line.TicketUltimaCompra)
			str += fmt.Sprintf("%s,", line.CnpjLojaFrequente)
			str += fmt.Sprintf("%s", line.CnpjUltimaLoja)
			str += ")"
			values = append(values, str)
		}

		sql := "INSERT INTO ticketloja (cpf, private, incompleto, dataultcompra, ticketmedio, ticketultcompra, cnpjlojafreq, cnpjultloja) VALUES $1"

		_, err = db.Exec(sql, strings.Join(values, ", "))
		if err != nil {
			return false
		} else {
			return true
		}
	}
	return false
}

func FindAll() FileLines {
	sql := `SELECT * FROM ticketloja`

	var fileType FileType
	var fileLines FileLines

	// example from http://go-database-sql.org/retrieving.html
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&fileType.Cpf, &fileType.Private, &fileType.Incompleto, &fileType.DataUltimaCompra, &fileType.CompraTicketMedio, &fileType.TicketUltimaCompra, &fileType.CnpjLojaFrequente, &fileType.CnpjUltimaLoja)
		if err != nil {
			log.Fatal(err)
		}
		fileLines = append(fileLines, &fileType)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		return fileLines
	}
	return nil

}

//after spending the trial period at Neoway I finish this function :P
func Find(filter []string, values []string) {

}

//creates table to persist files content
func createTicketLojaTable() bool {

	sql := `CREATE TABLE IF NOT EXISTS ticketloja (
				seq serial PRIMARY KEY,
				cpf varchar(14),
				private boolean,
				incompleto boolean,
				dataultcompra date,
				ticketmedio numeric(17,2),
				ticketultcompra numeric(17,2),
				cnpjlojafreq varchar(18),
				cnpjultloja varchar(18)
			)`

	_, err = db.Exec(sql)
	if err != nil {
		return false
	} else {
		return true
	}
}
