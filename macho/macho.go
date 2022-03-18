package macho

type MachO struct {
	Header   Header
	Commands []LoadCommand
	Data     [][]byte
}

func NewMacho() MachO {
	return MachO{
		Header: Header{
			Magic:      0xfeed_facf,
			CpuType:    0x0100_0007, // CPU_TYPE_X86_64
			CpuSubType: 0x0000_0003, // CPU_SUBTYPE_X86_64_ALL
			FileType:   MhExecute,
			NoOfCmds:   0,
			SizeOfCmds: 0,
			Flags:      0,
			Reserved:   0,
		},
	}
}

func (m *MachO) AddCommand(cmd LoadCommand) {
	m.Commands = append(m.Commands, cmd)
	m.Header.NoOfCmds++
	m.Header.SizeOfCmds += cmd.SizeOf()
}

func (m MachO) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, m.Header.Bytes()...)
	for _, command := range m.Commands {
		bytes = append(bytes, command.Bytes()...)
	}
	p := len(bytes) % 4096
	// Pad binary if size isn't a multiple of 4096
	for ; p < 4096; p++ {
		bytes = append(bytes, 0)
	}

	bytes[2048] = 'H'
	bytes[2049] = 'e'
	bytes[2050] = 'l'
	bytes[2051] = 'l'
	bytes[2052] = 'o'
	bytes[2053] = ' '
	bytes[2054] = 'W'
	bytes[2055] = 'o'
	bytes[2056] = 'r'
	bytes[2057] = 'l'
	bytes[2058] = 'd'
	bytes[2059] = '\n'

	bytes[1024] = 0xBA
	bytes[1025] = 0x0C

	bytes[1029] = 0xB8
	bytes[1030] = 0x04

	bytes[1033] = 0x02
	bytes[1034] = 0x0F
	bytes[1035] = 0x05
	bytes[1036] = 0x48
	bytes[1037] = 0x89
	bytes[1038] = 0xC7
	bytes[1039] = 0xB8
	bytes[1040] = 0x01

	bytes[1043] = 0x02
	bytes[1044] = 0x0F
	bytes[1045] = 0x05

	//bytes[0x168] = 0xBA
	//bytes[0x169] = 0x0B
	//
	//bytes[0x16D] = 0xB8
	//bytes[0x16E] = 0x04
	//
	//bytes[0x171] = 0x02
	//bytes[0x172] = 0x0F
	//bytes[0x173] = 0x05
	//bytes[0x174] = 0x48
	//bytes[0x175] = 0x89
	//bytes[0x176] = 0xC7
	//bytes[0x177] = 0xB8
	//bytes[0x178] = 0x01
	//
	//bytes[0x17B] = 0x02
	//bytes[0x17C] = 0x0F
	//bytes[0x17D] = 0x05

	return bytes
}
