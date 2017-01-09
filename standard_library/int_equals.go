package standard_library

import "github.com/grainlang/grain/ast"

func CreateIntEqualsAst() ast.Function {
	return ast.Function{
		Id: "f3be3f89-1a7b-4d49-af47-3970f05a05b9",
		Name: "equals",
		Description: "check two intergers for equality",
		Parameters: []ast.Parameter{
			{
				Id: "b92e126e-ad0f-4e0a-8548-0ae45e151d78",
				Name: "compared",
				ValueType: ast.Integer,
			},
			{
				Id: "cf098c89-d6be-4132-8776-9330dbf74df6",
				Name: "compared",
				ValueType: ast.Integer,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "6747da77-dbaa-4a65-b61d-49e1c13142b5",
				Name: "equal",
				ValueType: ast.Boolean,
			},
		},
		Body: []ast.Expression{
			ast.BinaryOperationCall{
				Id: "2a03cc2e-a770-4f37-b681-9b07c7b26820",
				Name: "=",
				LeftParameterBinding: ast.NativeBinding{
					FromParameter: "b92e126e-ad0f-4e0a-8548-0ae45e151d78",
				},
				RightParameterBinding: ast.NativeBinding{
					FromParameter: "cf098c89-d6be-4132-8776-9330dbf74df6",
				},
				ReturnId: "0840d95d-c76b-4d5f-afbc-a2ae907cecdf",
			},
			ast.Binding{
				FromId: "2a03cc2e-a770-4f37-b681-9b07c7b26820",
				FromReturnValue: "0840d95d-c76b-4d5f-afbc-a2ae907cecdf",
				ToReturnValue: "6747da77-dbaa-4a65-b61d-49e1c13142b5",
			},
		},
	}
}
