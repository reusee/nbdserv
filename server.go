package nbdserv

import (
	"encoding/binary"
	"net"
)

type Backend interface {
	Size() uint64
}

func Start(
	conn net.Conn,
	backend Backend,
) (err error) {
	defer he(&err)

	// handshake
	_, err = conn.Write([]byte("NBDMAGIC"))
	ce(err)
	_, err = conn.Write([]byte("IHAVEOPT"))
	ce(err)
	var handshakeFlags uint16
	err = binary.Write(conn, binary.BigEndian, handshakeFlags)
	var clientFlags uint32
	err = binary.Write(conn, binary.BigEndian, clientFlags)

	// options

	return
}
