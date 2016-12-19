package main

import (
	"llvm.org/llvm/bindings/go/llvm"
	"fmt"
	"os"
	"io/ioutil"
	"os/exec"
)

func main() {

	mainModule := createHelloGrainlangModule()

	machine, err := initMachine()
	if err != nil {
		os.Exit(-1)
	}

	buffer, err := machine.EmitToMemoryBuffer(mainModule, llvm.ObjectFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot emit object file to memory buffer:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	objectFileName := "hello.o"
	err = ioutil.WriteFile(objectFileName, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot save file:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	cmd := exec.Command("clang", objectFileName, "-o", "hello_program")
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot run clang:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	os.Remove(objectFileName)
}
