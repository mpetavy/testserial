package serialinterface

import (
	"github.com/tarm/serial"
	"time"
)

const (
	TARM_INTERFACE = "tarm"
)

type TarmProvider struct {
	tarm *serial.Port
}

func NewTarmProvider(comport string, baud int,readTimeout int) (TarmProvider,error) {
	var err error

	t := TarmProvider{}

	c := &serial.Config{Name: comport, Baud:baud, ReadTimeout: time.Second * time.Duration(readTimeout)}
	t.tarm, err = serial.OpenPort(c)

	return t,err
}

func (t TarmProvider) Close() error {
	return t.tarm.Close()
}

func (t TarmProvider) Read(p []byte) (int, error) {
	return t.tarm.Read(p)
}

func (t TarmProvider) Write(p []byte) (n int, err error) {
	return t.tarm.Write(p)
}
