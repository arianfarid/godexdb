package db

import (
	"log"
	"os"
)

type Conn struct {
	file *os.File
}
func (c Conn) Connect(path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	c.file = file
} 
