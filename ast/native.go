package ast

var NativeValueVoid = 0
var NativeValueInt = 1

type NativeFunctionCall struct {
	Id                string
	Name              string
	ParameterBindings []NativeBinding
	ReturnType        int
	ReturnId          string
}

type BinaryOperationCall struct {
	Id                    string
	Name                  string
	LeftParameterBinding  NativeBinding
	RightParameterBinding NativeBinding
	ReturnId              string
}

type NativeBinding struct {
	FromParameter   string
}
