package serialinterface

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
)

const (
	JACOBSA_INTERFACE = "jacobsa"
)

type JacobsaProvider struct {
	jacobsa io.ReadWriteCloser
}

func NewJacobsaProvider(comport string, baud int, readTimeout int) (Provider, error) {
	t := JacobsaProvider{}

	options := serial.OpenOptions{
		PortName:              comport,
		BaudRate:              uint(baud),
		DataBits:              8,
		StopBits:              1,
		InterCharacterTimeout: uint(readTimeout),
		MinimumReadSize:       0,
	}

	var err error

	t.jacobsa, err = serial.Open(options)

	return t, err
}

func (t JacobsaProvider) Close() error {
	return t.jacobsa.Close()

}

func (t JacobsaProvider) Read(p []byte) (int, error) {
	return t.jacobsa.Read(p)
}

func (t JacobsaProvider) Write(p []byte) (n int, err error) {
	return t.jacobsa.Write(p)
}
