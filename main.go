package main

import (
	"flag"
	"fmt"
	"github.com/mpetavy/common"
	"go.bug.st/serial"
	"os"
)

var (
	readMode *bool
	text     *string
	comport  *string
	fileName *string
	baud     *int
)

func init() {
	common.Init(false, "1.0.0", "", "2018", "test", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)

	readMode = flag.Bool("r", false, "Operation mode READ/WRITE")
	comport = flag.String("c", "", "COM port to use")
	text = flag.String("t", "Hello!", "Text to transmit")
	baud = flag.Int("b", 9600, "Baud rate")
	fileName = flag.String("f", "", "File to transmit")
}

func run() error {
	ports, err := serial.GetPortsList()
	if common.Error(err) {
		return err
	}
	if len(ports) == 0 {
		return fmt.Errorf("No serial ports found!")
	}
	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	options := &serial.Mode{
		BaudRate: *baud,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(*comport, options)
	if common.Error(err) {
		return err
	}

	defer func() {
		common.Error(port.Close())
	}()

	if *readMode {
		common.Info("read")
		err = read(port)
	} else {
		common.Info("write")
		err = write(port)
	}

	return err
}

func read(port serial.Port) error {
	common.Info("Reading ...\n")

	var err error
	var n int

	buf := make([]byte, 128)

	for {
		n, err = port.Read(buf)

		portError, ok := err.(*serial.PortError)
		if ok && portError.Code() == serial.PortClosed {
			return nil
		}

		if common.Error(err) {
			return nil
		}

		fmt.Printf("%s", string(buf[:n]))
	}
}

func write(port serial.Port) error {
	common.Info("Printing ...\n")

	var err error

	if *fileName != "" {
		ba, err := os.ReadFile(*fileName)
		if common.Error(err) {
			return err
		}

		*text = string(ba)
	}

	_, err = port.Write([]byte(*text))
	if common.Error(err) {
		return err
	}

	return nil
}

func main() {
	defer common.Done()

	common.Run([]string{"c"})
}
