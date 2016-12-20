package ast

type Function struct {
	Id           string
	Name         string
	Description  string
	Parameters   []Parameter
	ReturnValues []ReturnValue
	Body         []Expression
}

type Parameter struct {
	Id          string
	Name        string
	Description string
	ValueType   Type
}

type ReturnValue struct {
	Id          string
	Name        string
	Description string
	ValueType   Type
}

type Expression interface {
}

type FunctionUse struct {
	Id         string
	FunctionId string
	Bindings   []Binding
}

type Binding struct {
	FromFunctionUseId string
	FromReturnValue   string
	ToParameter       string
}
