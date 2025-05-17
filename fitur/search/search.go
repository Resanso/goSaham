package search

import (
	"fmt"
	"tubesAlpro/fitur/data"
)

// Menampilkan daftar semua saham
func DaftarSaham() {
	fmt.Println("Daftar Saham:")
	for i := 0; i < 5; i++ {
		s := data.ListSaham[i]
		fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
	}
	
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() 
	fmt.Scanln() 
}

// Mencari saham berdasarkan kode
func CariSaham() {
	var kodeSaham string
	fmt.Print("Masukkan kode saham yang ingin dicari: ")
	fmt.Scan(&kodeSaham)
	
	found := false
	for i := 0; i < 5; i++ {
		s := data.ListSaham[i]
		if s.Kode == kodeSaham {
			fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
			found = true
			break
		}
	}
	
	if !found {
		fmt.Println("Saham tidak ditemukan.")
	}
	
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() 
	fmt.Scanln() 
}
