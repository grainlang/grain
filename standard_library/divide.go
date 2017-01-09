package standard_library

import "github.com/grainlang/grain/ast"

func CreateDivideAst() ast.Function {
	return ast.Function{
		Id: "89ed3305-7b73-45ad-8475-f94e91ee867a",
		Name: "divide",
		Description: "divide two integers",
		Parameters: []ast.Parameter{
			{
				Id: "c557ed6f-bae9-4863-8b37-9fca854569d7",
				Name: "dividend",
				ValueType: ast.Integer,
			},
			{
				Id: "8aed9c05-6b5d-475f-95b6-952a9d346a32",
				Name: "divisor",
				ValueType: ast.Integer,
			},
		},
		ReturnValues: []ast.ReturnValue{
			{
				Id: "72a4c59b-619d-48e2-b108-4057be1c3eee",
				Name: "quotient",
				ValueType: ast.Integer,
			},
			{
				Id: "f2bc71ef-c0b9-42a5-847a-676408d768eb",
				Name: "remainder",
				ValueType: ast.Integer,
			},
		},
		Body: []ast.Expression{
			ast.BinaryOperationCall{
				Id: "85390ef5-3097-4e01-807f-09077c78e158",
				Name: "/",
				LeftParameterBinding: ast.NativeBinding{
					FromParameter: "c557ed6f-bae9-4863-8b37-9fca854569d7",
				},
				RightParameterBinding: ast.NativeBinding{
					FromParameter: "8aed9c05-6b5d-475f-95b6-952a9d346a32",
				},
				ReturnId: "1c062aa4-1873-4b35-a375-9f9c54c2e497",
			},
			ast.Binding{
				FromId: "85390ef5-3097-4e01-807f-09077c78e158",
				FromReturnValue: "1c062aa4-1873-4b35-a375-9f9c54c2e497",
				ToReturnValue: "72a4c59b-619d-48e2-b108-4057be1c3eee",
			},
			ast.BinaryOperationCall{
				Id: "ea81866d-cc2f-49c5-ac2c-3330bcc569e8",
				Name: "%",
				LeftParameterBinding: ast.NativeBinding{
					FromParameter: "c557ed6f-bae9-4863-8b37-9fca854569d7",
				},
				RightParameterBinding: ast.NativeBinding{
					FromParameter: "8aed9c05-6b5d-475f-95b6-952a9d346a32",
				},
				ReturnId: "e287b6b5-8525-42f8-9369-778608cad0a7",
			},
			ast.Binding{
				FromId: "ea81866d-cc2f-49c5-ac2c-3330bcc569e8",
				FromReturnValue: "e287b6b5-8525-42f8-9369-778608cad0a7",
				ToReturnValue: "f2bc71ef-c0b9-42a5-847a-676408d768eb",
			},
		},
	}
}
