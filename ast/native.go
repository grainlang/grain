package ast

var NativeValueVoid = "void"
var NativeValueInt = "int"

type NativeFunctionCall struct {
	Name        string
	Parameters  []Parameter
	ReturnValue string
}
