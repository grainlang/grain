package ast

var NativeValueVoid = "void"
var NativeValueInt = "int"
var NativeValueInt64 = "int64"

type NativeFunctionCall struct {
	Name        string
	Parameters  []Parameter
	ReturnValue string
}

type BinaryOperationCall struct {
	Name           string
	LeftParameter  Parameter
	RightParameter Parameter
	ReturnValue    string
}
