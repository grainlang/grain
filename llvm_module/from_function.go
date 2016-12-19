package llvm_module

import (
	"github.com/grainlang/grain/ast"
	"llvm.org/llvm/bindings/go/llvm"
)

func CreateLlvmModuleFromFunction(function ast.Function, allFunctions []ast.Function) llvm.Module {
	context := llvm.GlobalContext()
	builder := context.NewBuilder()
	module := context.NewModule(function.Id)
	llvmFunction := createFunctionDeclarationInModule(function, module)
	bodyBlock := llvm.AddBasicBlock(llvmFunction, "body")
	builder.SetInsertPoint(bodyBlock, bodyBlock.FirstInstruction())
	body := function.Body
	switch typedBody := body.(type) {
	case ast.NativeFunctionCall:
		nativeFunctionParamTypes := make([]llvm.Type, len(typedBody.Parameters))
		for i := range typedBody.Parameters {
			nativeFunctionParamTypes[i] = llvm.Int32Type()
		}
		var nativeFunctionReturnType llvm.Type
		if typedBody.ReturnValue == ast.NativeValueVoid {
			nativeFunctionReturnType = llvm.VoidType()
		} else if typedBody.ReturnValue == ast.NativeValueInt {
			nativeFunctionReturnType = llvm.Int32Type()
		} else {
			panic("Unknown type")
		}
		nativeFunctionType := llvm.FunctionType(nativeFunctionReturnType, nativeFunctionParamTypes, false)
		nativeFunction := llvm.AddFunction(module, typedBody.Name, nativeFunctionType)
		nativeFunctionParamValues := make([]llvm.Value, len(typedBody.Parameters))
		for i, nativeFunctionParam := range typedBody.Parameters {
			for index, param := range function.Parameters {
				if nativeFunctionParam.Id == param.Id {
					nativeFunctionParamValues[i] = llvmFunction.Param(index)
					break
				}
			}
		}
		nativeFunctionReturnValue := builder.CreateCall(nativeFunction, nativeFunctionParamValues, "ret")
		builder.CreateRet(nativeFunctionReturnValue)
	case ast.Binding:
		var providingFunction ast.Function
		var consumingFunction ast.Function
		for _, fn := range allFunctions {
			for _, returnValue := range fn.ReturnValues {
				if typedBody.From == returnValue.Id {
					providingFunction = fn
					break
				}
			}
			for _, param := range fn.Parameters {
				if typedBody.To == param.Id {
					consumingFunction = fn
					break
				}
			}
		}
		providingLlvmFunction := createFunctionDeclarationInModule(providingFunction, module)
		consumingLlvmFunction := createFunctionDeclarationInModule(consumingFunction, module)
		providingFunctionReturnValue := builder.CreateCall(providingLlvmFunction, []llvm.Value{}, "ret")
		consumingFunctionReturnValue := builder.CreateCall(consumingLlvmFunction, []llvm.Value{providingFunctionReturnValue}, "ret")
		builder.CreateRet(consumingFunctionReturnValue)
	}
	return module
}

func createFunctionDeclarationInModule(function ast.Function, module llvm.Module) llvm.Value {
	paramTypes := make([]llvm.Type, len(function.Parameters))
	for i := range function.Parameters {
		paramTypes[i] = llvm.Int32Type()
	}
	returnTypes := make([]llvm.Type, len(function.ReturnValues))
	for i := range function.ReturnValues {
		returnTypes[i] = llvm.Int32Type()
	}
	returnType := llvm.StructType(returnTypes, false)
	llvmFunctionType := llvm.FunctionType(returnType, paramTypes, false)
	llvmFunction := llvm.AddFunction(module, "$" + function.Id, llvmFunctionType)
	return llvmFunction
}
