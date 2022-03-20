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
		load_commands.NewSectionHeader64("__text", "__TEXT", 4_294_983_571, 37, 16275, 1, 0, 0, 0x80000400),
	}
	textSegment := load_commands.NewSegment64("__TEXT", 4_294_967_296, 16_384, 0, 16_384, 0x5, 0x5, 0, textSectionHeaders)
	m.AddCommand(textSegment)

	dataSectionHeaders := []load_commands.SectionHeader64{
		load_commands.NewSectionHeader64("__data", "__DATA", 4_294_983_680, 12, 16_384, 1, 0, 0, 0),
	}
	dataSegment := load_commands.NewSegment64("__DATA", 4_294_983_680, 16_384, 16_384, 16_384, 0x3, 0x3, 0, dataSectionHeaders)
	m.AddCommand(dataSegment)

	linkeditSegment := load_commands.NewSegment64("__LINKEDIT", 4_295_000_064, 16_384, 32_768, 176, 0x1, 0x1, 0, nil)
	m.AddCommand(linkeditSegment)

	//lcUnixThread := load_commands.NewLcUnixThread(load_commands.X86_Thread_State64)
	//m.AddCommand(lcUnixThread)

	dyldInfoOnly := load_commands.NewLcDyldInfoOnly(0, 0, 0, 0, 0, 0, 0, 0, 32_768, 48)
	m.AddCommand(dyldInfoOnly)

	lcSymTab := load_commands.NewLcSymTab(32_824, 32_888)
	//lcSymTab.StringTable.AddString("hello world!\n")
	lcSymTab.AddSymbol("__mh_execute_header", 0x0F, 0x01, 0x0010, 4_294_967_296)
	lcSymTab.AddSymbol("_main", 0x0F, 0x01, 0x00, 4_294_983_571)
	lcSymTab.AddSymbol("dyld_stub_binder", 0x01, 0x00, 0x0100, 0)
	lcSymTab.AddSymbol("message", 0x0E, 0x03, 0x00, 4_294_983_680)
	m.AddCommand(lcSymTab)

	lcDySymTab := load_commands.NewLcDySymTab(0, 1, 1, 2, 3, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	m.AddCommand(lcDySymTab)

	lcLoadDylinker := load_commands.NewLcLoadDynlinker("/usr/lib/dyld")
	m.AddCommand(lcLoadDylinker)

	lcMain := load_commands.NewLcMain(16_275, 0)
	m.AddCommand(lcMain)

	lcLoadDylib := load_commands.NewLcLoadDylib("/usr/lib/libSystem.B.dylib")
	m.AddCommand(lcLoadDylib)

	// Strings start at offset 1024
	instructions := []byte{
		0xB8, 0x04, 0x00, 0x00, 0x02, 0xBF, 0x01, 0x00,
		0x00, 0x00, 0x48, 0xBE, 0x00, 0x40, 0x00, 0x00,

		0x01, 0x00, 0x00, 0x00, 0xBA, 0x0D, 0x00, 0x00,
		0x00, 0x0F, 0x05, 0xB8, 0x01, 0x00, 0x00, 0x02,

		0x48, 0x31, 0xFF, 0x0F, 0x05,
	}
	// Start instructions at offset 16_275
	m.AddData(instructions, 16_275)

	// Start DATA section at offset 16_384
	data := []byte{
		0x48, 0x65, 0x6C, 0x6C,
		0x6F, 0x20, 0x57, 0x6F,
		0x72, 0x6C, 0x64, 0x0A,
	}
	m.AddData(data, 16_384)

	// Bytes for dynamic linker, from offset 32_768
	linkerStuff := []byte{
		0x00, 0x01, 0x5F, 0x00, 0x05, 0x00, 0x02, 0x5F,
		0x6D, 0x68, 0x5F, 0x65, 0x78, 0x65, 0x63, 0x75,

		0x74, 0x65, 0x5F, 0x68, 0x65, 0x61, 0x64, 0x65,
		0x72, 0x00, 0x21, 0x6D, 0x61, 0x69, 0x6E, 0x00,

		0x25, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x93,
		0x7F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	m.AddData(linkerStuff, 32_768)

	// Start Symbol table data at 32_824
	m.AddData(lcSymTab.SymbolTableBytes(), 32_824)

	// Start String table at 32_888
	m.AddData(lcSymTab.StringTable.Bytes(), 32_888)
	//Pad up to 33KB // 33_792
	//m.AddData([]byte{}, 32_944)
	//m.AddData([]byte{}, 33_792)
	dump(m)

	fmt.Printf("%+v\n", m)
}
