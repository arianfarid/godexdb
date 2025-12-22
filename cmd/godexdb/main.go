package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
		input, err := reader.ReadString(';')
		if err != nil {
			log.Fatal("Syntax error: ", err)
		}
		fmt.Printf("%s\n", input)
	}
}
