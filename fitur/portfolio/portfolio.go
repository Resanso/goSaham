package portfolio

import (
	"fmt"
	"tubesAlpro/fitur/data"
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

// Fungsi untuk menambah saham ke portofolio
func Transaksi() {
	var kodeSaham string
	var jumlah int
	var hargaBeli float64

	fmt.Print("Masukkan kode saham yang ingin dibeli: ")
	fmt.Scan(&kodeSaham)
	fmt.Print("Masukkan jumlah saham yang ingin dibeli: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan harga beli per saham: ")
	fmt.Scan(&hargaBeli)
	
	if data.JumlahPortofolio < 100 {
		data.PortofolioUser[data.JumlahPortofolio] = data.PortofolioItem{KodeSaham: kodeSaham, Jumlah: jumlah, HargaBeli: hargaBeli}
		data.JumlahPortofolio++
		fmt.Println("Transaksi berhasil. Saham telah ditambahkan ke portofolio Anda.")
	} else {
		fmt.Println("Portofolio penuh, tidak dapat menambahkan saham baru.")
	}
		
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() 
	fmt.Scanln() 
}
