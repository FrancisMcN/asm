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

	pageZeroSegment := load_commands.NewSegment64("__PAGEZERO", 0, 4_294_967_296, 0, 0, 0, 0, nil)
	m.AddCommand(pageZeroSegment)

	textSectionHeaders := []load_commands.SectionHeader64{
		load_commands.NewSectionHeader64("__text", "__TEXT", 4_294_967_296+1024, 1024, 1024, 5, 0, 0, 0),
		load_commands.NewSectionHeader64("__cstring", "__TEXT", 4_294_967_296+2048, 1024, 2048, 5, 0, 0, 0),
	}
	textSegment := load_commands.NewSegment64("__TEXT", 4_294_967_296, 2048, 0, 2048, 0x7, 0x5, textSectionHeaders)
	m.AddCommand(textSegment)

	//linkeditSegment := load_commands.NewSegment64("__LINKEDIT", 4_294_967_296 + 2048, 1024, 2048, 1024, 0x1, 0x1, nil)
	//m.AddCommand(linkeditSegment)

	//textSectionHeaders := []load_commands.SectionHeader64{
	//	load_commands.NewSectionHeader64("__text", "__TEXT", 4_294_967_296 + 4096, 1024, 1024, 5, 0, 0, 0),
	//}
	//textSegment := load_commands.NewSegment64("__TEXT", 4_294_967_296, 3072, 0, 3072, 0x7, 0x5, textSectionHeaders)
	//m.AddCommand(textSegment)

	lcUnixThread := load_commands.NewLcUnixThread(load_commands.X86_Thread_State64)
	m.AddCommand(lcUnixThread)

	//dyldInfoOnly := load_commands.NewLcDyldInfoOnly(0,0,0,0,0,0,0,0,0,0)
	//m.AddCommand(dyldInfoOnly)

	//lcLoadDylinker := load_commands.NewLcLoadDynlinker("/usr/lib/dyld")
	//m.AddCommand(lcLoadDylinker)
	//
	//lcLoadDylib := load_commands.NewLcLoadDylib("/usr/lib/libSystem.B.dylib")
	//m.AddCommand(lcLoadDylib)

	dump(m)

	fmt.Printf("%+v\n", m)
}
