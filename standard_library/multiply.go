package standard_library

import "github.com/grainlang/grain/ast"

func CreateMultiplyAst() ast.Function {
	return ast.Function{
		Id: "b45f802e-39ca-48d5-9368-e9ffcdea0d27",
		Name: "multiply",
		Description: "multiply two integers",
		Parameters: []ast.Parameter{
			{
				Id: "b787a111-3f85-4053-a8c3-e1dae3ce6fe4",
				ValueType: ast.Integer,
			},
			{
				Id: "a0a9a1b3-ef5f-469f-ae78-9dc9aedcf510",
				ValueType: ast.Integer,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "f30fdb67-d967-47e8-8745-bd2f9770bd4f",
				ValueType: ast.Integer,
			},
		},
		Body: []ast.Expression{
			ast.BinaryOperationCall{
				Id: "*id",
				Name: "*",
				LeftParameter: ast.Parameter{
					Id: "b787a111-3f85-4053-a8c3-e1dae3ce6fe4",
				},
				RightParameter: ast.Parameter{
					Id: "a0a9a1b3-ef5f-469f-ae78-9dc9aedcf510",
				},
				ReturnId: "*ret",
			},
			ast.Binding{
				FromId: "*id",
				FromReturnValue: "*ret",
				ToReturnValue: "f30fdb67-d967-47e8-8745-bd2f9770bd4f",
			},
		},
	}
}
