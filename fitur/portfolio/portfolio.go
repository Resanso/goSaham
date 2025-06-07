package portfolio

import (
	"fmt"

	data "github.com/Resanso/goSaham/fitur/data"
)

// Menampilkan portofolio pengguna
func Portofolio() {
	if data.JumlahPortofolio == 0 {
		fmt.Println("Portofolio Anda kosong.")
	} else {
		fmt.Println("Portofolio Anda:")
		for i := 0; i < data.JumlahPortofolio; i++ {
			p := data.PortofolioUser[i]
			fmt.Printf("Kode Saham: %s, Jumlah: %d, Harga Beli: %.2f\n", p.KodeSaham, p.Jumlah, p.HargaBeli)
		}
	}
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}


