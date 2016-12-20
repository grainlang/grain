package ast

type Type struct {
	Id          string
	Name        string
	Description string
}

var Character = Type{
	Id: "c9ab1171-8e9d-496f-8010-5681f47469a6",
	Name: "character",
}

var Integer = Type{
	Id: "87f4d4f1-52e2-4291-bfad-dfcb61878bce",
	Name: "integer",
	Description: "64 bits",
}
