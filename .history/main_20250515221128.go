package main

import "fmt"
func main() {
	var isLoggedIn bool
	if isLoggedIn == false {
		login(isLoggedIn)
	}
}

func login(isLoggedIn bool) bool {
	fmt.Println("User logged in")
	return true
}
