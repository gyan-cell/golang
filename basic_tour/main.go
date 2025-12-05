package main

//  This code groups the imports into a parenthesized, "factored" import statement.
//
// You can also write multiple import statements, like:
//
// import "fmt"
// import "math"
//

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My fav No. is ", rand.Intn(10))
}

// Exported names
//
// In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
//
// pizza and pi do not start with a capital letter, so they are not exported.
//
// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
//
// Run the code. Notice the error message.
//
// To fix the error, rename math.pi to math.Pi and try it again.
// the below function is exported , because it starts with the capital letter
func Sum(a, b int) int {
	return a + b
}

// this function is not exported at all
func funkyman() {
	fmt.Println("seee my funck")
}

//
//  A function can take zero or more arguments.
//
// In this example, add takes two parameters of type int.
//
// Notice that the type comes after the variable name.
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
//
// In this example, we shortened
//
//  A function can return any number of results.
//
// The swap function returns two strings.

func swap(x, y string) (string, string) {
	return y, x
}

func swapCaller() {
	a, b := "hello ", "World"
	swap(a, b)
}

//  Go's return values may be named. If so, they are treated as variables defined at the top of the function.
//
// These names should be used to document the meaning of the return values.
//
// A return statement without arguments returns the named return values. This is known as a "naked" return.
//
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//
// The var statement declares a list of variables; as in function argument lists, the type is last.
//
// A var statement can be at package or function level. We see both in this example.
//  A var declaration can include initializers, one per variable.
//
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

//  Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
//
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

// Basic types
//
// Go's basic types are
//
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // alias for uint8
//
// rune // alias for int32
//      // represents a Unicode code point
//
// float32 float64
//
// complex64 complex128
//
// The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
//
// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
//

// Constants are declared like variables, but with the const keyword.
//
// Constants can be character, string, boolean, or numeric values.
//
// Constants cannot be declared using the := syntax.

//  Go has only one looping construct, the for loop.
//
// The basic for loop has three components separated by semicolons:
//
//     the init statement: executed before the first iteration
//     the condition expression: evaluated before every iteration
//     the post statement: executed at the end of every iteration
//
// The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
//
// The loop will stop iterating once the boolean condition evaluates to false.

//  A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
//
// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

//  A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
//
// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

func defer_func() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
