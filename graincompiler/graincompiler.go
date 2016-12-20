package main

import (
	"llvm.org/llvm/bindings/go/llvm"
	"fmt"
	"os"
	"io/ioutil"
	"os/exec"
	"github.com/grainlang/grain/ast"
	"github.com/grainlang/grain/standard_library"
	"github.com/grainlang/grain/hello"
	"github.com/grainlang/grain/llvm_module"
	"strconv"
)

func main() {
	machine, err := initMachine()
	if err != nil {
		os.Exit(-1)
	}
	allFunctions := []ast.Function{
		standard_library.CreateGetCharacterAst(),
		standard_library.CreatePutCharacterAst(),
		standard_library.CreateToUppercaseAst(),
		hello.CreateGetCharPutCharAst(),
		hello.CreateGetCharToUpperPutCharAst(),
	}
	allModules := []llvm.Module{
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[0], allFunctions),
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[1], allFunctions),
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[2], allFunctions),
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[3], allFunctions),
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[4], allFunctions),
		llvm_module.CreateMainModuleWithCallToFunction(allFunctions[3]),
		llvm_module.CreateMainModuleWithCallToFunction(allFunctions[4]),
	}
	for i, module := range allModules {
		module.Dump()
		buffer, err := machine.EmitToMemoryBuffer(module, llvm.ObjectFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot emit object file to memory buffer:")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		objectFileName := "grainlang" + strconv.Itoa(i) + ".o"
		fmt.Println(objectFileName)
		err = ioutil.WriteFile(objectFileName, buffer.Bytes(), 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot save file:")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	}
	cmd := exec.Command("clang", "grainlang0.o", "grainlang1.o", "grainlang3.o", "grainlang5.o", "-o", "get_put")
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot run clang:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	cmd = exec.Command("clang", "grainlang0.o", "grainlang1.o", "grainlang2.o", "grainlang4.o", "grainlang6.o", "-o", "get_upper_put")
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot run clang:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	os.Remove("grainlang0.o")
	os.Remove("grainlang1.o")
	os.Remove("grainlang2.o")
	os.Remove("grainlang3.o")
	os.Remove("grainlang4.o")
	os.Remove("grainlang5.o")
	os.Remove("grainlang6.o")
}
