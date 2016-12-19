package standard_library

import "github.com/grainlang/grain/ast"

func CreateGetCharacterAst() ast.Function {
	return ast.Function{
		Id: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
		Name: "get character",
		Description: "get character from standard input",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
				ValueType: ast.Character,
			},
		},
		Body: []ast.Expression{
			ast.NativeFunctionCall{
				Name: "getchar",
				Parameters: []ast.Parameter{},
				ReturnValue: ast.NativeValueInt,
			},
		},
	}
}
