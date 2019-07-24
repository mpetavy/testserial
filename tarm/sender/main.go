package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"os"
)
func main() {
	c := &serial.Config{Name: os.Args[1], Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := s.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Printing ...\n")

	_, err = s.Write([]byte(fmt.Sprintf("%s%c","ÄÖÜßäöü - This is a very long line # This is a very long line # This is a very long line # This is a very long line !",10)))
	if err != nil {
		log.Fatal(err)
	}
}


