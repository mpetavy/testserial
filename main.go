package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

var (
	mode    *string
	text    *string
	comport *string
	library *string
	baud    *int
	jacobsa io.ReadWriteCloser
)

func init() {
	mode = flag.String("m", "", "Operation mode READ/WRITE")
	comport = flag.String("c", "", "COM port to use")
	text = flag.String("t", "Hello!", "Text to transmit")
	baud = flag.Int("b", 9600, "Baud rate")
}

func run() error {
	flag.Parse()

	if *mode == "" {
		flag.Usage()
		os.Exit(0)
	}

	var err error
	var i int

	options := serial.OpenOptions{
		PortName:          *comport,
		BaudRate:          uint(*baud),
		DataBits:          8,
		StopBits:          1,
		ParityMode:        0,
		RTSCTSFlowControl: false,

		InterCharacterTimeout:   0,
		MinimumReadSize:         1,
		Rs485Enable:             false,
		Rs485RtsHighDuringSend:  false,
		Rs485RtsHighAfterSend:   false,
		Rs485RxDuringTx:         false,
		Rs485DelayRtsBeforeSend: 0,
		Rs485DelayRtsAfterSend:  0,
	}

	for i = 0; i < 20; i++ {
		log.Printf("try #%d to open %s ...", i, *comport)
		jacobsa, err = serial.Open(options)

		if err != nil {
			time.Sleep(time.Millisecond * 100)
		} else {
			break
		}
	}

	if err != nil {
		return err
	}

	log.Printf("open successfull after %d tries", i)

	defer func() {
		err := jacobsa.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	switch strings.ToUpper(*mode) {
	case "READ":
		log.Printf("read")
		err = read()
	case "WRITE":
		log.Printf("write")
		err = write()
	default:
		return fmt.Errorf("unknown mode: %s", *mode)
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
		n, err = jacobsa.Read(buf)
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

	return nil

}

func write() error {
	log.Printf("Printing ...\n")

	var err error

	_, err = jacobsa.Write([]byte(*text))

	return err
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
