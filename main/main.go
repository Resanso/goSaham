package main

import (
	"fmt"

	portfolio "github.com/Resanso/goSaham/fitur/portfolio"
	search "github.com/Resanso/goSaham/fitur/search"
	sort "github.com/Resanso/goSaham/fitur/sort"
	stats "github.com/Resanso/goSaham/fitur/stats"
	transaksi "github.com/Resanso/goSaham/fitur/transaksi"
)

func main() {
	menu()
}

func menu() {
	running := true
	
	for running {
		fmt.Println("\n===== APLIKASI MANAJEMEN PORTOFOLIO SAHAM =====")
		fmt.Println("1. Portofolio Saya")
		fmt.Println("2. Lihat Daftar Saham")
		fmt.Println("3. Mencari Saham (Sequential Search - Berdasarkan Kode)")
		fmt.Println("4. Mencari Saham (Binary Search - Berdasarkan Nama)")
		fmt.Println("5. Transaksi Saham")
		fmt.Println("6. Lihat Saham Terurut (Selection Sort)")
		fmt.Println("7. Lihat Saham Terurut (Insertion Sort)")
		fmt.Println("8. Statistik Portofolio")
		fmt.Println("9. keluar")
		
		var pilihan int
		fmt.Print("Pilih menu (1-10): ")
		_, err := fmt.Scanf("%d", &pilihan)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
				switch pilihan {
		case 1:
			fmt.Println("\n--- PORTOFOLIO SAYA ---")
			portfolio.Portofolio()
		case 2:
			fmt.Println("\n--- DAFTAR SAHAM ---")
			search.DaftarSaham()
		case 3:
			fmt.Println("\n--- PENCARIAN SAHAM (SEQUENTIAL SEARCH) ---")
			search.CariSaham()
		case 4:
			fmt.Println("\n--- PENCARIAN SAHAM (BINARY SEARCH) ---")
			search.CariSahamBinarySearch()
		case 5:
			fmt.Println("\n--- TRANSAKSI SAHAM ---")
			transaksi.MenuTransaksi()
		case 6:
			fmt.Println("\n--- DAFTAR SAHAM TERURUT (SELECTION SORT) ---")
			sort.DaftarSahamTerurut("selection")
		case 7:
			fmt.Println("\n--- DAFTAR SAHAM TERURUT (INSERTION SORT) ---")
			sort.DaftarSahamTerurut("insertion")
		case 8:
			fmt.Println("\n--- STATISTIK PORTOFOLIO ---")
			stats.ShowPortfolioStats()
		case 9:
			running = false
			keluar()
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1-10.")
		}
	}
}

func keluar() {
	fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
}
