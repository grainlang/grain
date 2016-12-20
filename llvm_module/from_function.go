package llvm_module

import (
	"github.com/grainlang/grain/ast"
	"llvm.org/llvm/bindings/go/llvm"
)

func CreateLlvmModuleFromFunction(function ast.Function, allFunctions []ast.Function) llvm.Module {
	context := llvm.GlobalContext()
	builder := context.NewBuilder()
	module := context.NewModule(function.Id + " " + function.Name)
	llvmFunction := createFunctionDeclarationInModule(function, module)
	bodyBlock := llvm.AddBasicBlock(llvmFunction, "body")
	builder.SetInsertPoint(bodyBlock, bodyBlock.FirstInstruction())
	returnValueToLlvmValue := make(map[string]llvm.Value)
	functionToLlvmDeclaration := make(map[string]llvm.Value)
	for expressionIndex, body := range function.Body {
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
		case ast.BinaryOperationCall:
			var opcode llvm.Opcode
			if typedBody.Name == "+" {
				opcode = llvm.Add
			} else if typedBody.Name == "-" {
				opcode = llvm.Sub
			} else if typedBody.Name == "*" {
				opcode = llvm.Mul
			} else {
				panic("Unknown operator: " + typedBody.Name)
			}
			var leftParam, rightParam llvm.Value
			for index, param := range function.Parameters {
				if typedBody.LeftParameter.Id == param.Id {
					leftParam = llvmFunction.Param(index)
				} else if typedBody.RightParameter.Id == param.Id {
					rightParam = llvmFunction.Param(index)
				}
			}
			binaryOperationReturnValue := builder.CreateBinOp(opcode, leftParam, rightParam, "ret")
			builder.CreateRet(binaryOperationReturnValue)
		case ast.FunctionUse:
			consumingFunction := FindUsedFunction(typedBody, allFunctions)
			var consumingLlvmFunction llvm.Value
			if foundDeclaration, ok := functionToLlvmDeclaration[consumingFunction.Id]; ok {
				consumingLlvmFunction = foundDeclaration
			} else {
				consumingLlvmFunction = createFunctionDeclarationInModule(consumingFunction, module)
				functionToLlvmDeclaration[consumingFunction.Id] = consumingLlvmFunction
			}
			llvmParams := make([]llvm.Value, len(typedBody.Bindings))
			for i, binding := range typedBody.Bindings {
				llvmParams[i] = returnValueToLlvmValue[binding.FromFunctionUseId + " " + binding.FromReturnValue]
			}
			consumingFunctionReturnValue := builder.CreateCall(consumingLlvmFunction, llvmParams, "ret")
			if expressionIndex == len(function.Body) - 1 {
				builder.CreateRet(consumingFunctionReturnValue)
			} else {
				if len(consumingFunction.ReturnValues) > 0 {
					returnValueToLlvmValue[typedBody.Id + " " + consumingFunction.ReturnValues[0].Id] = consumingFunctionReturnValue
				}
			}
		}
	}
	return module
}

func createFunctionDeclarationInModule(function ast.Function, module llvm.Module) llvm.Value {
	paramTypes := make([]llvm.Type, len(function.Parameters))
	for i, param := range function.Parameters {
		if param.ValueType == ast.Integer {
			paramTypes[i] = llvm.Int32Type()
		} else {
			paramTypes[i] = llvm.Int32Type()
		}
	}
	returnTypes := make([]llvm.Type, len(function.ReturnValues))
	for i, returnValue := range function.ReturnValues {
		if returnValue.ValueType == ast.Integer {
			returnTypes[i] = llvm.Int32Type()
		} else {
			returnTypes[i] = llvm.Int32Type()
		}
	}
	returnType := llvm.StructType(returnTypes, false)
	llvmFunctionType := llvm.FunctionType(returnType, paramTypes, false)
	llvmFunction := llvm.AddFunction(module, "$" + function.Id, llvmFunctionType)
	return llvmFunction
}

func FindUsedFunction(bodyPart ast.FunctionUse, allFunctions []ast.Function) ast.Function {
	for _, fn := range allFunctions {
		if bodyPart.FunctionId == fn.Id {
			return fn
		}
	}
	panic("No such function " + bodyPart.FunctionId)
}