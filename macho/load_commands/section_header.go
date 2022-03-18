package load_commands

import "encoding/binary"

type SectionHeader64 struct {
	SectName  []byte
	SegName   []byte
	Addr      uint64
	Size      uint64
	Offset    uint32
	Align     uint32
	Reloff    uint32
	Nreloc    uint32
	Flags     uint32
	Reserved1 uint32
	Reserved2 uint32
	Reserved3 uint32
}

func (s SectionHeader64) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, s.SectName...)
	bytes = append(bytes, s.SegName...)
	for i := 0; i < 48; i++ {
		bytes = append(bytes, 0)
	}
	binary.LittleEndian.PutUint64(bytes[32:], s.Addr)
	binary.LittleEndian.PutUint64(bytes[40:], s.Size)
	binary.LittleEndian.PutUint32(bytes[48:], s.Offset)
	binary.LittleEndian.PutUint32(bytes[52:], s.Align)
	binary.LittleEndian.PutUint32(bytes[56:], s.Reloff)
	binary.LittleEndian.PutUint32(bytes[60:], s.Nreloc)
	binary.LittleEndian.PutUint32(bytes[64:], s.Flags)
	binary.LittleEndian.PutUint32(bytes[68:], s.Reserved1)
	binary.LittleEndian.PutUint32(bytes[72:], s.Reserved2)
	binary.LittleEndian.PutUint32(bytes[76:], s.Reserved3)
	return bytes
}

func (s SectionHeader64) SizeOf() uint32 {

	size := 0
	size += len(s.SegName)
	size += len(s.SectName)
	size += 16
	size += 32
	return uint32(size)

}

func NewSectionHeader64(sectName string, segName string, addr uint64, size uint64, offset uint32, align uint32, reloff uint32, nreloc uint32, flags uint32) SectionHeader64 {

	segmentName := make([]byte, 16)
	for i := 0; i < len(segName) && i < 16; i++ {
		segmentName[i] = segName[i]
	}

	sectionName := make([]byte, 16)
	for i := 0; i < len(sectName) && i < 16; i++ {
		sectionName[i] = sectName[i]
	}

	return SectionHeader64{
		SegName:  segmentName,
		SectName: sectionName,
		Addr:     addr,
		Size:     size,
		Offset:   offset,
		Align:    align,
		Reloff:   reloff,
		Nreloc:   nreloc,
		Flags:    flags,
	}
}
