package hello

import "github.com/grainlang/grain/ast"

func CreateVoodooCalculationsAst() ast.Function {
	return ast.Function{
		Id: "daa2455e-bf63-458f-9ac0-c37dca59476a",
		Name: "gets three characters from input and prints out some different characters",
		Description: "gets three characters from input and prints out some different characters",
		Parameters: []ast.Parameter{},
		ReturnValues: []ast.ReturnValue{},
		Body: []ast.Expression{
			ast.FunctionUse{
				Id: "2b590032-208b-47dd-a192-b23be7203aa6",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "6e3de88a-b146-4396-9f08-ceaf4a5b94ce",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "118efdcd-be74-467f-805d-1bb00c8c49e6",
				FunctionId: "41a75f38-63bd-46f5-8cbe-78b32f0698bc",
			},
			ast.FunctionUse{
				Id: "9842b1d8-0cb0-46c8-a594-d034f0df8f6d",
				FunctionId: "b45f802e-39ca-48d5-9368-e9ffcdea0d27",
				Bindings: []ast.Binding{
					{
						FromId: "2b590032-208b-47dd-a192-b23be7203aa6",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "b787a111-3f85-4053-a8c3-e1dae3ce6fe4",
					},
					{
						FromId: "6e3de88a-b146-4396-9f08-ceaf4a5b94ce",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "a0a9a1b3-ef5f-469f-ae78-9dc9aedcf510",
					},
				},
			},
			ast.FunctionUse{
				Id: "5b95597f-cc45-47f3-868f-d100f3cb5cba",
				FunctionId: "89ed3305-7b73-45ad-8475-f94e91ee867a",
				Bindings: []ast.Binding{
					{
						FromId: "9842b1d8-0cb0-46c8-a594-d034f0df8f6d",
						FromReturnValue: "f30fdb67-d967-47e8-8745-bd2f9770bd4f",
						ToParameter: "c557ed6f-bae9-4863-8b37-9fca854569d7",
					},
					{
						FromId: "118efdcd-be74-467f-805d-1bb00c8c49e6",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "8aed9c05-6b5d-475f-95b6-952a9d346a32",
					},
				},
			},
			ast.FunctionUse{
				Id: "6c2c29ba-5562-44cb-a2d0-dfc21b0034a1",
				FunctionId: "b45f802e-39ca-48d5-9368-e9ffcdea0d27",
				Bindings: []ast.Binding{
					{
						FromId: "5b95597f-cc45-47f3-868f-d100f3cb5cba",
						FromReturnValue: "f2bc71ef-c0b9-42a5-847a-676408d768eb",
						ToParameter: "b787a111-3f85-4053-a8c3-e1dae3ce6fe4",
					},
					{
						FromId: "118efdcd-be74-467f-805d-1bb00c8c49e6",
						FromReturnValue: "6d1a6e69-a263-422f-8f19-1d19299af5c9",
						ToParameter: "a0a9a1b3-ef5f-469f-ae78-9dc9aedcf510",
					},
				},
			},
			ast.FunctionUse{
				Id: "83f265b1-5ce8-479d-9e98-ad8ea9b577c8",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromId: "5b95597f-cc45-47f3-868f-d100f3cb5cba",
						FromReturnValue: "72a4c59b-619d-48e2-b108-4057be1c3eee",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
			ast.FunctionUse{
				Id: "efb67f26-5341-4899-a994-150e39cedfcb",
				FunctionId: "0f6e40da-37a1-42ff-b98a-a25d0da5ca45",
				Bindings: []ast.Binding{
					{
						FromId: "6c2c29ba-5562-44cb-a2d0-dfc21b0034a1",
						FromReturnValue: "f30fdb67-d967-47e8-8745-bd2f9770bd4f",
						ToParameter: "acbbd3f4-858a-40cb-a465-b2996bc2cf8f",
					},
				},
			},
		},
	}
}
