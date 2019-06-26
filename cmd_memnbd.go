// +build ignore

package main

import (
	"fmt"
	"net"

	"github.com/reusee/nbdserv"
)

func main() {
	ln, err := net.Listen(
		"tcp",
		":10809",
	)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("accept: %v", err)
			continue
		}
		go func() {
			defer conn.Close()

			if err := nbdserv.Start(
				conn,
				NewMem(124*1024*1024*64),
			); err != nil {
				panic(err)
			}

		}()
	}
}

type Mem struct {
	size uint64
}

var _ nbdserv.Backend = new(Mem)

func NewMem(size uint64) *Mem {
	return &Mem{
		size: size,
	}
}

func (m Mem) Size() uint64 {
	return m.size
}
