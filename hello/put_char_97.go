package hello

import "github.com/grainlang/grain/ast"

func CreatePutCharConst97Ast() ast.Function {
	return ast.Function{
		Id: "2c2cefed-5b1d-4a63-82a2-f1a275b60bc6",
		Name: "print out a",
		Description: "print out a",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.Constant{
				Id: "6002f7ae-748f-40cc-a625-edfe6973000e",
				Name: "int representation of a",
				Value: "97",
				ValueType: ast.Integer,
			},
			ast.FunctionUse{
				Id: "28761c4f-2341-408f-bdca-4c8fff3362ad",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromConstant: "6002f7ae-748f-40cc-a625-edfe6973000e",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
		},
	}
}
