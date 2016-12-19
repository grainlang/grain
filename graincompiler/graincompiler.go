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

	allFunctions := []ast.Function{
		standard_library.CreateGetCharacterAst(),
		standard_library.CreatePutCharacterAst(),
		hello.CreateGetCharPutCharAst(),
	}
	allModules := []llvm.Module{
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[0], allFunctions),
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[1], allFunctions),
		llvm_module.CreateLlvmModuleFromFunction(allFunctions[2], allFunctions),
		llvm_module.CreateMainModuleWithCallToFunction(allFunctions[2]),
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
	cmd = exec.Command("clang", "grainlang0.o", "grainlang1.o", "grainlang2.o", "grainlang3.o", "-o", "get_put")
	err = cmd.Run()
	os.Remove("grainlang0.o")
	os.Remove("grainlang1.o")
	os.Remove("grainlang2.o")
	os.Remove("grainlang3.o")
}
