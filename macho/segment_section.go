package macho

//
//import "francis.ie/asm/macho/load_commands"
//
//type SegmentSectionHeaders64 struct {
//	Segment        load_commands.Segment64
//	SectionHeaders []load_commands.SectionHeader64
//}
//
//func (s SegmentSectionHeaders64) Bytes() []byte {
//	bytes := make([]byte, 0)
//	bytes = append(bytes, s.Segment.Bytes()...)
//	for _, sh := range s.SectionHeaders {
//		bytes = append(bytes, sh.Bytes()...)
//	}
//	return bytes
//}
//
//func (s SegmentSectionHeaders64) SizeOf() uint32 {
//
//	size := s.Segment.SizeOf()
//
//	for _, h := range s.SectionHeaders {
//		size += h.SizeOf()
//	}
//
//	return size
//}
