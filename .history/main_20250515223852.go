package main

import "fmt"

// Struktur untuk data saham
type Saham struct {
	Kode     string
	Nama     string
	Harga    float64
	Kategori string
}

// Variabel global untuk menyimpan data saham
var listSaham = []Saham{
	{"BBCA", "Bank Central Asia", 9000.0, "Perbankan"},
	{"TLKM", "Telkom Indonesia", 3500.0, "Telekomunikasi"},
	{"ASII", "Astra International", 6800.0, "Otomotif"},
	{"UNVR", "Unilever Indonesia", 4500.0, "Konsumer"},
	{"GOTO", "GoTo Gojek Tokopedia", 75.0, "Teknologi"},
}

// Struktur untuk data portofolio
type PortofolioItem struct {
	KodeSaham string
	Jumlah    int
	HargaBeli float64
}

// Variabel global untuk menyimpan portofolio pengguna
var portofolioUser []PortofolioItem

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
		Portofolio()
	case 2:
		fmt.Println("Lihat daftar saham")
		daftarSaham()
	case 3:
		fmt.Println("mencari saham")
		cariSaham()
	case 4:
		fmt.Println("transaksi saham")
		transaksi()
	case 5:
		fmt.Println("keluar")
		keluar()
	}
}

func Portofolio() {
	if len(portofolioUser) == 0 {
		fmt.Println("Portofolio Anda kosong.")
	} else {
		fmt.Println("\n===== Portofolio Anda =====")
		fmt.Println("----------------------------------------------------------------")
		fmt.Printf("%-8s %-6s %-10s %-12s %-12s %-10s\n", 
			"Kode", "Jumlah", "Harga Beli", "Harga Saat Ini", "Nilai Total", "Profit/Loss")
		fmt.Println("----------------------------------------------------------------")
		
		var totalNilai, totalInvestasi float64
		
		for _, p := range portofolioUser {
			// Cari harga saham saat ini
			hargaSekarang := 0.0
			for _, s := range listSaham {
				if s.Kode == p.KodeSaham {
					hargaSekarang = s.Harga
					break
				}
			}
			
			// Hitung nilai total dan profit/loss
			investasi := float64(p.Jumlah) * p.HargaBeli
			nilaiSekarang := float64(p.Jumlah) * hargaSekarang
			profitLoss := nilaiSekarang - investasi
			profitLossPercentage := (profitLoss / investasi) * 100
			
			// Tampilkan data per saham
			fmt.Printf("%-8s %-6d %-10.2f %-12.2f %-12.2f ", 
				p.KodeSaham, p.Jumlah, p.HargaBeli, hargaSekarang, nilaiSekarang)
			
			// Tampilkan profit/loss dengan warna (tidak bisa diimplementasikan di terminal standar)
			if profitLoss >= 0 {
				fmt.Printf("+%-9.2f (+%.2f%%)\n", profitLoss, profitLossPercentage)
			} else {
				fmt.Printf("%-10.2f (%.2f%%)\n", profitLoss, profitLossPercentage)
			}
			
			// Akumulasi total
			totalNilai += nilaiSekarang
			totalInvestasi += investasi
		}
		
		// Tampilkan total portofolio
		totalProfitLoss := totalNilai - totalInvestasi
		totalProfitLossPercentage := (totalProfitLoss / totalInvestasi) * 100
		
		fmt.Println("----------------------------------------------------------------")
		fmt.Printf("%-39s %-12.2f ", "Total Nilai Portofolio:", totalNilai)
		if totalProfitLoss >= 0 {
			fmt.Printf("+%-9.2f (+%.2f%%)\n", totalProfitLoss, totalProfitLossPercentage)
		} else {
			fmt.Printf("%-10.2f (%.2f%%)\n", totalProfitLoss, totalProfitLossPercentage)
		}
		fmt.Println("----------------------------------------------------------------")
	}
	
	// Kembali ke menu utama setelah menampilkan portofolio
	fmt.Println("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	menu()
}

func daftarSaham() {
	fmt.Println("Daftar Saham:")
	for _, s := range listSaham {
		fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
	}
}

func cariSaham() {
	var kodeSaham string
	fmt.Print("Masukkan kode saham yang ingin dicari: ")
	fmt.Scan(&kodeSaham)

	found := false
	for _, s := range listSaham {
		if s.Kode == kodeSaham {
			fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Saham tidak ditemukan.")
	}
}

func transaksi() {
	var kodeSaham string
	var jumlah int
	var hargaBeli float64

	fmt.Print("Masukkan kode saham yang ingin dibeli: ")
	fmt.Scan(&kodeSaham)
	fmt.Print("Masukkan jumlah saham yang ingin dibeli: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan harga beli per saham: ")
	fmt.Scan(&hargaBeli)

	portofolioUser = append(portofolioUser, PortofolioItem{KodeSaham: kodeSaham, Jumlah: jumlah, HargaBeli: hargaBeli})
	fmt.Println("Transaksi berhasil. Saham telah ditambahkan ke portofolio Anda.")
}

func keluar() {
	fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
}