package main

import (
	"encoding/binary"
	"fmt"
	"francis.ie/asm/macho"
	"francis.ie/asm/macho/load_commands"
	"log"
	"os"
)

func dump(m macho.MachO) {
	f, err := os.Create("test-asm-bin")
	if err != nil {
		log.Fatal("Couldn't open file")
	}
	defer f.Close()

	err = binary.Write(f, binary.LittleEndian, m.Bytes())
	if err != nil {
		log.Fatal("Write failed")
	}
}

func main() {

	m := macho.NewMacho()

	pageZeroSegment := load_commands.NewSegment64("__PAGEZERO", 0, 4_294_967_296, 0, 0, 0, 0, 0, nil)
	m.AddCommand(pageZeroSegment)

	textSectionHeaders := []load_commands.SectionHeader64{
		load_commands.NewSectionHeader64("__text", "__TEXT", 4_294_967_296, 64, 1024, 5, 0, 0, 0x80000400),
	}
	textSegment := load_commands.NewSegment64("__TEXT", 4_294_967_296, 16_384, 0, 16_384, 0x7, 0x5, 0, textSectionHeaders)
	m.AddCommand(textSegment)

	dataSectionHeaders := []load_commands.SectionHeader64{
		load_commands.NewSectionHeader64("__data", "__DATA", 4_294_967_296+16_384, 1024, 4096, 1, 0, 0, 0),
	}
	dataSegment := load_commands.NewSegment64("__DATA", 4_294_967_296+16_384, 4096, 4096, 4096, 0x7, 0x5, 0, dataSectionHeaders)
	m.AddCommand(dataSegment)

	linkeditSegment := load_commands.NewSegment64("__LINKEDIT", 4_294_967_296+20_480, 4096, 12_288, 4096, 0x1, 0x1, 0, nil)
	m.AddCommand(linkeditSegment)

	lcUnixThread := load_commands.NewLcUnixThread(load_commands.X86_Thread_State64)
	m.AddCommand(lcUnixThread)

	dyldInfoOnly := load_commands.NewLcDyldInfoOnly(0, 0, 0, 0, 0, 0, 0, 0, 4096, 48)
	m.AddCommand(dyldInfoOnly)

	lcSymTab := load_commands.NewLcSymTab(8192)
	lcSymTab.StringTable.AddString("hello world!\n")
	lcSymTab.AddSymbol("_main", 4_294_967_296 + 1024)
	lcSymTab.AddSymbol("test_symbol", 4_294_967_296 + 1024 + 64)
	m.AddCommand(lcSymTab)

	lcDySymTab := load_commands.NewLcDySymTab(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	m.AddCommand(lcDySymTab)

	lcLoadDylinker := load_commands.NewLcLoadDynlinker("/usr/lib/dyld")
	m.AddCommand(lcLoadDylinker)

	lcLoadDylib := load_commands.NewLcLoadDylib("/usr/lib/libSystem.B.dylib")
	m.AddCommand(lcLoadDylib)

	// Strings start at offset 1024
	instructions := []byte{
		0xBA, 0x0D, 0x00, 0x00,
		0x00, 0xB8, 0x04, 0x00,
		0x00, 0x02, 0x0F, 0x05,
		0x48, 0x89, 0xC7, 0xB8,
		0x01, 0x00, 0x00, 0x02,
		0x0F, 0x05,
	}
	// Start instructions at offset 1024
	m.AddData(instructions, 1024)

	// Start DATA section at offset 2048
	m.AddData([]byte{}, 2048)

	//// Start string table at offset 3072
	//m.AddData([]byte("hello world!\n"), 3072)

	// Leave 4096 bytes for dynamic linker, from offset 4096 to offset 8192
	linkerStuff := []byte{
		0x00,
		//0x00, 0x01, 0x5F, 0x00,
		//0x05, 0x00, 0x02, 0x5F,
		//0x6D, 0x68, 0x5F, 0x65,
		//0x78, 0x65, 0x63, 0x75,
		//0x74, 0x65, 0x5F, 0x68,
		//0x65, 0x61, 0x64, 0x65,
		//0x72, 0x00, 0x21, 0x6D,
		//0x61, 0x69, 0x6E, 0x00,
		//0x25, 0x02, 0x00, 0x00,
		//0x00, 0x03, 0x00, 0x94,
		//0x7F, 0x00, 0x00, 0x00,
		//0x00, 0x00, 0x00, 0x00,
		//0x94, 0x7F, 0x00, 0x00,
		//0x00, 0x00, 0x00, 0x00,
		//0x2D, 0x00, 0x00, 0x00,
		//0x0E, 0x03, 0x00, 0x00,
		//0x00, 0x40, 0x00, 0x00,
		//0x01, 0x00, 0x00, 0x00,
		//0x02, 0x00, 0x00, 0x00,
		//0x03, 0x01, 0x10, 0x00,
		//0x00, 0x00, 0x00, 0x00,
		//0x01, 0x00, 0x00, 0x00,
		//0x16, 0x00, 0x00, 0x00,
		//0x0F, 0x01, 0x00, 0x00,
		//0x94, 0x3F, 0x00, 0x00,
		//0x01, 0x00, 0x00, 0x00,
		//0x1C, 0x00, 0x00, 0x00,
		//0x01, 0x00, 0x00, 0x01,
		//0x00, 0x00, 0x00, 0x00,
		//0x00, 0x00, 0x00, 0x00,
	}
	m.AddData(linkerStuff, 4096)

	// Start Symbol table data at 8192
	m.AddData(lcSymTab.SymbolTableBytes(), 8192)

	// Start String table at 1024
	m.AddData(lcSymTab.StringTable.Bytes(), 10240)

	// Pad up to 12KB
	m.AddData([]byte{}, 16_384)

	dump(m)

	fmt.Printf("%+v\n", m)
}
