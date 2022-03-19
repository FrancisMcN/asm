package load_commands

import "encoding/binary"

type LcDySymTab struct {
	header         LcHeader
	ilocalsym      uint32
	nlocalsym      uint32
	iextdefsym     uint32
	nextdefsym     uint32
	iundefsym      uint32
	nundefsym      uint32
	tocoff         uint32
	ntoc           uint32
	modtaboff      uint32
	nmodtab        uint32
	extrefsymoff   uint32
	nextrefsyms    uint32
	indirectsymoff uint32
	nindirectsyms  uint32
	extreloff      uint32
	nextrel        uint32
	locreloff      uint32
	nlocrel        uint32
}

func (l LcDySymTab) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, l.header.Bytes()...)
	for i := 0; i < 18*4; i++ {
		bytes = append(bytes, 0)
	}
	offset := 8
	binary.LittleEndian.PutUint32(bytes[offset:], l.ilocalsym)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nlocalsym)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.iextdefsym)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nextdefsym)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.iundefsym)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nundefsym)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.tocoff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.ntoc)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.modtaboff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nmodtab)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.extrefsymoff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nextrefsyms)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.indirectsymoff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nindirectsyms)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.extreloff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nextrel)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.locreloff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nlocrel)

	//p := len(bytes) % 32
	//// Pad binary if size isn't a multiple of 32
	//for ; p < 32; p++ {
	//	bytes = append(bytes, 0)
	//}
	return bytes
}

func (l LcDySymTab) SizeOf() uint32 {
	return uint32(len(l.Bytes()))
}

func NewLcDySymTab(ilocalsym, nlocalsym, iextdefsym, nextdefsym, iundefsym, nundefsym, tocoff, ntoc, modtaboff, nmodtab, extrefsymoff, nextrefsyms, indirectsymoff, nindirectsyms, extreloff, nextrel, locreloff, nlocrel uint32) LcDySymTab {
	lcDySymtab := LcDySymTab{
		header: LcHeader{
			Cmd: DySymTabCommand,
		},
		ilocalsym:      ilocalsym,
		nlocalsym:      nlocalsym,
		iextdefsym:     iextdefsym,
		nextdefsym:     nextdefsym,
		iundefsym:      iundefsym,
		nundefsym:      nundefsym,
		tocoff:         tocoff,
		ntoc:           ntoc,
		modtaboff:      modtaboff,
		nmodtab:        nmodtab,
		extrefsymoff:   extrefsymoff,
		nextrefsyms:    nextrefsyms,
		indirectsymoff: indirectsymoff,
		nindirectsyms:  nindirectsyms,
		extreloff:      extreloff,
		nextrel:        nextrel,
		locreloff:      locreloff,
		nlocrel:        nlocrel,
	}
	lcDySymtab.header.CmdSize = lcDySymtab.SizeOf()
	return lcDySymtab
}
