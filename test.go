package main

type myInterface interface {
	xPlus()
	xMinor()
}

type myStruct struct{}

func (myStruct) xPlus() {
	return
}
func (myStruct) xMinor() {
	return
}

func Test() {
	var a myInterface
	var b myStruct

	a = b

	a.xMinor()
}
