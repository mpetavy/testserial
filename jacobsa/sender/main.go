package main

import (
	"github.com/jacobsa/go-serial/serial"
	"fmt"
	"log"
	"os"
)

func main() {
	options := serial.OpenOptions{
		PortName:        os.Args[1],
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	s, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
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
