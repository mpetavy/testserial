package serialinterface

import "io"

type Provider interface {
	io.ReadWriteCloser
}
