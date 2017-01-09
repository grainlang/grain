package standard_library

import "github.com/grainlang/grain/ast"

func CreateToUppercaseAst() ast.Function {
	return ast.Function{
		Id: "f2705a1f-59b5-4ae6-bee1-bf6d220e2794",
		Name: "to uppercase",
		Description: "convert character to it's uppercase version, e.g. a -> A",
		Parameters: []ast.Parameter{
			{
				Id: "3e647713-dabc-48a2-9103-347bf53c9813",
				ValueType: ast.Character,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "430e2b1d-089f-48b1-9e47-d85b5b592d1a",
				ValueType: ast.Character,
			},
		},
		Body: []ast.Expression{
			ast.NativeFunctionCall{
				Id: "07322ea3-d7c1-49b2-a2e7-114b168830c9",
				Name: "toupper",
				Parameters: []ast.Parameter{
					{
						Id: "3e647713-dabc-48a2-9103-347bf53c9813",
					},
				},
				ReturnType: ast.NativeValueInt,
				ReturnId: "463081ce-8785-4efe-b6c5-4a814a87cd93",
			},
			ast.Binding{
				FromId: "07322ea3-d7c1-49b2-a2e7-114b168830c9",
				FromReturnValue: "463081ce-8785-4efe-b6c5-4a814a87cd93",
				ToReturnValue: "430e2b1d-089f-48b1-9e47-d85b5b592d1a",
			},
		},
	}
}
