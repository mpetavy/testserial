package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"testserial/serialinterface"
	"time"
)

var (
	mode        *string
	text        *string
	com         *string
	library     *string
	baud        *int
	readTimeout *int

	serialIf serialinterface.Provider
)

func init() {
	mode = flag.String("m", "", "Operation mode READ/WRITE")
	com = flag.String("c", "", "COM port to use")
	library = flag.String("l", "jacobsa", "Library to use")
	text = flag.String("t", "Hello!", "Text to transmit")
	baud = flag.Int("b", 19200, "Baud rate")
	readTimeout = flag.Int("rt", 1, "READ timeout")
}

func run() error {
	flag.Parse()

	if *mode == "" {
		flag.Usage()
		os.Exit(0)
	}

	var err error
	var i int

	for i = 0; i < 20; i++ {
		switch *library {
		case serialinterface.JACOBSA_INTERFACE:
			serialIf, err = serialinterface.NewJacobsaProvider(*com, *baud, *readTimeout)
		case serialinterface.TARM_INTERFACE:
			serialIf, err = serialinterface.NewTarmProvider(*com, *baud, *readTimeout)
		case serialinterface.HUIN_INTERFACE:
			serialIf, err = serialinterface.NewHuinInterface(*com, *baud, *readTimeout)
		default:
			return fmt.Errorf("unknown library: %s", *library)
		}

		if err != nil {
			time.Sleep(time.Millisecond * 100)
		} else {
			break
		}
	}

	fmt.Printf("after init %v\n", err)

	if err != nil {
		return err
	}

	fmt.Printf("succeeded to open after %d tries\n", i)

	defer func() {
		fmt.Printf("defer")
		err := serialIf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Printf("after defer")

	switch strings.ToUpper(*mode) {
	case "READ":
		fmt.Printf("read")
		err = read()
	case "WRITE":
		fmt.Printf("write")
		err = write()
	default:
		return fmt.Errorf("unknown mode: %s", *library)
	}

	return nil
}

func read() error {
	log.Printf("Reading ...\n")

	var err error
	var n int
	var a int

	buf := make([]byte, 128)
	for {
		n, err = serialIf.Read(buf)
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

	return nil

}

func write() error {
	log.Printf("Printing ...\n")

	var err error

	_, err = serialIf.Write([]byte(*text))

	return err
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
