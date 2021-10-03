package main

import "fmt"

type Role int

const (
	User Role = 1 << iota
	Manager
	Admin
	Customer
	RoleMask = (1 << (iota)) - 1
)

func (r Role) String() string {
	result := ""
	if User&r > 0 {
		result += "User "
	}
	if Manager&r > 0 {
		result += "Manager "
	}
	if Admin&r > 0 {
		result += "Admin "
	}
	if Customer&r > 0 {
		result += "Customer"
	}
	if r&RoleMask == 0 {
		result = "Invalid role"
	}
	return result
}

func main() {
	fmt.Printf("%s - %[1]d : %[1]T, Mask: %b\n", User+Manager, RoleMask)
	fmt.Printf("%s - %[1]d: %[1]T, Mask: %b\n", Manager, RoleMask)
	fmt.Printf("%s - %[1]d : %[1]T, Mask: %b\n", Admin, RoleMask)
	fmt.Printf("%s - %[1]d : %[1]T, Mask: %b\n", Customer, RoleMask)
}

// fmt package defines interface Stringer as:
type Stringer interface {
	String() string
}
