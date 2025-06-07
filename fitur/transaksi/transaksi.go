package transaksi

import (
	"fmt"

	data "github.com/Resanso/goSaham/fitur/data"
)

// TambahTransaksi menambah saham baru ke portofolio
func TambahTransaksi() {
	var kodeSaham string
	var jumlah int
	var hargaBeli float64

	fmt.Print("Masukkan kode saham: ")
	fmt.Scan(&kodeSaham)
	fmt.Print("Masukkan jumlah saham: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan harga beli: ")
	fmt.Scan(&hargaBeli)
	
	// Cek apakah masih bisa menambah saham
	if data.JumlahPortofolio < 100 {
		data.PortofolioUser[data.JumlahPortofolio] = data.PortofolioItem{
			KodeSaham: kodeSaham, 
			Jumlah: jumlah, 
			HargaBeli: hargaBeli,
		}
		data.JumlahPortofolio++
		fmt.Println("Saham berhasil ditambahkan!")
	} else {
		fmt.Println("Portofolio penuh!")
	}
}

// TampilkanDaftarTransaksi menampilkan semua transaksi yang ada
func TampilkanDaftarTransaksi() {
	fmt.Println("Daftar Transaksi:")
	fmt.Println("No. | Kode | Jumlah | Harga Beli")
	fmt.Println("----|----- |--------|----------")
	
	for i := 0; i < data.JumlahPortofolio; i++ {
		item := data.PortofolioUser[i]
		fmt.Printf("%d   | %s   | %d      | %.2f\n", 
			i+1, item.KodeSaham, item.Jumlah, item.HargaBeli)
	}
}

// EditTransaksi mengubah data transaksi yang sudah ada
func EditTransaksi() {
	if data.JumlahPortofolio == 0 {
		fmt.Println("Portofolio kosong!")
		return
	}

	TampilkanDaftarTransaksi()

	var nomor int
	fmt.Print("\nPilih nomor transaksi yang ingin diedit: ")
	fmt.Scan(&nomor)

	// Validasi input
	if nomor < 1 || nomor > data.JumlahPortofolio {
		fmt.Println("Nomor tidak valid!")
		return
	}

	index := nomor - 1
	item := data.PortofolioUser[index]
	
	fmt.Printf("Data sekarang: %s, %d saham, harga %.2f\n", 
		item.KodeSaham, item.Jumlah, item.HargaBeli)

	var kodeBaru string
	var jumlahBaru int
	var hargaBaru float64

	fmt.Print("Kode saham baru: ")
	fmt.Scan(&kodeBaru)
	fmt.Print("Jumlah baru: ")
	fmt.Scan(&jumlahBaru)
	fmt.Print("Harga beli baru: ")
	fmt.Scan(&hargaBaru)

	// Update data
	data.PortofolioUser[index] = data.PortofolioItem{
		KodeSaham: kodeBaru,
		Jumlah: jumlahBaru,
		HargaBeli: hargaBaru,
	}
	
	fmt.Println("Transaksi berhasil diupdate!")
}

// HapusTransaksi menghapus transaksi dari portofolio
func HapusTransaksi() {
	if data.JumlahPortofolio == 0 {
		fmt.Println("Portofolio kosong!")
		return
	}

	TampilkanDaftarTransaksi()

	var nomor int
	fmt.Print("\nPilih nomor transaksi yang ingin dihapus: ")
	fmt.Scan(&nomor)

	// Validasi input
	if nomor < 1 || nomor > data.JumlahPortofolio {
		fmt.Println("Nomor tidak valid!")
		return
	}

	index := nomor - 1
	item := data.PortofolioUser[index]
	
	var konfirmasi string
	fmt.Printf("Hapus %s (%d saham)? (y/n): ", item.KodeSaham, item.Jumlah)
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		// Geser elemen untuk menghapus
		for i := index; i < data.JumlahPortofolio-1; i++ {
			data.PortofolioUser[i] = data.PortofolioUser[i+1]
		}
		data.JumlahPortofolio--
		fmt.Println("Transaksi berhasil dihapus!")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

// MenuTransaksi menampilkan menu untuk transaksi saham
func MenuTransaksi() {
	for {
		fmt.Println("\n=== MENU TRANSAKSI ===")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. Edit Transaksi")
		fmt.Println("3. Hapus Transaksi")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih (1-4): ")
		
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahTransaksi()
		case 2:
			EditTransaksi()
		case 3:
			HapusTransaksi()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
