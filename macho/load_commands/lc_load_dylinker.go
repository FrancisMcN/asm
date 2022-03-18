package load_commands

import "encoding/binary"

type LcLoadDylinker struct {
	header    LcHeader
	StrOffset uint32
	StrName   string
}

func (l LcLoadDylinker) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, l.header.Bytes()...)
	for i := 0; i < 4; i++ {
		bytes = append(bytes, 0)
	}
	binary.LittleEndian.PutUint32(bytes[8:], l.StrOffset)
	for i := 0; i < len(l.StrName); i++ {
		bytes = append(bytes, l.StrName[i])
	}
	p := len(bytes) % 32
	// Pad binary if size isn't a multiple of 32
	for ; p < 32; p++ {
		bytes = append(bytes, 0)
	}
	return bytes
}

func (l LcLoadDylinker) SizeOf() uint32 {
	return uint32(len(l.Bytes()))
}

func NewLcLoadDynlinker(name string) LcLoadDylinker {
	lcLoadDylinker := LcLoadDylinker{
		header: LcHeader{
			Cmd: LoadDynlinkerCommand,
		},
		StrOffset: 12,
		StrName:   name,
	}
	lcLoadDylinker.header.CmdSize = lcLoadDylinker.SizeOf()
	return lcLoadDylinker
}
