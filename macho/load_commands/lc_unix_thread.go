package load_commands

import "encoding/binary"

const (
	X86_Thread_State64 = 4
)

type X86CpuThreadState struct {
	Rax    uint64
	Rbx    uint64
	Rcx    uint64
	Rdx    uint64
	Rdi    uint64
	Rsi    uint64
	Rbp    uint64
	Rsp    uint64
	R8     uint64
	R9     uint64
	R10    uint64
	R11    uint64
	R12    uint64
	R13    uint64
	R14    uint64
	R15    uint64
	Rip    uint64
	Rflags uint64
	Cs     uint64
	Gs     uint64
	Fs     uint64
}

func (s *X86CpuThreadState) Count() uint32 {
	// Number of registers is 21
	// 'Count' must be expressed as a number of 32-bit integers
	// X86_64 registers are 64 bit so count = 21 * 2
	return 42
}

func (s *X86CpuThreadState) Bytes() []byte {
	bytes := make([]byte, 4*s.Count())
	binary.LittleEndian.PutUint64(bytes, s.Rax)
	i := 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rbx)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rcx)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rdx)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rdi)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rsi)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rbp)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rsp)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R8)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R9)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R10)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R11)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R12)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R13)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R14)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.R15)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rip)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Rflags)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Cs)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Gs)
	i += 8
	binary.LittleEndian.PutUint64(bytes[i:], s.Fs)

	return bytes
}

type LcUnixThread struct {
	header LcHeader
	Flavor uint32
	Count  uint32
	State  X86CpuThreadState
}

func (l LcUnixThread) Bytes() []byte {
	bytes := make([]byte, 0)
	bytes = append(bytes, l.header.Bytes()...)
	for i := 0; i < 8; i++ {
		bytes = append(bytes, 0)
	}
	binary.LittleEndian.PutUint32(bytes[8:], l.Flavor)
	binary.LittleEndian.PutUint32(bytes[12:], l.Count)
	bytes = append(bytes, l.State.Bytes()...)
	return bytes
}

func (l LcUnixThread) SizeOf() uint32 {
	return uint32(len(l.Bytes()))
}

func NewLcUnixThread(flav uint32) LcUnixThread {
	threadState := X86CpuThreadState{
		//Rdi: 1,
		//Rsi: 4_294_967_336,
		//Rip: 4_294_967_656,
		Rdi: 1,
		Rsi: 4_294_967_296 + 10240,
		Rip: 4_294_967_296 + 1024,
		// 4_507_069_039
		// 212_101_743
	}
	lcUnixThread := LcUnixThread{
		header: LcHeader{
			Cmd: UnixThreadCommand,
		},
		Flavor: flav,
		Count:  threadState.Count(),
		State:  threadState,
	}
	lcUnixThread.header.CmdSize = lcUnixThread.SizeOf()
	return lcUnixThread
}
