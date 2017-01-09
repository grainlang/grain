package ast

type Constant struct {
	Id          string
	Name        string
	Description string
	Value       string
	ValueType   Type
}
