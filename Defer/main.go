/*
Defer Note:

Goal: Prove that defer is execute "before" whole surrounding func return.

Note: Some people in stack overflow my acclaim that defer execute after func return.
That is because defer value execute after "return" set result parameter in return statement

	Return of result parameter example:
		func foo(a int)(b int) {
			return 2
		}

	Can be wrote as below:
		func foo(a int)(b int) {
			b = 2
			return
		}
	(e.g. return 2 --> a = 2)

But the defer function will be executed before the whole surrounding function return.
For detail info, see the exampled below.

Reference: https://go.dev/ref/spec#Defer_statements
*/
package main

import (
	"fmt"
)

// expect defer: 1, main: 1
func num1() (a int) {
	/* 
		Defer function will be executed after return assgin a = 2
		As you can see in defer anonymous function I assign a = 1
		Thus make the real return value a become 1.
	*/
	defer func() {
		a = 1
		fmt.Printf("num() defer invoke here: %v\n", a)
	}()

	return 2
}

// expect defer: 1, main: 1
func num2() (a int) {
	/* 
		return will be execute olny all the command is executed.
		to refactor num()
		"return 2" + "defer function" can be regard as

		a = 2
		a = 1
		fmt.printf(...)
		return

		Howerver, defer function assign 1 to a, as a result, at the time of return, num2 return 1
	*/
	defer func() {
		a = 1
		fmt.Printf("num() defer invoke here: %v\n", a)
	}()

	a = 2
	return
}

// expected 0, 9.
// At the time declare defer, parameters of fmt.Printf: "x" has been assigns as 0 (zero value)
func test1() (x int) {
	defer fmt.Printf("in defer: x = %d\n", x)

	x = 7
	return 9
}

// expect print defer: 9 & main: 9
// At the time declare defer, function is defined but is invoked at once.
// Defer function is executed after x is assigned to 9.
func test2() (x int) {
	defer func() {
			fmt.Printf("in defer: x = %d\n", x)
	}()

	x = 7
	return 9
}

// expect defer: 7, main:9
// At the time we declare defer function, x is assigned as 7.
func test3() (x int) {
	/*
		1. Define x as return value
		3. Assign 7 to x
		2. Invoke defer, x is evaluated as 7
		4. Set return parameter 9 (x is assigned to value = 9)
		5. Execute deferred function, at this time, x have already become 9
		6. Return
	*/
	
	x = 7
	defer fmt.Printf("in defer: x = %d\n", x)
	return 9
}

/* 
	expect n = 0, x = 9, return 9
	n in func is getting zero value of input parameter x
	x get the value at the end of outer function
*/
func test4() (x int) {
	defer func(n int) {
		fmt.Printf("in defer x as parameter: x = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}(x)
		
	x = 7
	return 9
}

/*
	function & parameter 
	defer:  50, 90
	main: 100, 9
*/
func test5() (n int, x int) {
	n = 10
	y := 50
	z :=100
	defer func(n int) {
		// In inner function block, the parameter n shadow（遮蔽）the n outside its function
		// Thus, n in inner block get parameter y = 50;
		// While z -> get the value outside the function (coupling)
		// Think as z getting value at the end of outer function. 
		fmt.Printf("in defer n as parameter: n = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", z)
	}(y)

	n = 100
	z = 90
	x = 9
	return n, x
}

// expected return 100, 9
func test6() (n int, x int) {
	n = 10
	defer func() {
		// with no parameters, defer get the value after return value is assigned.
		fmt.Printf("in defer n as parameter: n = %d\n", n)
		fmt.Printf("in defer x after return: x = %d\n", x)
	}()
		
	return 100, 9
}

func main() {
	fmt.Printf("num1() return: %d\n", num1()) // get 1
	fmt.Printf("num2() return: %d\n", num2()) // get 1
	fmt.Println("test1")
	fmt.Printf("in main: x = %d\n", test1())
	fmt.Println("test2")
	fmt.Printf("in main: x = %d\n", test2())
	fmt.Println("test3")
	fmt.Printf("in main: x = %d\n", test3())
	fmt.Println("test4")
	fmt.Printf("in main: x = %d\n", test4())
	fmt.Println("test5")
	n5, x5 := test5()
	fmt.Printf("Test 5 in main: n= %d, x = %d\n", n5, x5)
	n6, x6 := test6()
	fmt.Printf("Test 6 in main: n= %d, x = %d\n", n6, x6)
}
