package main

import "fmt"

func main() {
	// Type
	typeDeclare()

	// Loop
	loop()

	// If statement
	ifStatement(1, 2, 3)

	// Switch statement
	choose(2)
}

func typeDeclare() {
	// We have to use the var, otherwise compiler will throw exception
	// main.go:21:2: S1021: should merge variable declaration with assignment on next line (gosimple)
	// var floatValue32 float32
	// floatValue32 = 3.4
	var floatValue32 float32 = 3.4
	var floatValue64 float64 = 3.44 // float default is float64
	fmt.Println(floatValue32, floatValue64)

	var stringValue string = "Hello World"
	// Cannot edit the string while string has been initial.
	// stringValue[1] = '2'
	fmt.Println(stringValue)

	var boolValue bool
	boolValue = true
	fmt.Println(boolValue)

	// We can use := to declare a variable
	goodValue := 111
	fmt.Println(goodValue)
}

func loop() {
	// There is only for loop in Golang
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	// Or can keep empty in for loop
	// This usage just like while loop
	sum2 := 0
	for sum2 < 1000 { // for ; sum2 < 1000;
		sum2 += sum2
	}
	fmt.Println(sum2)

	// infinity loop
	//for {
	//}
}

func ifStatement(c, a int, b int) {
	if c < a && b < a {
		fmt.Println(a, b, c)
	}

	// We can define a variable in side the if statement
	if v := a + b; v < 1000 {
		fmt.Println(v, a, b)
	}
}

func choose(i int) {
	switch i {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4:
		// fallthrough will go to next case
		fallthrough
	case 5:
		fmt.Println("5")
	}
}
