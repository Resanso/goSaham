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
		fmt.Println("Portofolio Anda:")
		for _, p := range portofolioUser {
			fmt.Printf("Kode Saham: %s, Jumlah: %d, Harga Beli: %.2f\n", p.KodeSaham, p.Jumlah, p.HargaBeli)
		}
	}
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