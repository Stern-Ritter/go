package main

type Apple struct{}

func NewApple() Apple {
	return Apple{}
}

func (a Apple) String() string {
	return "Яблоко"
}
