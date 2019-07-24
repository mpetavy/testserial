package main

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
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

	log.Printf("Reading ...\n")

	var n int
	var a int

	buf := make([]byte, 128)
	for {
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			a += n

			fmt.Printf("%s", string(buf[:n]))
		} else {
			fmt.Printf(".")
		}
	}

	log.Printf("%d bytes read", a)
}
