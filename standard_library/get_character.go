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
				Id: "c69a023e-45ed-41fc-b569-a4c197b9f32e",
				Name: "getchar",
				Parameters: []ast.Parameter{},
				ReturnType: ast.NativeValueInt,
				ReturnId: "68cfeb4e-4bad-451b-af77-c717a159e238",
			},
			ast.Binding{
				FromId: "c69a023e-45ed-41fc-b569-a4c197b9f32e",
				FromReturnValue: "68cfeb4e-4bad-451b-af77-c717a159e238",
				ToReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
			},
		},
	}
}
