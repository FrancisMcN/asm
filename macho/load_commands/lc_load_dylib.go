package load_commands

import "encoding/binary"

type LcLoadDylib struct {
	header               LcHeader
	StrOffset            uint32
	Timestamp            uint32
	CurrentVersion       uint32
	CompatibilityVersion uint32
	StrName              string
}

func (l LcLoadDylib) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, l.header.Bytes()...)
	for i := 0; i < 16; i++ {
		bytes = append(bytes, 0)
	}
	offset := 8
	binary.LittleEndian.PutUint32(bytes[offset:], l.StrOffset)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.Timestamp)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.CurrentVersion)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.CompatibilityVersion)

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

func (l LcLoadDylib) SizeOf() uint32 {
	return uint32(len(l.Bytes()))
}

func NewLcLoadDylib(name string) LcLoadDylib {
	lcLoadDylib := LcLoadDylib{
		header: LcHeader{
			Cmd: LoadDylibCommand,
		},
		StrOffset:            24,
		Timestamp:            0,
		CurrentVersion:       0,
		CompatibilityVersion: 0,
		StrName:              name,
	}
	lcLoadDylib.header.CmdSize = lcLoadDylib.SizeOf()
	return lcLoadDylib
}
