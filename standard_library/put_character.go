package standard_library

import "github.com/grainlang/grain/ast"

func CreatePutCharacterAst() ast.Function {
	return ast.Function{
		Id: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
		Name: "put character",
		Description: "put character into standard output",
		Parameters: []ast.Parameter{
			{
				Id: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
				ValueType: ast.Character,
			},
		},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.NativeFunctionCall{
				Name: "putchar",
				ParameterBindings: []ast.NativeBinding{
					{
						FromParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
				ReturnType: ast.NativeValueVoid,
			},
			ast.Binding{
			},
		},
	}
}
