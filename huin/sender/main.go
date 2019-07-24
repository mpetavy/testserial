package main

import (
	"github.com/huin/goserial"
	"fmt"
	"log"
	"os"
)

func main() {
	c := &goserial.Config{Name: os.Args[1], Baud: 115200}
	s, err := goserial.OpenPort(c)
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

	_, err = s.Write([]byte(fmt.Sprintf("%s%c", "ÄÖÜßäöü - This is a very long line # This is a very long line # This is a very long line # This is a very long line !", 10)))
	if err != nil {
		log.Fatal(err)
	}
}
