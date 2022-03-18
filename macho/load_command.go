package macho

type LoadCommand interface {
	Bytes() []byte
	SizeOf() uint32
}
