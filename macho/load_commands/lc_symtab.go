package load_commands

import (
	"encoding/binary"
)

type Nlist struct {
	nstrx  uint32
	ntype  uint8
	nsect  uint8
	ndesc  uint16
	nvalue uint64
}

func NewNlist(nstrx uint32, ntype uint8, nsect uint8, ndesc uint16, nvalue uint64) Nlist {
	return Nlist{
		nstrx:  nstrx,
		ntype:  ntype,
		nsect:  nsect,
		ndesc:  ndesc,
		nvalue: nvalue,
	}
}

func (l Nlist) Bytes() []byte {
	bytes := make([]byte, 4)
	offset := 0
	binary.LittleEndian.PutUint32(bytes[offset:], l.nstrx)
	offset += 4
	bytes = append(bytes, l.ntype)
	offset += 1
	bytes = append(bytes, l.nsect)
	offset += 1
	bytes = append(bytes, 0, 0)
	binary.LittleEndian.PutUint16(bytes[offset:], l.ndesc)
	offset += 2
	bytes = append(bytes, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(bytes[offset:], l.nvalue)

	return bytes
}

type LcSymTab struct {
	header  LcHeader
	symoff  uint32
	nsyms   uint32
	stroff  uint32
	strsize uint32

	symbols     []Nlist
	StringTable StringTable
}

func (l LcSymTab) SymbolTableBytes() []byte {
	bytes := make([]byte, 0)
	for _, s := range l.symbols {
		bytes = append(bytes, s.Bytes()...)
	}
	return bytes
}

func (l LcSymTab) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, l.header.Bytes()...)
	for i := 0; i < 16; i++ {
		bytes = append(bytes, 0)
	}
	offset := 8
	binary.LittleEndian.PutUint32(bytes[offset:], l.symoff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.nsyms)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.stroff)
	offset += 4
	binary.LittleEndian.PutUint32(bytes[offset:], l.strsize)

	return bytes
}

func (l LcSymTab) SizeOf() uint32 {
	return uint32(len(l.Bytes()))
}

func (l *LcSymTab) AddSymbol(symbol string, ntype uint8, nsect uint8, ndesc uint16, addr uint64) {
	l.StringTable.AddString(symbol)
	l.nsyms++
	l.strsize = l.StringTable.SizeOf()
	nstrx := l.StringTable.GetStringOffset(symbol)
	nlist := NewNlist(nstrx, ntype, nsect, ndesc, addr)
	l.symbols = append(l.symbols, nlist)
}

func NewLcSymTab(symoff uint32, stroff uint32) LcSymTab {
	lcSymTab := LcSymTab{
		header: LcHeader{
			Cmd: SymTabCommand,
		},
		symoff:      symoff,
		stroff:      stroff,
		symbols:     make([]Nlist, 0),
		StringTable: NewStringTable(),
	}
	lcSymTab.header.CmdSize = lcSymTab.SizeOf()
	return lcSymTab
}
