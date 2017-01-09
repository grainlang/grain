package ast

var NativeValueVoid = 0
var NativeValueInt = 1

type NativeFunctionCall struct {
	Id         string
	Name       string
	Parameters []Parameter
	ReturnType int
	ReturnId   string
}

type BinaryOperationCall struct {
	Id             string
	Name           string
	LeftParameter  Parameter
	RightParameter Parameter
	ReturnId       string
}
