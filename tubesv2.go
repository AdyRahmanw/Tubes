package main

import "fmt"

const NMAX = 100

// 1. Struct Profil Kendaraan (Tanpa data servis tunggal)
type Kendaraan struct {
	plat  string
	model string
	tahun string
}

// 2. Struct Transaksi Servis Baru
type RiwayatServis struct {
	platMobil string // Hubungan relasi ke Kendaraan.plat
	tanggal   string
	tipeCek   string
	deskripsi string
}

// Deklarasi 2 jenis array
type arrKendaraan [NMAX]Kendaraan
type arrServis [NMAX]RiwayatServis

// ==========================================
// FUNGSI 1: PENDAFTARAN KENDARAAN (TAMBAH AKUN)
// ==========================================
func isiData(nKendaraan *int, tabKendaraan *arrKendaraan) {
	// Deklarasi semua variabel di awal fungsi
	var plat, model, tahun string
	var i int

	if *nKendaraan >= NMAX {
		fmt.Println("\n[ERROR] Kapasitas memori pendaftaran kendaraan penuh!")
		return
	}

	fmt.Print("Masukan Plat Nomor Kendaraan (tanpa spasi, contoh: D1234AB): ")
	fmt.Scan(&plat)

	// Validasi keunikan plat nomor (pencarian linier)
	for i = 0; i < *nKendaraan; i++ {
		if tabKendaraan[i].plat == plat {
			fmt.Printf("\n[ERROR] Kendaraan dengan plat '%s' sudah terdaftar sebagai %s!\n", plat, tabKendaraan[i].model)
			return
		}
	}

	fmt.Print("Masukan Merk Mobil (tanpa spasi, contoh: Toyota_Avanza): ")
	fmt.Scan(&model)

	fmt.Print("Masukan Tahun Produksi Mobil: ")
	fmt.Scan(&tahun)

	// Menyimpan data secara manual ke array
	tabKendaraan[*nKendaraan].plat = plat
	tabKendaraan[*nKendaraan].model = model
	tabKendaraan[*nKendaraan].tahun = tahun
	*nKendaraan = *nKendaraan + 1
	fmt.Println("\n[SUKSES] Akun kendaraan baru berhasil didaftarkan!")
}

// ==========================================
// FUNGSI 2: PENCATATAN TRANSAKSI SERVIS BARU (LOGGED-IN SESSION)
// ==========================================
// Parameter tidak lagi mencari plat nomor, melainkan langsung menggunakan plat login aktif
func catatServis(nServis *int, tabServis *arrServis, platMobil string) {
	// Deklarasi semua variabel di awal fungsi
	var tanggal, tipeCek, deskripsi string

	if *nServis >= NMAX {
		fmt.Println("\n[ERROR] Kapasitas memori catatan servis bengkel penuh!")
		return
	}

	fmt.Printf("\n--- Catat Servis Baru untuk %s ---\n", platMobil)

	fmt.Print("Masukan tanggal servis (contoh: 28-Mei-2026): ")
	fmt.Scan(&tanggal)

	fmt.Print("Masukan tipe pengecekan (tanpa spasi, contoh: Ganti_Oli): ")
	fmt.Scan(&tipeCek)

	fmt.Print("Masukan tindakan/keterangan (tanpa spasi, contoh: Ganti_oli_mesin): ")
	fmt.Scan(&deskripsi)

	// Menyimpan catatan servis ke array servis atas nama plat login aktif
	tabServis[*nServis].platMobil = platMobil
	tabServis[*nServis].tanggal = tanggal
	tabServis[*nServis].tipeCek = tipeCek
	tabServis[*nServis].deskripsi = deskripsi
	*nServis = *nServis + 1
	fmt.Println("\n[SUKSES] Catatan transaksi servis berhasil ditambahkan!")
}

