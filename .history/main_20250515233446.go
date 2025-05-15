package main

import "fmt"

// Struktur untuk data saham
type Saham struct {
	Kode     string
	Nama     string
	Harga    float64
	Kategori string
}

// Gunakan array dengan ukuran tetap 5 sebagai tipe data untuk saham
var listSaham = [5]Saham{
	{"BBCA", "Bank Central Asia", 9000.0, "Perbankan"},
	{"TLKM", "Telkom Indonesia", 3500.0, "Telekomunikasi"},
	{"ASII", "Astra International", 6800.0, "Otomotif"},
	{"UNVR", "Unilever Indonesia", 4500.0, "Konsumer"},
	{"GOTO", "GoTo Gojek Tokopedia", 75.0, "Teknologi"},
}

// Fungsi selection sort untuk mengurutkan saham berdasarkan harga tertinggi
func selectionSortByPrice(saham [5]Saham) [5]Saham {
	var result [5]Saham
	
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

// Fungsi insertion sort untuk mengurutkan saham berdasarkan harga tertinggi
func insertionSortByPrice(saham [5]Saham) [5]Saham {
	// Membuat array hasil dengan ukuran tetap 5
	var result [5]Saham
	
	// Salin data dari parameter ke array result
	for i := 0; i < 5; i++ {
		result[i] = saham[i]
	}
	
	// Mengurutkan dengan insertion sort
	for i := 1; i < 5; i++ {
		key := result[i]
		j := i - 1
		
		// Geser elemen yang lebih kecil dari key ke kanan
		for j >= 0 && result[j].Harga < key.Harga {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	
	return result
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
	running := true
	
	for running {
		fmt.Println("\n===== APLIKASI MANAJEMEN PORTOFOLIO SAHAM =====")
		fmt.Println("1. Portofolio Saya")
		fmt.Println("2. Lihat Daftar Saham")
		fmt.Println("3. Mencari Saham")
		fmt.Println("4. Transaksi Saham")
		fmt.Println("5. Lihat Saham Terurut (Selection Sort)")
		fmt.Println("6. Lihat Saham Terurut (Insertion Sort)")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu (1-7): ")
		
		fmt.Scan(&pilihan)
		
		switch pilihan {
		case 1:
			fmt.Println("\n--- PORTOFOLIO SAYA ---")
			Portofolio()
		case 2:
			fmt.Println("\n--- DAFTAR SAHAM ---")
			daftarSaham()
		case 3:
			fmt.Println("\n--- PENCARIAN SAHAM ---")
			cariSaham()
		case 4:
			fmt.Println("\n--- TRANSAKSI SAHAM ---")
			transaksi()
		case 5:
			fmt.Println("\n--- DAFTAR SAHAM TERURUT (SELECTION SORT) ---")
			daftarSahamTerurut("selection")
		case 6:
			fmt.Println("\n--- DAFTAR SAHAM TERURUT (INSERTION SORT) ---")
			daftarSahamTerurut("insertion")
		case 7:
			running = false
			keluar()
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1-7.")
		}
	}
}

func Portofolio() {
	// Alternatif untuk memeriksa portofolio kosong tanpa len()
	kosong := true
	
	// Coba akses elemen pertama
	if portofolioUser != nil {
		// Gunakan blok defer-recover untuk menghindari panic dari indeks di luar jangkauan
		func() {
			defer func() {
				if r := recover(); r != nil {
					// Panic terjadi, portofolio masih kosong
				}
			}()
			
			// Jika bisa mengakses elemen pertama, portofolio tidak kosong
			_ = portofolioUser[0]
			kosong = false
		}()
	}
	
	if kosong {
		fmt.Println("Portofolio Anda kosong.")
	} else {
		fmt.Println("Portofolio Anda:")
		
		// Iterasi portofolio sampai mendapatkan error akses indeks
		i := 0
		for {
			// Coba akses elemen di indeks i
			var p PortofolioItem
			valid := func() bool {
				defer func() {
					recover()
				}()
				if i >= 0 {
					p = portofolioUser[i]
					return true
				}
				return false
			}()
			
			if !valid {
				break
			}
			
			fmt.Printf("Kode Saham: %s, Jumlah: %d, Harga Beli: %.2f\n", p.KodeSaham, p.Jumlah, p.HargaBeli)
			i++
		}
	}
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}

// Update fungsi yang menggunakan listSaham untuk mengakses sebagai array, bukan slice
func daftarSaham() {
	fmt.Println("Daftar Saham:")
	for i := 0; i < 5; i++ {
		s := listSaham[i]
		fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
	}
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}

func cariSaham() {
	var kodeSaham string
	fmt.Print("Masukkan kode saham yang ingin dicari: ")
	fmt.Scan(&kodeSaham)

	portofolioUser = append(portofolioUser, PortofolioItem{KodeSaham: kodeSaham, Jumlah: jumlah, HargaBeli: hargaBeli})
	fmt.Println("Transaksi berhasil. Saham telah ditambahkan ke portofolio Anda.")for {
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
	fmt.Scanln() // Kadang perlu dua kali karena buffer input			found = true
}				break

func keluar() {
	fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")	} else {
}			break 

// Fungsi untuk menampilkan daftar saham terurut
func daftarSahamTerurut(metode string) {
	var sortedSaham [5]Saham
	var metodeName string	fmt.Println("Saham tidak ditemukan.")
	
	// Mengurutkan dengan metode yang dipilih
	if metode == "selection" {
		sortedSaham = selectionSortByPrice(listSaham) kembali ke menu utama...")
		metodeName = "Selection Sort"ter
	} else if metode == "insertion" {uffer input
		sortedSaham = insertionSortByPrice(listSaham)
		metodeName = "Insertion Sort"
	}unc transaksi() {
	
	fmt.Printf("Daftar Saham Terurut Berdasarkan Harga Tertinggi (%s):\n", metodeName)
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("%-6s %-25s %-10s %-15s\n", "Kode", "Nama", "Harga", "Kategori")
	fmt.Println("-----------------------------------------------------------")fmt.Print("Masukkan kode saham yang ingin dibeli: ")
	
	for i := 0; i < 5; i++ {jumlah saham yang ingin dibeli: ")
		s := sortedSaham[i]
		fmt.Printf("%-6s %-25s %-10.2f %-15s\n", s.Kode, s.Nama, s.Harga, s.Kategori)mt.Print("Masukkan harga beli per saham: ")
	}fmt.Scan(&hargaBeli)
	
	// Memberi jeda sebelum kembali ke menu utamadeSaham: kodeSaham, Jumlah: jumlah, HargaBeli: hargaBeli})
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")am telah ditambahkan ke portofolio Anda.")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input// Memberi jeda sebelum kembali ke menu utama

}	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}

func keluar() {
	fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
}

// Fungsi untuk menampilkan daftar saham terurut
func daftarSahamTerurut(metode string) {
	var sortedSaham [5]Saham
	var metodeName string
	
	// Mengurutkan dengan metode yang dipilih
	if metode == "selection" {
		sortedSaham = selectionSortByPrice(listSaham)
		metodeName = "Selection Sort"
	} else if metode == "insertion" {
		sortedSaham = insertionSortByPrice(listSaham)
		metodeName = "Insertion Sort"
	}
	
	fmt.Printf("Daftar Saham Terurut Berdasarkan Harga Tertinggi (%s):\n", metodeName)
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("%-6s %-25s %-10s %-15s\n", "Kode", "Nama", "Harga", "Kategori")
	fmt.Println("-----------------------------------------------------------")
	
	for i := 0; i < 5; i++ {
		s := sortedSaham[i]
		fmt.Printf("%-6s %-25s %-10.2f %-15s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
	}
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}