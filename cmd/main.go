package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Conn struct {
	file *os.File
}
 
func main() {

	//args := os.Args
	dbPointer := flag.String("db", "./test.db", "Location of the db file.")
	flag.Parse()


	fmt.Println(*dbPointer)
	file, err := os.OpenFile(*dbPointer, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	var connection = Conn {
		file,
	}
	fmt.Println(connection)
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
