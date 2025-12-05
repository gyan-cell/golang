//methods in Golang is bit weird , it does not support a whole lot of object oriented programming but
// it allows us to create the methods on types or structs but instead of defining it in our main methods we
// define it seprately through methods , using reciever argument

package main

import "fmt"

type User struct {
	Name   string
	Id     int
	status bool
}

// now suppose i wanna have a method getStatus in that struct like we can define mthods in js objects what we will do is we create a function

func (user User) GetStatus() bool { // This function is now exportable and now this function belongs to struct User  notice how it does not change the
	// the orignal struct , because it does not take the original value
	return user.status
}

func (user *User) ChangeStatus() bool {
	user.status = !user.status
	return user.status
}

// Value Receiver (WON'T modify original)
func (user User) TryToChangeStatus() bool {
	user.status = !user.status // Only changes the copy!
	return user.status
}

//  You can declare a method on non-struct types, too.
//
// In this example we see a numeric type MyFloat with an Abs method.
//
// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as i

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	fmt.Println("Explaining Go methods")

	userBlade := User{
		Name:   "Blade", // Use named fields (recommended)
		Id:     12,
		status: false,
	}

	fmt.Println("Initial Status:", userBlade.GetStatus())

	// This works but only changes the copy
	fmt.Println("Try to change (value receiver):", userBlade.TryToChangeStatus())
	fmt.Println("Status after TryToChangeStatus:", userBlade.GetStatus()) // Still false!

	// This actually changes the original
	fmt.Println("Change status (pointer receiver):", userBlade.ChangeStatus())
	fmt.Println("Status after ChangeStatus:", userBlade.GetStatus()) // Now true!

	// Alternative syntax with pointer
	userPtr := &userBlade
	fmt.Println("\nUsing pointer directly:")
	fmt.Println("Status via pointer:", userPtr.GetStatus()) // Go automatically dereferences!
	userPtr.ChangeStatus()
	fmt.Println("Status after change:", userPtr.GetStatus())
}
