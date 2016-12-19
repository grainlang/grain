package hello

import "github.com/grainlang/grain/ast"

func CreateGetCharToUpperPutCharAst() ast.Function {
	return ast.Function{
		Id: "8e0b6230-c25c-482f-a2a4-b5ce009d6edd",
		Description: "prints out first character from standard input, after uppercasing it",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.Binding{
				From: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
				To: "3e647713-dabc-48a2-9103-347bf53c9813",
			},
			ast.Binding{
				From: "430e2b1d-089f-48b1-9e47-d85b5b592d1a",
				To: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
			},
		},
	}
}
