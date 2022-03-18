package load_commands

import "encoding/binary"

type LcHeader struct {
	Cmd     uint32
	CmdSize uint32
}

func (l LcHeader) Bytes() []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(bytes, l.Cmd)
	binary.LittleEndian.PutUint32(bytes[4:], l.CmdSize)
	return bytes
}

func (l LcHeader) SizeOf() uint32 {
	return 8
}
