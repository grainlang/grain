package standard_library

import "github.com/grainlang/grain/ast"

func CreateAddAst() ast.Function {
	return ast.Function{
		Id: "96f1af5d-9d94-428e-a2f5-fc3c63111f27",
		Name: "add",
		Description: "add two integers",
		Parameters: []ast.Parameter{
			{
				Id: "1a6f49fa-5273-4c00-84c5-0278ce1f7041",
				ValueType: ast.Integer,
			},
			{
				Id: "b08c9d64-04a7-4fe2-b9fa-9febb59f5689",
				ValueType: ast.Integer,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "d010a324-fe47-40df-bdb3-15c073d46c86",
				ValueType: ast.Integer,
			},
		},
		Body: []ast.Expression{
			ast.BinaryOperationCall{
				Id: "+id",
				Name: "+",
				LeftParameter: ast.Parameter{
					Id: "1a6f49fa-5273-4c00-84c5-0278ce1f7041",
				},
				RightParameter: ast.Parameter{
					Id: "b08c9d64-04a7-4fe2-b9fa-9febb59f5689",
				},
				ReturnId: "+return",
			},
			ast.Binding{
				FromId: "+id",
				FromReturnValue: "+return",
				ToReturnValue: "d010a324-fe47-40df-bdb3-15c073d46c86",
			},
		},
	}
}
