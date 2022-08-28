package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Packet struct {
	packet_length    uint32
	header_length    uint16
	protocol_version uint16
	operation        uint32
	sequence_id      uint32
	body             []byte
}

func decode(rd io.Reader) (*Packet, error) {
	var packet = Packet{}

	error := binary.Read(rd, binary.BigEndian, &packet.packet_length)
	if error != nil {
		return nil, error
	}

	error = binary.Read(rd, binary.BigEndian, &packet.header_length)
	if error != nil {
		return nil, error
	}

	error = binary.Read(rd, binary.BigEndian, &packet.protocol_version)
	if error != nil {
		return nil, error
	}

	error = binary.Read(rd, binary.BigEndian, &packet.operation)
	if error != nil {
		return nil, error
	}

	error = binary.Read(rd, binary.BigEndian, &packet.sequence_id)
	if error != nil {
		return nil, error
	}

	if packet.packet_length < uint32(packet.header_length) {
		return nil, errors.New("packet length overflow")
	}

	var body_len = packet.packet_length - uint32(packet.header_length)
	packet.body = make([]byte, body_len)

	if body_len > 0 {
		n, error := rd.Read(packet.body)
		if error != nil {
			return nil, error
		}

		if n != int(body_len) {
			return nil, fmt.Errorf("packet read body fail, expect %d, but %d", int(body_len), n)
		}
	}

	return &packet, nil
}
