package standard_library

import "github.com/grainlang/grain/ast"

func CreateSubtractAst() ast.Function {
	return ast.Function{
		Id: "c98cad91-869c-4cae-b8ff-c12fcc791dd0",
		Name: "subtract",
		Description: "subtract two integers",
		Parameters: []ast.Parameter{
			{
				Id: "d142b061-7aff-4c39-a3ad-ff0ff0ebf3de",
				Name: "minuend",
				ValueType: ast.Integer,
			},
			{
				Id: "301ee05c-3b79-4c61-a0be-482966161abb",
				Name: "subtrahend",
				ValueType: ast.Integer,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "dc49938b-bdd3-4e69-bbae-c81b0132ef58",
				ValueType: ast.Integer,
			},
		},
		Body: []ast.Expression{
			ast.BinaryOperationCall{
				Name: "-",
				LeftParameter: ast.Parameter{
					Id: "d142b061-7aff-4c39-a3ad-ff0ff0ebf3de",
				},
				RightParameter: ast.Parameter{
					Id: "301ee05c-3b79-4c61-a0be-482966161abb",
				},
				ReturnValue: ast.NativeValueInt64,
			},
		},
	}
}
