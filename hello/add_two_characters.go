package hello

import "github.com/grainlang/grain/ast"

func CreateAddTwoCharactersAst() ast.Function {
	return ast.Function{
		Id: "86bf30d1-3eb9-48db-98f3-66c1c0840943",
		Name: "adds two characters from input stream and prints sum as a character",
		Description: "adds two characters from input stream and prints sum as a character",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.FunctionUse{
				Id: "18c36dca-f03c-4ce0-b695-c7cdc48a9dca",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "56c5e818-d957-4f1b-96a6-72bc5e738d58",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "ae693ca3-3327-4eeb-9ce7-9c3768cd86dd",
				FunctionId: "96f1af5d-9d94-428e-a2f5-fc3c63111f27",
				Bindings: []ast.Binding{
					{
						FromId: "18c36dca-f03c-4ce0-b695-c7cdc48a9dca",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "1a6f49fa-5273-4c00-84c5-0278ce1f7041",
					},
					{
						FromId: "56c5e818-d957-4f1b-96a6-72bc5e738d58",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "b08c9d64-04a7-4fe2-b9fa-9febb59f5689",
					},
				},
			},
			ast.FunctionUse{
				Id: "e10415cf-c3a9-40cf-a968-2288470e1436",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromId: "ae693ca3-3327-4eeb-9ce7-9c3768cd86dd",
						FromReturnValue: "d010a324-fe47-40df-bdb3-15c073d46c86",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
		},
	}
}
