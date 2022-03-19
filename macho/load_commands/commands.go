package load_commands

const (
	MainCommand          = 0x8000_0028
	UnixThreadCommand    = 0x5
	Segment64Command     = 0x19
	DyldInfoOnlyCommand  = 0x8000_0022
	SymTabCommand        = 0x2
	DySymTabCommand      = 0xB
	LoadDynlinkerCommand = 0xE
	LoadDylibCommand     = 0xC
)
