package load_commands

type StringTable struct {
	stringMap   map[string]uint32
	stringIndex map[string]uint32
	stringTable []string
}

func (s *StringTable) AddString(str string) {
	if _, found := s.stringMap[str]; !found {
		s.stringMap[str] = uint32(len(s.Bytes()))
		s.stringIndex[str] = uint32(len(s.stringIndex))
		s.stringTable = append(s.stringTable, str)
	}
}

func (s StringTable) HasString(str string) bool {
	if _, found := s.stringMap[str]; found {
		return true
	}
	return false
}

func (s StringTable) GetStringIndex(str string) uint32 {
	if val, found := s.stringIndex[str]; found {
		return val
	}
	return 0
}

func (s StringTable) GetStringOffset(str string) uint32 {
	if val, found := s.stringMap[str]; found {
		return val
	}
	return 0
}

func (s StringTable) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, 0x20, 0x00)
	for _, str := range s.stringTable {
		bytes = append(bytes, []byte(str)...)
		bytes = append(bytes, 0x00)
	}
	p := len(bytes) % 56
	// Pad binary if size isn't a multiple of 32
	for ; p < 56; p++ {
		bytes = append(bytes, 0)
	}
	return bytes
}

func (s StringTable) SizeOf() uint32 {
	//var size uint32 = 0
	//for _, str := range s.stringTable {
	//	size += uint32(len([]byte(str)))
	//}
	return uint32(len(s.Bytes()))
}

func NewStringTable() StringTable {

	return StringTable{
		stringMap:   make(map[string]uint32, 0),
		stringIndex: make(map[string]uint32, 1),
		stringTable: make([]string, 0),
	}

}
