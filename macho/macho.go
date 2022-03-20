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
			Flags:      0x85,
			Reserved:   0,
		},
	}
}

func (m *MachO) AddCommand(cmd LoadCommand) {
	m.Commands = append(m.Commands, cmd)
	m.Header.NoOfCmds++
	m.Header.SizeOfCmds += cmd.SizeOf()
}

func (m *MachO) AddPaddingUpTo(num int) []byte {
	bytes := make([]byte, 0)
	if num == 0 || num <= len(m.Bytes()) {
		return bytes
	}
	p := len(m.Bytes()) % num
	// Pad binary if size isn't a multiple of num
	for ; p < num; p++ {
		bytes = append(bytes, 0)
	}
	return bytes
}

func (m *MachO) AddData(bytes []byte, padding int) {
	m.Data = append(m.Data, m.AddPaddingUpTo(padding))
	m.Data = append(m.Data, bytes)
}

func (m MachO) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, m.Header.Bytes()...)
	for _, command := range m.Commands {
		bytes = append(bytes, command.Bytes()...)
	}

	for _, data := range m.Data {
		bytes = append(bytes, data...)
	}

	return bytes
}
