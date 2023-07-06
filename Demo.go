package main

import (
	"demo1/demo/check"
	"fmt"
)

func main() {
	//var x float64
	//x = 20.0
	//var y float32 = 30.0
	//z := 40
	//fmt.Println(x)
	//fmt.Printf("Kiểu dữ liệu của x là %T\n", x)
	//fmt.Println(y)
	//fmt.Printf("Kiểu dữ liệu của y là %T\n", y)
	//fmt.Println(z)
	//fmt.Printf("Kiểu dữ liệu của z là %T\n", z)
	//var a, b, c = 1, 2, "rikkei"
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Printf("Kiểu dữ liệu của a là %T\n", a)
	//fmt.Printf("Kiểu dữ liệu của b là %T\n", b)
	//fmt.Printf("Kiểu dữ liệu của c là %T\n", c)

	//Access modifier
	check.Check()
	fmt.Println(check.Other2)

	//Const
	const hello = "Sam"
	name := hello
	fmt.Printf("type %T value %v \n", name, name)

	//const type boolean
	//const trueConst = true
	//type myBool bool
	//var defaultBool = trueConst       //chạy bình thường
	//var customBool myBool = trueConst //chạy bình thường
	//defaultBool = customBool          //lỗi do 2 biến có kiểu khác nhau

	// Const type number
	const test = 5
	var intVar int = test
	var int32Var int32 = test
	var float64Var float64 = test
	var complex64Var complex64 = test
	fmt.Println(test == intVar)
	fmt.Println(test == int32Var)
	fmt.Println(test == float64Var)
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	//For
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// If -else
	a := 10
	if a > 20 {
		fmt.Println(">")
	} else if a < 20 {
		fmt.Println("<")
	} else {
		fmt.Println("=")
	}

	//Switch-case
	x := 42.0
	switch x {
	case 0:
	case 1, 2:
		fmt.Println("Multiple matches")
	case 42: // Don't "fall through".
		fmt.Println("reached")
	case 43:
		fmt.Println("Unreached")
	default:
		fmt.Println("Optional")
	}
}
