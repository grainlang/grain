package hello

import "github.com/grainlang/grain/ast"

func CreateGetCharPutCharAst() ast.Function {
	return ast.Function{
		Id: "e34d79c1-134c-4339-a3f6-143c9944ab93",
		Description: "prints out first character from standard input",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.FunctionUse{
				Id: "2a2ecdba-e6f6-40d8-b7ee-7ce8014e65fe",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "908ee06f-a75a-4b3c-86b2-c30271f8ea08",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromFunctionUseId: "2a2ecdba-e6f6-40d8-b7ee-7ce8014e65fe",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
		},
	}
}
