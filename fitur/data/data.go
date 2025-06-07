package data

// Struktur untuk data saham
type Saham struct {
	Kode     string
	Nama     string
	Harga    float64
	Kategori string
	Volume   int
}

// Struktur untuk data portofolio
type PortofolioItem struct {
	KodeSaham string
	Jumlah    int
	HargaBeli float64
}

// Data saham yang tersedia (diurutkan berdasarkan nama perusahaan untuk binary search)
var ListSaham = [5]Saham{
	{"ASII", "Astra International", 6800.0, "Otomotif", 1250000},
	{"BBCA", "Bank Central Asia", 9000.0, "Perbankan", 21000},
	{"GOTO", "GoTo Gojek Tokopedia", 75.0, "Teknologi", 850000},
	{"TLKM", "Telkom Indonesia", 3500.0, "Telekomunikasi", 1750000},
	{"UNVR", "Unilever Indonesia", 4500.0, "Konsumer", 950000},
}

// Variabel global untuk menyimpan portofolio pengguna
var PortofolioUser [100]PortofolioItem
var JumlahPortofolio int = 0 // Untuk melacak jumlah item di portofolio
