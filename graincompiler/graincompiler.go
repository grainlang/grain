package main

import (
	"llvm.org/llvm/bindings/go/llvm"
	"fmt"
	"os"
	"io/ioutil"
	"os/exec"
)

func main() {
	context := llvm.GlobalContext()
	builder := context.NewBuilder()
	mainModule := context.NewModule("mainModule")

	putsFuncType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, false)
	putsFunc := llvm.AddFunction(mainModule, "puts", putsFuncType)

	printfFuncType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, true)
	printfFunc := llvm.AddFunction(mainModule, "printf", printfFuncType)

	mainFuncType := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
	mainFunc := llvm.AddFunction(mainModule, "main", mainFuncType)
	body := llvm.AddBasicBlock(mainFunc, "entry")
	builder.SetInsertPoint(body, body.FirstInstruction())

	hello := builder.CreateGlobalStringPtr("Hello, Grainlang!", "hello")
	format := builder.CreateGlobalStringPtr("[%s]", "format")
	builder.CreateCall(putsFunc, []llvm.Value{hello}, "res")
	builder.CreateCall(printfFunc, []llvm.Value{format, hello}, "res")
	builder.CreateRetVoid()

	mainModule.Dump()

	var err error
	var target llvm.Target

	llvm.LinkInMCJIT()

	err = llvm.InitializeNativeTarget()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Native target initialization error:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	err = llvm.InitializeNativeAsmPrinter()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ASM printer initialization error:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	target, err = llvm.GetTargetFromTriple(llvm.DefaultTargetTriple())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get target:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	fmt.Println("Initialize: TargetTriple = " + llvm.DefaultTargetTriple())
	fmt.Println("Initialize: TargetDescription = " + target.Description())

	machine := target.CreateTargetMachine(llvm.DefaultTargetTriple(),
		"", "",
		llvm.CodeGenLevelNone,
		llvm.RelocDefault,
		llvm.CodeModelSmall)
	buffer, err := machine.EmitToMemoryBuffer(mainModule, llvm.ObjectFile)
	objectFileName := "hello.o"
	ioutil.WriteFile(objectFileName, buffer.Bytes(), 0644)
	cmd := exec.Command("clang", objectFileName, "-o", "hello_program")
	cmd.Run()
	os.Remove(objectFileName)
}
