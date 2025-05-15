package main

import "fmt"

func main() {
	menu()
}

func menu() {
var pilihan int
fmt.Println("1. portofolio saya")
fmt.Println("2. Lihat daftar saham")
fmt.Println("3. mencari saham")
fmt.Println("4. transaksi saham")
fmt.Println("5. keluar")

fmt.Scan(&pilihan)
switch pilihan {
case 1:
	fmt.Println("portofolio saya")
case 2:
	fmt.Println("Lihat daftar saham")
case 3:
	fmt.Println("mencari saham")
case 4:
	fmt.Println("transaksi saham")
case 5:
	fmt.Println("keluar")
}
}