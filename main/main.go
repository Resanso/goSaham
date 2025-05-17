package main

import (
	"fmt"
	"tubesAlpro/fitur/portfolio"
	"tubesAlpro/fitur/search"
	"tubesAlpro/fitur/sort"
	"tubesAlpro/fitur/stats"
)

func main() {
	menu()
}

func menu() {
	var pilihan int
	running := true
	
	for running {
		fmt.Println("\n===== APLIKASI MANAJEMEN PORTOFOLIO SAHAM =====")
		fmt.Println("1. Portofolio Saya")
		fmt.Println("2. Lihat Daftar Saham")
		fmt.Println("3. Mencari Saham")
		fmt.Println("4. Transaksi Saham")
		fmt.Println("5. Lihat Saham Terurut (Selection Sort)")
		fmt.Println("6. Lihat Saham Terurut (Insertion Sort)")
		fmt.Println("7. Statistik Portofolio")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu (1-8): ")
		
		fmt.Scan(&pilihan)
		
		switch pilihan {
		case 1:
			fmt.Println("\n--- PORTOFOLIO SAYA ---")
			portfolio.Portofolio()
		case 2:
			fmt.Println("\n--- DAFTAR SAHAM ---")
			search.DaftarSaham()
		case 3:
			fmt.Println("\n--- PENCARIAN SAHAM ---")
			search.CariSaham()
		case 4:
			fmt.Println("\n--- TRANSAKSI SAHAM ---")
			portfolio.Transaksi()
		case 5:
			fmt.Println("\n--- DAFTAR SAHAM TERURUT (SELECTION SORT) ---")
			sort.DaftarSahamTerurut("selection")
		case 6:
			fmt.Println("\n--- DAFTAR SAHAM TERURUT (INSERTION SORT) ---")
			sort.DaftarSahamTerurut("insertion")
		case 7:
			fmt.Println("\n--- STATISTIK PORTOFOLIO ---")
			stats.ShowPortfolioStats()
		case 8:
			running = false
			keluar()
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1-8.")
		}
	}
}

func keluar() {
	fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
}
