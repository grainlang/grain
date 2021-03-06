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
				Name: "addend",
				ValueType: ast.Integer,
			},
			{
				Id: "b08c9d64-04a7-4fe2-b9fa-9febb59f5689",
				Name: "addend",
				ValueType: ast.Integer,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "d010a324-fe47-40df-bdb3-15c073d46c86",
				Name: "sum",
				ValueType: ast.Integer,
			},
		},
		Body: []ast.Expression{
			ast.BinaryOperationCall{
				Id: "a4d648a1-29b9-457e-af95-77965a61c4a8",
				Name: "+",
				LeftParameterBinding: ast.NativeBinding{
					FromParameter: "1a6f49fa-5273-4c00-84c5-0278ce1f7041",
				},
				RightParameterBinding: ast.NativeBinding{
					FromParameter: "b08c9d64-04a7-4fe2-b9fa-9febb59f5689",
				},
				ReturnId: "ba1937fa-14d0-45bd-8b2d-cc38b424ea07",
			},
			ast.Binding{
				FromId: "a4d648a1-29b9-457e-af95-77965a61c4a8",
				FromReturnValue: "ba1937fa-14d0-45bd-8b2d-cc38b424ea07",
				ToReturnValue: "d010a324-fe47-40df-bdb3-15c073d46c86",
			},
		},
	}
}
