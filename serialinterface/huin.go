package serialinterface

import (
	"github.com/huin/goserial"
	"io"
)

const (
	HUIN_INTERFACE = "huin"
)

type HuinProvider struct {
	huin io.ReadWriteCloser
}

func NewHuinInterface(comport string, baud int, readTimeout int) (HuinProvider,error) {
	var err error

	t := HuinProvider{}

	c := &goserial.Config{Name: comport, Baud: baud}
	t.huin, err = goserial.OpenPort(c)

	return t,err
}

func (t HuinProvider) Close() error {
	return t.huin.Close()
}

func (t HuinProvider) Read(p []byte) (int, error) {
	return t.huin.Read(p)
}

func (t HuinProvider) Write(p []byte) (n int, err error) {
	return t.huin.Write(p)
}
