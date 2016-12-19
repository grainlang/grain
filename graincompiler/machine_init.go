package main

import (
	"llvm.org/llvm/bindings/go/llvm"
	"fmt"
	"os"
)

func initMachine() (machine llvm.TargetMachine, err error) {
	err = llvm.InitializeNativeTarget()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Native target initialization error:")
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = llvm.InitializeNativeAsmPrinter()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ASM printer initialization error:")
		fmt.Fprintln(os.Stderr, err)
		return
	}
	target, err := llvm.GetTargetFromTriple(llvm.DefaultTargetTriple())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get target:")
		fmt.Fprintln(os.Stderr, err)
		return
	}
	machine = target.CreateTargetMachine(
		llvm.DefaultTargetTriple(),
		"",
		"",
		llvm.CodeGenLevelNone,
		llvm.RelocDefault,
		llvm.CodeModelSmall)
	return
}
