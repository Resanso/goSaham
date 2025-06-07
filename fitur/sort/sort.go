package sort

import (
	"fmt"

	data "github.com/Resanso/goSaham/fitur/data"
)

// Fungsi untuk menampilkan daftar saham terurut
func DaftarSahamTerurut(metode string) {
	var sortedSaham [5]data.Saham
	var metodeName string
	
	// Mengurutkan dengan metode yang dipilih
	if metode == "selection" {
		sortedSaham = SelectionSortByPrice(data.ListSaham)
		metodeName = "Selection Sort (Berdasarkan Harga)"
	} else if metode == "insertion" {
		sortedSaham = InsertionSortByVolume(data.ListSaham)
		metodeName = "Insertion Sort (Berdasarkan Volume)"
	}
	
	fmt.Printf("Daftar Saham Terurut (%s):\n", metodeName)
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Printf("%-6s %-25s %-10s %-15s %-12s\n", "Kode", "Nama", "Harga", "Kategori", "Volume")
	fmt.Println("-------------------------------------------------------------------------------")
	
	for i := 0; i < 5; i++ {
		s := sortedSaham[i]
		fmt.Printf("%-6s %-25s %-10.2f %-15s %-12d\n", s.Kode, s.Nama, s.Harga, s.Kategori, s.Volume)
	}
	
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln()
	fmt.Scanln() 
}

// Mengurutkan saham berdasarkan harga tertinggi menggunakan selection sort
func SelectionSortByPrice(saham [5]data.Saham) [5]data.Saham {
	var result [5]data.Saham
		
	// Salin data dari parameter ke array result
	for i := 0; i < 5; i++ {
		result[i] = saham[i]
	}
		
	// Mengurutkan dengan selection sort
	for i := 0; i < 5-1; i++ {
		// Mencari indeks harga tertinggi di array yang belum terurut
		maxIndex := i
		for j := i + 1; j < 5; j++ {
			if result[j].Harga > result[maxIndex].Harga {
				maxIndex = j
			}
		}
		// Menukar elemen jika indeks maksimum berubah
		if maxIndex != i {
			// Swap/menukar elemen
			temp := result[i]
			result[i] = result[maxIndex]
			result[maxIndex] = temp
		}
	}
		
	return result
}
	
// Fungsi insertion sort untuk mengurutkan saham berdasarkan volume tertinggi
func InsertionSortByVolume(saham [5]data.Saham) [5]data.Saham {
	// Membuat array hasil dengan ukuran tetap 5
	var result [5]data.Saham
		
	// Salin data dari parameter ke array result
	for i := 0; i < 5; i++ {
		result[i] = saham[i]
	}
		
	// Mengurutkan dengan insertion sort berdasarkan volume tertinggi
	for i := 1; i < 5; i++ {
		key := result[i]
		j := i - 1
			
		// Geser elemen yang memiliki volume lebih kecil dari key ke kanan
		for j >= 0 && result[j].Volume < key.Volume {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
		
	return result
}
