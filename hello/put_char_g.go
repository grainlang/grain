package hello

import "github.com/grainlang/grain/ast"

func CreatePutCharConstGAst() ast.Function {
	return ast.Function{
		Id: "e0061370-d389-43b3-a203-8f58d7a3b374",
		Name: "print out G",
		Description: "print out G",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.Constant{
				Id: "263fce84-647e-4c2f-a1a9-448f129185a4",
				Name: "value G",
				Value: "G",
				ValueType: ast.Character,
			},
			ast.FunctionUse{
				Id: "ca93212b-db27-439d-adc6-054c62f10767",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromConstant: "263fce84-647e-4c2f-a1a9-448f129185a4",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
		},
	}
}
