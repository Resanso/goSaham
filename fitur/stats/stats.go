package stats

import (
	"fmt"
	data "github.com/Resanso/goSaham/fitur/data"
)

// CalculatePortfolioValue menghitung total nilai portofolio
func CalculatePortfolioValue() float64 {
	var totalValue float64 = 0.0
	
	// Hitung total nilai berdasarkan harga beli dan jumlah
	for i := 0; i < data.JumlahPortofolio; i++ {
		item := data.PortofolioUser[i]
		totalValue += float64(item.Jumlah) * item.HargaBeli
	}
	
	return totalValue
}

// ShowPortfolioStats menampilkan statistik portofolio
func ShowPortfolioStats() {
	if data.JumlahPortofolio == 0 {
		fmt.Println("Portofolio kosong. Tidak ada statistik yang bisa ditampilkan.")
		return
	}
	
	totalValue := CalculatePortfolioValue()
	
	fmt.Println("\n--- STATISTIK PORTOFOLIO ---")
	fmt.Printf("Jumlah Jenis Saham: %d\n", data.JumlahPortofolio)
	fmt.Printf("Total Nilai Portofolio: Rp %.2f\n", totalValue)
	
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln()
	fmt.Scanln()
}
