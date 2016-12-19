package main

import "llvm.org/llvm/bindings/go/llvm"

func createHelloGrainlangModule() llvm.Module {
	context := llvm.GlobalContext()
	builder := context.NewBuilder()
	mainModule := context.NewModule("mainModule")

	putsFuncType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, false)
	putsFunc := llvm.AddFunction(mainModule, "puts", putsFuncType)

	printfFuncType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, true)
	printfFunc := llvm.AddFunction(mainModule, "printf", printfFuncType)

	mainFuncType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	mainFunc := llvm.AddFunction(mainModule, "main", mainFuncType)
	body := llvm.AddBasicBlock(mainFunc, "entry")
	builder.SetInsertPoint(body, body.FirstInstruction())

	getcharFuncType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	getcharFunc := llvm.AddFunction(mainModule, "getchar", getcharFuncType)
	putcharFuncType := llvm.FunctionType(llvm.VoidType(), []llvm.Type{llvm.Int32Type()}, false)
	putcharFunc := llvm.AddFunction(mainModule, "putchar", putcharFuncType)

	hello := builder.CreateGlobalStringPtr("Hello, Grainlang!", "hello")
	format := builder.CreateGlobalStringPtr("[%s]", "format")
	format2 := builder.CreateGlobalStringPtr("[%c]\n", "format2")
	builder.CreateCall(putsFunc, []llvm.Value{hello}, "res")
	builder.CreateCall(printfFunc, []llvm.Value{format, hello}, "res")

	char := builder.CreateCall(getcharFunc, []llvm.Value{}, "char")
	builder.CreateCall(putcharFunc, []llvm.Value{char}, "")
	char = builder.CreateCall(getcharFunc, []llvm.Value{}, "char")
	builder.CreateCall(putcharFunc, []llvm.Value{char}, "")
	char = builder.CreateCall(getcharFunc, []llvm.Value{}, "char")
	builder.CreateCall(putcharFunc, []llvm.Value{char}, "")
	builder.CreateCall(printfFunc, []llvm.Value{format2, char}, "res")

	builder.CreateRet(llvm.ConstInt(llvm.Int32Type(), 0, true))

	mainModule.Dump()

	return mainModule
}
