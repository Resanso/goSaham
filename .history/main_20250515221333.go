package main

import "fmt"
func main() {
	var isLoggedIn bool
	if isLoggedIn == false {
		login(isLoggedIn)
	}
}

func login(isLoggedIn bool) bool {
	var username string
	var password string
	username = "admin"
	password = "admin"
	fmt.Println("masukkan username dan password")
	fmt.Scanln(&username)
	fmt.Scanln(&password)
	if username != "admin" || password != "admin" {
		fmt.Println("username atau password salah")
		return false
	} else {
		fmt.Println("login berhasil")
		fmt.Println("selamat datang", username)
		return true
	}
}
