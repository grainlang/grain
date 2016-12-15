package main

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	context := llvm.GlobalContext()
	builder := context.NewBuilder()
	mainModule := context.NewModule("mainModule")

	putsFuncType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, false)
	putsFunc := llvm.AddFunction(mainModule, "puts", putsFuncType)

	mainFuncType := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
	mainFunc := llvm.AddFunction(mainModule, "main", mainFuncType)
	//
	body := llvm.AddBasicBlock(mainFunc, "entry")
	builder.SetInsertPoint(body, mainFunc)

	hello := builder.CreateGlobalStringPtr("hello", "hello")
	builder.SetInsertPoint(body, body.FirstInstruction())
	builder.CreateCall(putsFunc, []llvm.Value{hello}, "puts2")
	builder.CreateRetVoid()

	mainModule.Dump()
}