// ==========================================
// FUNGSI 3: TAMPILKAN SELURUH MOBIL YANG TERDAFTAR (ADMIN VIEW)
// ==========================================
func tampilData(nKendaraan int, tabKendaraan arrKendaraan) {
	// Deklarasi variabel loop di awal fungsi
	var i int

	if nKendaraan == 0 {
		fmt.Println("\nBelum ada data kendaraan yang terdaftar.")
		return
	}

	fmt.Println("\n--- Daftar Seluruh Kendaraan ---")
	for i = 0; i < nKendaraan; i++ {
		fmt.Printf("\nData ke-%d\n", i+1)
		fmt.Printf("Plat Nomor      : %s\n", tabKendaraan[i].plat)
		fmt.Printf("Model Kendaraan : %s\n", tabKendaraan[i].model)
		fmt.Printf("Tahun Produksi  : %s\n", tabKendaraan[i].tahun)
	}
	fmt.Println("--------------------------------")
}

// ==========================================
// FUNGSI 4: TAMPILKAN PROFIL & RIWAYAT SERVIS SAYA (LOGGED-IN SESSION)
// ==========================================
// Menampilkan histori pelayanan servis langsung dari mobil yang sedang aktif login
func tampilProfilDanServis(indeksMobil int, tabKendaraan arrKendaraan, nServis int, tabServis arrServis) {
	// Deklarasi semua variabel di awal fungsi
	var j, noServis int
	var plat string

	plat = tabKendaraan[indeksMobil].plat

	// Cetak Profil Kendaraan
	fmt.Println("\n========================================")
	fmt.Println("         PROFIL KENDARAAN (AKUN)        ")
	fmt.Println("========================================")
	fmt.Printf("Plat Nomor      : %s\n", plat)
	fmt.Printf("Model Kendaraan : %s\n", tabKendaraan[indeksMobil].model)
	fmt.Printf("Tahun Produksi  : %s\n", tabKendaraan[indeksMobil].tahun)
	fmt.Println("----------------------------------------")
	fmt.Println("             RIWAYAT SERVIS             ")
	fmt.Println("----------------------------------------")

	noServis = 1
	// Loop relasional: mencari log di array servis yang platMobil-nya cocok
	for j = 0; j < nServis; j++ {
		if tabServis[j].platMobil == plat {
			fmt.Printf(" [%d] Tanggal : %s\n", noServis, tabServis[j].tanggal)
			fmt.Printf("     Cek     : %s\n", tabServis[j].tipeCek)
			fmt.Printf("     Tindakan: %s\n", tabServis[j].deskripsi)
			fmt.Println("  --------------------------------------")
			noServis = noServis + 1
		}
	}

	if noServis == 1 {
		fmt.Println("  Belum ada riwayat servis untuk kendaraan ini.")
	}
	fmt.Println("========================================")
}

