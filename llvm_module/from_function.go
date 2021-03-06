package llvm_module

import (
	"github.com/grainlang/grain/ast"
	"llvm.org/llvm/bindings/go/llvm"
	"strconv"
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
	returnBindings := make([]llvm.Value, 0)
	for _, body := range function.Body {
		switch typedBody := body.(type) {
		case ast.Constant:
			var constValue llvm.Value
			if typedBody.ValueType == ast.Character {
				if len(typedBody.Value) != 1 {
					panic("Wrong character length: " + typedBody.Value)
				}
				constValue = llvm.ConstInt(llvm.Int32Type(), uint64(typedBody.Value[0]), true)
			} else if typedBody.ValueType == ast.Integer {
				constValue = llvm.ConstIntFromString(llvm.Int64Type(), typedBody.Value, 10)
			} else {
				panic("Unknown constant type.")
			}
			returnValueToLlvmValue[typedBody.Id] = constValue
		case ast.NativeFunctionCall:
			nativeFunctionParamTypes := make([]llvm.Type, len(typedBody.ParameterBindings))
			for i := range typedBody.ParameterBindings {
				nativeFunctionParamTypes[i] = llvm.Int32Type()
			}
			var nativeFunctionReturnType llvm.Type
			if typedBody.ReturnType == ast.NativeValueVoid {
				nativeFunctionReturnType = llvm.VoidType()
			} else if typedBody.ReturnType == ast.NativeValueInt {
				nativeFunctionReturnType = llvm.Int32Type()
			} else {
				panic("Unknown type")
			}
			nativeFunctionType := llvm.FunctionType(nativeFunctionReturnType, nativeFunctionParamTypes, false)
			nativeFunction := llvm.AddFunction(module, typedBody.Name, nativeFunctionType)
			nativeFunctionParamValues := make([]llvm.Value, len(typedBody.ParameterBindings))
			for i, nativeFunctionParam := range typedBody.ParameterBindings {
				for index, param := range function.Parameters {
					if nativeFunctionParam.FromParameter == param.Id {
						nativeFunctionParamValues[i] = llvmFunction.Param(index)
						break
					}
				}
			}
			nativeFunctionReturnValue := builder.CreateCall(nativeFunction, nativeFunctionParamValues, "ret")
			returnValueToLlvmValue[typedBody.Id + " " + typedBody.ReturnId] = nativeFunctionReturnValue
		case ast.BinaryOperationCall:
			var opcode llvm.Opcode
			if typedBody.Name == "+" {
				opcode = llvm.Add
			} else if typedBody.Name == "-" {
				opcode = llvm.Sub
			} else if typedBody.Name == "*" {
				opcode = llvm.Mul
			} else if typedBody.Name == "/" {
				opcode = llvm.SDiv
			} else if typedBody.Name == "%" {
				opcode = llvm.SRem
			} else if typedBody.Name == "=" {
				opcode = llvm.ICmp
			} else {
				panic("Unknown operator: " + typedBody.Name)
			}
			var leftParam, rightParam llvm.Value
			for index, param := range function.Parameters {
				if typedBody.LeftParameterBinding.FromParameter == param.Id {
					leftParam = llvmFunction.Param(index)
				} else if typedBody.RightParameterBinding.FromParameter == param.Id {
					rightParam = llvmFunction.Param(index)
				}
			}
			var binaryOperationReturnValue llvm.Value
			if typedBody.Name == "=" {
				binaryOperationReturnValue = builder.CreateICmp(llvm.IntEQ, leftParam, rightParam, "ret")
			} else {
				binaryOperationReturnValue = builder.CreateBinOp(opcode, leftParam, rightParam, "ret")
			}
			returnValueToLlvmValue[typedBody.Id + " " + typedBody.ReturnId] = binaryOperationReturnValue
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
				var returnValueId string
				if binding.FromConstant != "" {
					returnValueId = binding.FromConstant
				} else {
					returnValueId = binding.FromId + " " + binding.FromReturnValue
				}
				llvmParams[i] = returnValueToLlvmValue[returnValueId]
			}
			consumingFunctionReturnValue := builder.CreateCall(consumingLlvmFunction, llvmParams, "ret")
			for i, returnValue := range consumingFunction.ReturnValues {
				var name string
				if returnValue.Name != "" {
					name = returnValue.Name
				} else {
					name = "elem" + strconv.Itoa(i)
				}
				element := builder.CreateExtractValue(consumingFunctionReturnValue, i, name)
				returnValueToLlvmValue[typedBody.Id + " " + returnValue.Id] = element
			}
		case ast.Binding:
			returnBindings = append(returnBindings, returnValueToLlvmValue[typedBody.FromId + " " + typedBody.FromReturnValue])
		}
	}
	builder.CreateAggregateRet(returnBindings)
	return module
}

func createFunctionDeclarationInModule(function ast.Function, module llvm.Module) llvm.Value {
	paramTypes := make([]llvm.Type, len(function.Parameters))
	for i, param := range function.Parameters {
		if param.ValueType == ast.Integer {
			paramTypes[i] = llvm.Int64Type()
		} else if param.ValueType == ast.Character {
			paramTypes[i] = llvm.Int32Type()
		} else {
			panic("Unknown param type.")
		}
	}
	returnTypes := make([]llvm.Type, len(function.ReturnValues))
	for i, returnValue := range function.ReturnValues {
		if returnValue.ValueType == ast.Integer {
			returnTypes[i] = llvm.Int64Type()
		} else if returnValue.ValueType == ast.Character {
			returnTypes[i] = llvm.Int32Type()
		} else if returnValue.ValueType == ast.Boolean {
			returnTypes[i] = llvm.Int1Type()
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