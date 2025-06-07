package search

import (
	"fmt"
	"strings"

	data "github.com/Resanso/goSaham/fitur/data"
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

// Mencari saham berdasarkan kode (Sequential Search)
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

// Mencari saham berdasarkan nama perusahaan (Binary Search)
func CariSahamBinarySearch() {
	var namaSaham string
	fmt.Print("Masukkan nama perusahaan yang ingin dicari: ")
	fmt.Scanln() // untuk membersihkan buffer
	fmt.Scanln(&namaSaham)
	
	// Convert input to lowercase
	namaSaham = strings.ToLower(namaSaham)
	
	left := 0
	right := len(data.ListSaham) - 1
	found := false
	
	for left <= right {
		mid := (left + right) / 2
		midNama := strings.ToLower(data.ListSaham[mid].Nama)
		
		if strings.Contains(midNama, namaSaham) {
			// Ditemukan, tampilkan data saham
			s := data.ListSaham[mid]
			fmt.Printf("Saham ditemukan!\n")
			fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
			found = true
			
			// Cari kemungkinan match lain di sekitar posisi ini
			// Cek ke kiri
			for i := mid - 1; i >= 0; i-- {
				if strings.Contains(strings.ToLower(data.ListSaham[i].Nama), namaSaham) {
					s := data.ListSaham[i]
					fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
				} else {
					break
				}
			}
			
			// Cek ke kanan
			for i := mid + 1; i < len(data.ListSaham); i++ {
				if strings.Contains(strings.ToLower(data.ListSaham[i].Nama), namaSaham) {
					s := data.ListSaham[i]
					fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
				} else {
					break
				}
			}
			break
		} else if midNama < namaSaham {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	if !found {
		fmt.Println("Saham dengan nama tersebut tidak ditemukan.")
	}
	
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() 
}
