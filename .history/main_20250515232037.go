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

// Fungsi selection sort untuk mengurutkan saham berdasarkan harga tertinggi
func selectionSortByPrice(saham []Saham) []Saham {
	result := []Saham{}
	
	// Mengisi result dengan saham dari parameter menggunakan append
	for i := 0; i < len(saham); i++ {
		result = append(result, saham[i])
	}
	
	// Mengurutkan dengan selection sort
	n := len(result) // len mengembalikan jumlah elemen dalam slice
	for i := 0; i < n-1; i++ {
		// Mencari indeks harga tertinggi di slice yang belum terurut
		maxIndex := i
		for j := i + 1; j < n; j++ {
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
func insertionSortByPrice(saham []Saham) []Saham {
	result := []Saham{}
	
	// Mengisi result dengan saham dari parameter menggunakan append
	for i := 0; i < len(saham); i++ {
		result = append(result, saham[i])
	}
	
	// Mengurutkan dengan insertion sort
	n := len(result) // len mengembalikan jumlah elemen dalam slice
	for i := 1; i < n; i++ {
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
	if len(portofolioUser) == 0 {
		fmt.Println("Portofolio Anda kosong.")
	} else {
		fmt.Println("Portofolio Anda:")
		for i := 0; i < len(portofolioUser); i++ {
			p := portofolioUser[i]
			fmt.Printf("Kode Saham: %s, Jumlah: %d, Harga Beli: %.2f\n", p.KodeSaham, p.Jumlah, p.HargaBeli)
		}
	}
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}

func daftarSaham() {
	fmt.Println("Daftar Saham:")
	for i := 0; i < len(listSaham); i++ {
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

	found := false
	i := 0
	for {
		if i >= 0 && i < 5 {
			s := listSaham[i]
			if s.Kode == kodeSaham {
				fmt.Printf("Kode: %s, Nama: %s, Harga: %.2f, Kategori: %s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
				found = true
				break
			}
			i++
		} else {
			break 
		}
	}

	if !found {
		fmt.Println("Saham tidak ditemukan.")
	}
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
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
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}

func keluar() {
	fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
}

// Fungsi untuk menampilkan daftar saham terurut
func daftarSahamTerurut(metode string) {
	var sortedSaham []Saham
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
	
	for i := 0; i < len(sortedSaham); i++ {
		s := sortedSaham[i]
		fmt.Printf("%-6s %-25s %-10.2f %-15s\n", s.Kode, s.Nama, s.Harga, s.Kategori)
	}
	
	// Memberi jeda sebelum kembali ke menu utama
	fmt.Print("\nTekan Enter untuk kembali ke menu utama...")
	fmt.Scanln() // Menunggu input Enter
	fmt.Scanln() // Kadang perlu dua kali karena buffer input
}