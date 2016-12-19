package llvm_module

import (
	"github.com/grainlang/grain/ast"
	"llvm.org/llvm/bindings/go/llvm"
)

func CreateMainModuleWithCallToFunction(function ast.Function) llvm.Module {
	context := llvm.GlobalContext()
	builder := context.NewBuilder()
	module := context.NewModule("main with " + function.Id + " " + function.Name)
	mainFunctionType := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
	mainFunction := llvm.AddFunction(module, "main", mainFunctionType)
	bodyBlock := llvm.AddBasicBlock(mainFunction, "body")
	builder.SetInsertPoint(bodyBlock, bodyBlock.FirstInstruction())
	calledFunction := createFunctionDeclarationInModule(function, module)
	builder.CreateCall(calledFunction, []llvm.Value{}, "")
	builder.CreateRetVoid()
	return module
}
