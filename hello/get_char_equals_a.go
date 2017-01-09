package hello

import "github.com/grainlang/grain/ast"

func CreateGetCharEqualsAAst() ast.Function {
	return ast.Function{
		Id: "f187aa24-ed45-4a04-950e-d0cf250e96be",
		Name: "prints something when character is equal to a, prints something else otherwise",
		Description: "prints something when character is equal to a, prints something else otherwise",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.Constant{
				Id: "076bb862-ab45-47d9-911f-7efa52444745",
				Name: "int representation of a",
				Value: "97",
				ValueType: ast.Integer,
			},
			ast.FunctionUse{
				Id: "3e27029f-0973-4fa7-b9a0-6dc4c8306899",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "ebe07811-c787-4085-93a1-fc18889e40a7",
				FunctionId: "f3be3f89-1a7b-4d49-af47-3970f05a05b9",
				Bindings: []ast.Binding{
					{
						FromId: "3e27029f-0973-4fa7-b9a0-6dc4c8306899",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "b92e126e-ad0f-4e0a-8548-0ae45e151d78",
					},
					{
						FromConstant: "076bb862-ab45-47d9-911f-7efa52444745",
						ToParameter: "cf098c89-d6be-4132-8776-9330dbf74df6",
					},
				},
			},
			ast.FunctionUse{
				Id: "d05c9851-cf86-4834-8a21-61807ebe14ec",
				FunctionId: "96f1af5d-9d94-428e-a2f5-fc3c63111f27",
				Bindings: []ast.Binding{
					{
						FromId: "ebe07811-c787-4085-93a1-fc18889e40a7",
						FromReturnValue: "6747da77-dbaa-4a65-b61d-49e1c13142b5",
						ToParameter: "1a6f49fa-5273-4c00-84c5-0278ce1f7041",
					},
					{
						FromConstant: "076bb862-ab45-47d9-911f-7efa52444745",
						ToParameter: "b08c9d64-04a7-4fe2-b9fa-9febb59f5689",
					},
				},
			},
			ast.FunctionUse{
				Id: "3c1f536a-d934-489a-9399-2c6664e7a09b",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromId: "d05c9851-cf86-4834-8a21-61807ebe14ec",
						FromReturnValue: "d010a324-fe47-40df-bdb3-15c073d46c86",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
		},
	}
}
