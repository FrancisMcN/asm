package load_commands

import "encoding/binary"

type LcMain struct {
	header    LcHeader
	entryoff  uint64
	stacksize uint64
}

func (l LcMain) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, l.header.Bytes()...)
	for i := 0; i < 16; i++ {
		bytes = append(bytes, 0)
	}
	offset := 8
	binary.LittleEndian.PutUint64(bytes[offset:], l.entryoff)
	offset += 8
	binary.LittleEndian.PutUint64(bytes[offset:], l.stacksize)
	return bytes
}

func (l LcMain) SizeOf() uint32 {
	return uint32(len(l.Bytes()))
}

func NewLcMain(entryoff, stacksize uint64) LcMain {
	main := LcMain{
		header: LcHeader{
			Cmd: MainCommand,
		},
		entryoff:  entryoff,
		stacksize: stacksize,
	}
	main.header.CmdSize = main.SizeOf()
	return main
}
