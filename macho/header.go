package macho

import "encoding/binary"

const (
	MhExecute = 2
)

type Header struct {
	Magic      uint32
	CpuType    uint32
	CpuSubType uint32
	FileType   uint32
	NoOfCmds   uint32
	SizeOfCmds uint32
	Flags      uint32
	Reserved   uint32
}

func (h Header) Bytes() []byte {
	bytes := make([]byte, 32)

	binary.LittleEndian.PutUint32(bytes, h.Magic)
	binary.LittleEndian.PutUint32(bytes[4:], h.CpuType)
	binary.LittleEndian.PutUint32(bytes[8:], h.CpuSubType)
	binary.LittleEndian.PutUint32(bytes[12:], h.FileType)
	binary.LittleEndian.PutUint32(bytes[16:], h.NoOfCmds)
	binary.LittleEndian.PutUint32(bytes[20:], h.SizeOfCmds)
	binary.LittleEndian.PutUint32(bytes[24:], h.Flags)
	binary.LittleEndian.PutUint32(bytes[28:], h.Reserved)
	return bytes
}
