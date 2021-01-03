package easymodels

type Pet interface {
	Name() string
	SetName(name string)

	Age() int
	SetAge(age int)

	Kind() string
}
