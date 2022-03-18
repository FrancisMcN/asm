package load_commands

import "encoding/binary"

type LcDyldInfoOnly struct {
	header                LcHeader
	RebaseInfoOffset      uint32
	RebaseInfoSize        uint32
	BindingInfoOffset     uint32
	BindingInfoSize       uint32
	WeakBindingInfoOffset uint32
	WeakBindingInfoSize   uint32
	LazyBindingInfoOffset uint32
	LazyBindingInfoSize   uint32
	ExportInfoOffset      uint32
	ExportInfoSize        uint32
}

func (l LcDyldInfoOnly) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, l.header.Bytes()...)
	for i := 0; i < 40; i++ {
		bytes = append(bytes, 0)
	}
	i := 8
	binary.LittleEndian.PutUint32(bytes[i:], l.RebaseInfoOffset)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.RebaseInfoSize)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.BindingInfoOffset)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.BindingInfoSize)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.WeakBindingInfoOffset)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.WeakBindingInfoSize)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.LazyBindingInfoOffset)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.LazyBindingInfoSize)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.ExportInfoOffset)
	i += 4
	binary.LittleEndian.PutUint32(bytes[i:], l.ExportInfoSize)

	return bytes
}

func (l LcDyldInfoOnly) SizeOf() uint32 {
	return uint32(len(l.Bytes()))
}

func NewLcDyldInfoOnly(rebaseInfoOffset, rebaseInfoSize, bindingInfoOffset, bindingInfoSize, weakBindingInfoOffset, weakBindingInfoSize, lazyBindingInfoOffset, lazyBindingInfoSize, exportInfoOffset, exportInfoSize uint32) LcDyldInfoOnly {
	lcDyldInfoOnly := LcDyldInfoOnly{
		header: LcHeader{
			Cmd: DyldInfoOnlyCommand,
		},
		RebaseInfoOffset:      rebaseInfoOffset,
		RebaseInfoSize:        rebaseInfoSize,
		BindingInfoOffset:     bindingInfoOffset,
		BindingInfoSize:       bindingInfoSize,
		WeakBindingInfoOffset: weakBindingInfoOffset,
		WeakBindingInfoSize:   weakBindingInfoSize,
		LazyBindingInfoOffset: lazyBindingInfoOffset,
		LazyBindingInfoSize:   lazyBindingInfoSize,
		ExportInfoOffset:      exportInfoOffset,
		ExportInfoSize:        exportInfoSize,
	}
	lcDyldInfoOnly.header.CmdSize = lcDyldInfoOnly.SizeOf()
	return lcDyldInfoOnly
}
