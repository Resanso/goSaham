package stats

import (
	"fmt"
	"math/rand"
	"time"

	data "github.com/Resanso/goSaham/fitur/data"
)

// init function untuk inisialisasi random seed sekali saja
func init() {
	rand.Seed(time.Now().UnixNano())
}

// CalculatePortfolioValue menghitung total modal yang sudah diinvestasikan
func CalculatePortfolioValue() float64 {
	var totalValue float64 = 0.0
	
	// Hitung total nilai berdasarkan harga beli dan jumlah
	for i := 0; i < data.JumlahPortofolio; i++ {
		item := data.PortofolioUser[i]
		totalValue += float64(item.Jumlah) * item.HargaBeli
	}
	
	return totalValue
}

// GetRandomPrice membuat harga saham yang berfluktuasi secara random
func GetRandomPrice(basePrice float64) float64 {
	// Fluktuasi antara -10% sampai +10%
	fluctuation := (rand.Float64() - 0.5) * 0.2 // menghasilkan angka antara -0.1 sampai +0.1
	newPrice := basePrice * (1 + fluctuation)
	
	// Pastikan harga tidak terlalu rendah (minimal 50% dari harga asli)
	if newPrice < basePrice*0.5 {
		newPrice = basePrice * 0.5
	}
	
	return newPrice
}

// ShowPortfolioStats menampilkan statistik keuntungan dan kerugian portofolio
func ShowPortfolioStats() {
	// Cek apakah portofolio kosong
	if data.JumlahPortofolio == 0 {
		fmt.Println("Portofolio kosong. Tidak ada statistik yang bisa ditampilkan.")
		return
	}
	
	// Header
	fmt.Println("=== STATISTIK KEUNTUNGAN DAN KERUGIAN ===")
	fmt.Println("*** SIMULASI FLUKTUASI PASAR ***")
	fmt.Printf("Jumlah jenis saham: %d\n", data.JumlahPortofolio)
	fmt.Println("(Harga saham berfluktuasi ±10%)")
	fmt.Println("----------------------------------------")
	
	// Hitung total modal yang sudah diinvestasikan
	totalModal := CalculatePortfolioValue()
	
	// Header tabel
	fmt.Printf("%-6s %-8s %-12s %-12s %-12s %-8s\n", 
		"Kode", "Jumlah", "Harga Beli", "Harga Simulasi", "Untung/Rugi", "Status")
	fmt.Println("----------------------------------------------------------------")
	
	var totalNilaiSekarang float64 = 0.0
	
	// Loop untuk setiap saham di portofolio
	for i := 0; i < data.JumlahPortofolio; i++ {
		item := data.PortofolioUser[i]
		
		// Cari harga saham di data master
		var hargaDasar float64
		for j := 0; j < len(data.ListSaham); j++ {
			if data.ListSaham[j].Kode == item.KodeSaham {
				hargaDasar = data.ListSaham[j].Harga
				break
			}
		}
		
		// Buat harga simulasi dengan fluktuasi random
		hargaSimulasi := GetRandomPrice(hargaDasar)
		
		// Hitung nilai investasi
		nilaiModal := float64(item.Jumlah) * item.HargaBeli
		nilaiSekarang := float64(item.Jumlah) * hargaSimulasi
		untungRugi := nilaiSekarang - nilaiModal
		
		totalNilaiSekarang += nilaiSekarang
		
		// Tentukan status
		status := "IMPAS"
		if untungRugi > 0 {
			status = "UNTUNG"
		} else if untungRugi < 0 {
			status = "RUGI"
		}
		
		// Tampilkan detail per saham
		fmt.Printf("%-6s %-8d %-12.2f %-12.2f %-12.2f %-8s\n", 
			item.KodeSaham, item.Jumlah, item.HargaBeli, hargaSimulasi, untungRugi, status)
	}
	
	// Hitung total untung/rugi
	totalUntungRugi := totalNilaiSekarang - totalModal
	persentase := (totalUntungRugi / totalModal) * 100
	
	fmt.Println("----------------------------------------------------------------")
	
	// Tampilkan ringkasan
	fmt.Printf("Total Modal Investasi    : Rp %.2f\n", totalModal)
	fmt.Printf("Total Nilai Saat Ini     : Rp %.2f\n", totalNilaiSekarang)
	fmt.Printf("Total Keuntungan/Kerugian: Rp %.2f\n", totalUntungRugi)
	fmt.Printf("Persentase Return        : %.2f%%\n", persentase)
	
	// Status keseluruhan
	if totalUntungRugi > 0 {
		fmt.Println("Status: UNTUNG - Portofolio mengalami keuntungan")
	} else if totalUntungRugi < 0 {
		fmt.Println("Status: RUGI - Portofolio mengalami kerugian")
	} else {
		fmt.Println("Status: IMPAS - Tidak ada keuntungan atau kerugian")
	}
	
	// Tunggu input user untuk kembali ke menu
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln()
	fmt.Scanln()
}
