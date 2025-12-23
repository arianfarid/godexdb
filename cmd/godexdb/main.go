package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arianfarid/godexdb/internal/db"
)

func main() {

	//args := os.Args
	dbPointer := flag.String("db", "./test.db", "Location of the db file.")
	flag.Parse()

	var connection db.Conn
	connection.Connect(*dbPointer)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		raw, err := reader.ReadString(';')
		if err != nil {
			log.Fatal("Syntax error: ", err)
		}


		input := db.ParseInput(raw)
		switch input.Type {
		case db.InputEmpty:
			continue
		case db.InputMeta:
			if input.Meta.Type == db.MetaCommandQuit {
				return
			}
		case db.InputSQL:
			
			vals := strings.Split(input.SQL, " ")
			first := strings.ToLower(vals[0])

			statement := db.ParseStatemnt(first)
				
			fmt.Println("SQL:", input.SQL)
			fmt.Println("Statment:", statement)

			 
		}
		
	}
}