// ==========================================
// FUNGSI UTAMA (MAIN PROGRAM)
// ==========================================
func main() {
	// Deklarasi semua variabel di awal fungsi main
	var pilihan int
	var nKendaraan int
	var nServis int
	var tabKendaraan arrKendaraan
	var tabServis arrServis
	var platLogin string
	var indeksLogin, i int
	var keluarMenuSession bool

	nKendaraan = 5
	nServis = 6

	// Mock Data Kendaraan bawaan
	tabKendaraan[0] = Kendaraan{plat: "D1234AB", model: "Toyota_Avanza", tahun: "2018"}
	tabKendaraan[1] = Kendaraan{plat: "B9999XYZ", model: "Honda_Civic", tahun: "2020"}
	tabKendaraan[2] = Kendaraan{plat: "Z7777AA", model: "Suzuki_Ertiga", tahun: "2019"}
	tabKendaraan[3] = Kendaraan{plat: "F5555BB", model: "Mitsubishi_Pajero", tahun: "2021"}
	tabKendaraan[4] = Kendaraan{plat: "DK1111CC", model: "Daihatsu_Xenia", tahun: "2022"}

	// Mock Data Servis bawaan (relasi)
	tabServis[0] = RiwayatServis{platMobil: "D1234AB", tanggal: "12-Mei-2025", tipeCek: "Cek_Mesin", deskripsi: "Ganti_Oli_dan_Filter"}
	tabServis[1] = RiwayatServis{platMobil: "D1234AB", tanggal: "10-Oktober-2025", tipeCek: "Cek_Rem", deskripsi: "Ganti_Kampas_Rem"}
	tabServis[2] = RiwayatServis{platMobil: "B9999XYZ", tanggal: "01-Januari-2026", tipeCek: "Cek_Aki", deskripsi: "Ganti_Aki_GS_Astra"}
	tabServis[3] = RiwayatServis{platMobil: "Z7777AA", tanggal: "15-Maret-2025", tipeCek: "Cek_AC", deskripsi: "Pengisian_Freon_AC"}
	tabServis[4] = RiwayatServis{platMobil: "Z7777AA", tanggal: "20-September-2025", tipeCek: "Cek_Roda", deskripsi: "Spooring_dan_Balancing"}
	tabServis[5] = RiwayatServis{platMobil: "F5555BB", tanggal: "20-April-2026", tipeCek: "Cek_Mesin", deskripsi: "Tune_up_rutin"}

	for {
		// Halaman Awal Bengkel (Guest View)
		fmt.Println("\n=== SISTEM BENGKEL (HALAMAN UTAMA) ===")
		fmt.Println("1. Login Pelanggan (Masuk Akun)")
		fmt.Println("2. Registrasi Kendaraan Baru (Daftar Akun)")
		fmt.Println("3. Lihat Seluruh Kendaraan Terdaftar (Admin View)")
		fmt.Println("4. Keluar Aplikasi")
		fmt.Print("Pilih menu (1-4): ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("\nMasukan Plat Nomor Kendaraan Anda: ")
			fmt.Scan(&platLogin)

			// Proses Login: cari plat nomor di database
			indeksLogin = -1
			for i = 0; i < nKendaraan; i++ {
				if tabKendaraan[i].plat == platLogin {
					indeksLogin = i
				}
			}

			if indeksLogin == -1 {
				fmt.Println("\n[ERROR] Plat nomor belum terdaftar! Silakan lakukan Registrasi di Menu 2 terlebih dahulu.")
			} else {
				fmt.Printf("\n[LOGIN BERHASIL] Selamat datang, Pemilik %s (%s)!\n", tabKendaraan[indeksLogin].model, platLogin)

				// Loop Sesi Aktif Pelanggan (Logged-in View)
				keluarMenuSession = false
				for !keluarMenuSession {
					fmt.Printf("\n=== MENU MOBIL ANDA (%s) ===\n", platLogin)
					fmt.Println("1. Lihat Profil & Riwayat Servis Saya")
					fmt.Println("2. Catat Pelayanan Servis Baru")
					fmt.Println("3. Logout (Keluar Sesi)")
					fmt.Print("Pilih menu (1-3): ")

					fmt.Scan(&pilihan)

					switch pilihan {
					case 1:
						// Langsung menampilkan data tanpa menanyakan plat nomor lagi
						tampilProfilDanServis(indeksLogin, tabKendaraan, nServis, tabServis)
					case 2:
						// Langsung mencatat servis menggunakan plat session saat ini
						catatServis(&nServis, &tabServis, platLogin)
					case 3:
						fmt.Println("\n[LOGOUT] Keluar sesi akun berhasil.")
						keluarMenuSession = true
					default:
						fmt.Println("\n[ERROR] Opsi tidak valid! Ketik angka 1 sampai 3.")
					}
				}
			}
		case 2:
			isiData(&nKendaraan, &tabKendaraan)
		case 3:
			tampilData(nKendaraan, tabKendaraan)
		case 4:
			fmt.Println("\nTerima kasih telah menggunakan sistem bengkel kami!")
			return
		default:
			fmt.Println("\n[ERROR] Opsi tidak valid! Silakan masukkan angka 1 sampai 4.")
		}
	}
}
