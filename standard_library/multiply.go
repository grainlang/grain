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
				Name: "factor",
				ValueType: ast.Integer,
			},
			{
				Id: "a0a9a1b3-ef5f-469f-ae78-9dc9aedcf510",
				Name: "factor",
				ValueType: ast.Integer,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "f30fdb67-d967-47e8-8745-bd2f9770bd4f",
				Name: "product",
				ValueType: ast.Integer,
			},
		},
		Body: []ast.Expression{
			ast.BinaryOperationCall{
				Id: "ef20c5c2-6f58-4408-af25-2ca90d95e454",
				Name: "*",
				LeftParameter: ast.Parameter{
					Id: "b787a111-3f85-4053-a8c3-e1dae3ce6fe4",
				},
				RightParameter: ast.Parameter{
					Id: "a0a9a1b3-ef5f-469f-ae78-9dc9aedcf510",
				},
				ReturnId: "1acccafb-2ddb-4ef9-8e2d-94a0f7301cfd",
			},
			ast.Binding{
				FromId: "ef20c5c2-6f58-4408-af25-2ca90d95e454",
				FromReturnValue: "1acccafb-2ddb-4ef9-8e2d-94a0f7301cfd",
				ToReturnValue: "f30fdb67-d967-47e8-8745-bd2f9770bd4f",
			},
		},
	}
}
