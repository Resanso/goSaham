package main

import "fmt"
func main() {
	var isLoggedIn bool
	if isLoggedIn == false {
		login(isLoggedIn)
	}
	menu()
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

func menu() {
	fmt.Println("1. menu 1")
	fmt.Println("2. menu 2")
	fmt.Println("3. menu 3")
	fmt.Println("4. menu 4")
	fmt.Println("5. menu 5")
	fmt.Println("6. menu 6")
	fmt.Println("7. menu 7")
	fmt.Println("8. menu 8")
	fmt.Println("9. menu 9")
	fmt.Println("10. menu 10")
	fmt.Println("11. menu 11")
	fmt.Println("12. menu 12")
	fmt.Println("13. menu 13")
	fmt.Println("14. menu 14")
	fmt.Println("15. menu 15")
}