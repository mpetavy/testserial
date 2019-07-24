package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"os"
	"time"
)

func main() {
	c := &serial.Config{Name: os.Args[1], Baud: 115200, ReadTimeout: time.Second * 1}
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

			fmt.Printf("%s",string(buf[:n]))
		} else {
			fmt.Printf(".")
		}
	}

	log.Printf("%d bytes read", a)
}
