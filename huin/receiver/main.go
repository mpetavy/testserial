package main

import (
	"fmt"
	"github.com/huin/goserial"
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
