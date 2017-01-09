package hello

import "github.com/grainlang/grain/ast"

func CreateGetCharToUpperPutCharAst() ast.Function {
	return ast.Function{
		Id: "8e0b6230-c25c-482f-a2a4-b5ce009d6edd",
		Name: "prints out first character from standard input, after uppercasing it",
		Description: "prints out first character from standard input, after uppercasing it",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.FunctionUse{
				Id: "850865e4-167d-4742-85b8-08fa276ded53",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "e01ba0b0-4556-40d3-afc2-6a2fe1672b9c",
				FunctionId: "f2705a1f-59b5-4ae6-bee1-bf6d220e2794",
				Bindings: []ast.Binding{
					{
						FromId: "850865e4-167d-4742-85b8-08fa276ded53",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "3e647713-dabc-48a2-9103-347bf53c9813",
					},
				},
			},
			ast.FunctionUse{
				Id: "10bc2496-30d4-4063-b3e8-a0e7448b5aab",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromId: "e01ba0b0-4556-40d3-afc2-6a2fe1672b9c",
						FromReturnValue: "430e2b1d-089f-48b1-9e47-d85b5b592d1a",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
		},
	}
}
