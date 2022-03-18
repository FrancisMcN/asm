package load_commands

import (
	"encoding/binary"
)

type LcSegment64 struct {
	header         LcHeader
	SegName        []byte
	VmAddr         uint64
	VmSize         uint64
	FileOffset     uint64
	FileSize       uint64
	MaxProtection  uint32
	InitProtection uint32
	NoOfSections   uint32
	Flags          uint32
	SectionHeaders []SectionHeader64
}

func (s LcSegment64) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, s.header.Bytes()...)
	bytes = append(bytes, s.SegName...)
	for i := 0; i < 48; i++ {
		bytes = append(bytes, 0)
	}
	binary.LittleEndian.PutUint64(bytes[8+16:], s.VmAddr)
	binary.LittleEndian.PutUint64(bytes[8+24:], s.VmSize)
	binary.LittleEndian.PutUint64(bytes[8+32:], s.FileOffset)
	binary.LittleEndian.PutUint64(bytes[8+40:], s.FileSize)
	binary.LittleEndian.PutUint32(bytes[8+48:], s.MaxProtection)
	binary.LittleEndian.PutUint32(bytes[8+52:], s.InitProtection)
	binary.LittleEndian.PutUint32(bytes[8+56:], s.NoOfSections)
	binary.LittleEndian.PutUint32(bytes[8+60:], s.Flags)

	for _, sh := range s.SectionHeaders {
		bytes = append(bytes, sh.Bytes()...)
	}

	return bytes
}

func (s LcSegment64) SizeOf() uint32 {
	return uint32(len(s.Bytes()))
}

func NewSegment64(name string, vmAddr uint64, vmSize uint64, fileOffset uint64, fileSize uint64, maxProtection uint32, initProtection uint32, sectionHeaders []SectionHeader64) LcSegment64 {

	segName := make([]byte, 16)
	for i := 0; i < len(name) && i < 16; i++ {
		segName[i] = name[i]
	}

	var noOfSections uint32 = 0
	if sectionHeaders != nil {
		noOfSections = uint32(len(sectionHeaders))
	}

	if sectionHeaders == nil {
		sectionHeaders = make([]SectionHeader64, 0)
	}

	segment := LcSegment64{
		header: LcHeader{
			Cmd: Segment64Command,
		},
		SegName:        segName,
		VmAddr:         vmAddr,
		VmSize:         vmSize,
		FileOffset:     fileOffset,
		FileSize:       fileSize,
		MaxProtection:  maxProtection,
		InitProtection: initProtection,
		NoOfSections:   noOfSections,
		Flags:          0,
		SectionHeaders: sectionHeaders,
	}
	segment.header.CmdSize = segment.SizeOf()
	return segment

}
