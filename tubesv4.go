package main

import "fmt"

const NMAX = 100

type Kendaraan struct {
	plat  string
	model string
	tahun string
}

type RiwayatServis struct {
	platMobil string
	hari      int
	bulan     int
	tahun     int
	tipeCek   string
	deskripsi string
}

type arrKendaraan [NMAX]Kendaraan
type arrServis [NMAX]RiwayatServis

func editData(label string, data *string) {
	fmt.Printf("Masukkan %s Baru (saat ini: %s): ", label, *data)
	fmt.Scan(data)
}

func editDataInt(label string, data *int) {
	fmt.Printf("Masukkan %s Baru (saat ini: %d): ", label, *data)
	fmt.Scan(data)
}

func isiData(nKendaraan *int, tabKendaraan *arrKendaraan) {
	var plat, model, tahun string
	var i int

	fmt.Print("Masukan Plat Nomor Kendaraan (tanpa spasi, contoh: D1234AB): ")
	fmt.Scan(&plat)

	for i = 0; i < *nKendaraan; i++ {
		if tabKendaraan[i].plat == plat {
			fmt.Printf("\nKendaraan dengan plat '%s' sudah terdaftar sebagai %s!\n", plat, tabKendaraan[i].model)
			return
		}
	}

	fmt.Print("Masukan Merk Mobil (tanpa spasi, contoh: Toyota_Avanza): ")
	fmt.Scan(&model)

	fmt.Print("Masukan Tahun Produksi Mobil: ")
	fmt.Scan(&tahun)

	tabKendaraan[*nKendaraan].plat = plat
	tabKendaraan[*nKendaraan].model = model
	tabKendaraan[*nKendaraan].tahun = tahun
	*nKendaraan = *nKendaraan + 1
	
	fmt.Println("\nAkun kendaraan berhasil didaftarkan!")
}

func catatServis(nServis *int, tabServis *arrServis, platMobil string) {
	var hari, bulan, tahun int
	var tipeCek, deskripsi string

	fmt.Println("------------------------------------------")
	fmt.Printf("--- Catat Servis Baru untuk %s ---\n", platMobil)

	fmt.Print("Masukan hari (contoh: 28): ")
	fmt.Scan(&hari)
	fmt.Print("Masukan bulan (contoh: 5): ")
	fmt.Scan(&bulan)
	fmt.Print("Masukan tahun (contoh: 2026): ")
	fmt.Scan(&tahun)
	
	fmt.Print("Apa yang di cek: ")
	fmt.Scan(&tipeCek)
	
	fmt.Print("Tindakan: ")
	fmt.Scan(&deskripsi)

	tabServis[*nServis].platMobil = platMobil
	tabServis[*nServis].hari = hari
	tabServis[*nServis].bulan = bulan
	tabServis[*nServis].tahun = tahun
	tabServis[*nServis].tipeCek = tipeCek
	tabServis[*nServis].deskripsi = deskripsi
	*nServis = *nServis + 1
	
	fmt.Println("\nCatatan servis berhasil ditambahkan!")
}

func tampilData(nKendaraan int, tabKendaraan arrKendaraan) {
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

func tampilProfil(indeksMobil int, tabKendaraan arrKendaraan) {
	fmt.Println("            PROFIL KENDARAAN              ")
	fmt.Println("-------------------------------------------")
	fmt.Printf("Plat Nomor      : %s\n", tabKendaraan[indeksMobil].plat)
	fmt.Printf("Model Kendaraan : %s\n", tabKendaraan[indeksMobil].model)
	fmt.Printf("Tahun Produksi  : %s\n", tabKendaraan[indeksMobil].tahun)
	fmt.Println("-------------------------------------------")
}

func tampilServis(plat string, nServis int, tabServis arrServis) int {
	var j, noServis int

	fmt.Println("             RIWAYAT SERVIS             ")
	fmt.Println("-------------------------------------------")
	noServis = 1
	for j = 0; j < nServis; j++ {
		if tabServis[j].platMobil == plat {
			fmt.Printf(" %d. Tanggal : %d-%d-%d\n", noServis, tabServis[j].hari, tabServis[j].bulan, tabServis[j].tahun)
			fmt.Printf("     Cek     : %s\n", tabServis[j].tipeCek)
			fmt.Printf("     Tindakan: %s\n", tabServis[j].deskripsi)
			fmt.Println("------------------------------------------")
			noServis = noServis + 1
		}
	}

	if noServis == 1 {
		fmt.Println(" Belum ada riwayat servis untuk kendaraan ini.")
	}
	return noServis - 1
}

func insertionSortKendaraan(nKendaraan int, tabKendaraan *arrKendaraan) {
	var i, j int
	var key Kendaraan
	
	for i = 1; i < nKendaraan; i++ {
		key = tabKendaraan[i]
		j = i - 1
		
		for j >= 0 && tabKendaraan[j].plat > key.plat {
			tabKendaraan[j+1] = tabKendaraan[j]
			j = j - 1
		}
		
		tabKendaraan[j+1] = key
	}
}

func lebihAwal(s1, s2 RiwayatServis) bool {
	if s1.tahun != s2.tahun {
		return s1.tahun < s2.tahun
	}
	if s1.bulan != s2.bulan {
		return s1.bulan < s2.bulan
	}
	return s1.hari < s2.hari
}

func selectionSortServisAsc(tabServis *arrServis, nServis int) {
	var i, j, idxMin int
	var temp RiwayatServis
	for i = 0; i < nServis-1; i++ {
		idxMin = i
		for j = i + 1; j < nServis; j++ {
			if lebihAwal(tabServis[j], tabServis[idxMin]) {
				idxMin = j
			}
		}
		temp = tabServis[i]
		tabServis[i] = tabServis[idxMin]
		tabServis[idxMin] = temp
	}
}

func selectionSortServisDesc(tabServis *arrServis, nServis int) {
	var i, j, idxMax int
	var temp RiwayatServis
	for i = 0; i < nServis-1; i++ {
		idxMax = i
		for j = i + 1; j < nServis; j++ {
			if lebihAwal(tabServis[idxMax], tabServis[j]) {
				idxMax = j
			}
		}
		temp = tabServis[i]
		tabServis[i] = tabServis[idxMax]
		tabServis[idxMax] = temp
	}
}

func binarySearch(tab arrKendaraan, n int, x string) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = n - 1
	
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		
		if tab[tengah].plat == x {
			return tengah
		} else if tab[tengah].plat < x {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	
	return -1
}

func hapusKendaraan(tab *arrKendaraan, n *int, indeks int) {
	var i int
	
	if indeks < 0 || indeks >= *n {
		return
	}
	
	for i = indeks; i < *n-1; i++ {
		tab[i] = tab[i+1]
	}
	
	tab[*n-1] = Kendaraan{}
	*n = *n - 1
}

func hapusServis(tab *arrServis, n *int, indeks int) {
	var i int
	
	if indeks < 0 || indeks >= *n {
		return
	}
	
	for i = indeks; i < *n-1; i++ {
		tab[i] = tab[i+1]
	}
	
	tab[*n-1] = RiwayatServis{}
	*n = *n - 1
}

func main() {
	var pilihan, nKendaraan, nServis, indeksLogin, indeksHapus, i int
	var pilihanEdit, pilihanHapus, jumlahServis, count, indeksEdit, j int
	var tabKendaraan arrKendaraan
	var tabServis arrServis
	var platLogin, platHapus string
	var keluarMenuSession bool

	nKendaraan = 5
	nServis = 6

	tabKendaraan[0] = Kendaraan{plat: "D1234AB", model: "Toyota_Avanza", tahun: "2018"}
	tabKendaraan[1] = Kendaraan{plat: "B9999XYZ", model: "Honda_Civic", tahun: "2020"}
	tabKendaraan[2] = Kendaraan{plat: "Z7777AA", model: "Suzuki_Ertiga", tahun: "2019"}
	tabKendaraan[3] = Kendaraan{plat: "F5555BB", model: "Mitsubishi_Pajero", tahun: "2021"}
	tabKendaraan[4] = Kendaraan{plat: "DK1111CC", model: "Daihatsu_Xenia", tahun: "2022"}

	insertionSortKendaraan(nKendaraan, &tabKendaraan)

	tabServis[0] = RiwayatServis{platMobil: "D1234AB", hari: 12, bulan: 5, tahun: 2025, tipeCek: "Cek_Mesin", deskripsi: "Ganti_Oli_dan_Filter"}
	tabServis[1] = RiwayatServis{platMobil: "D1234AB", hari: 10, bulan: 10, tahun: 2025, tipeCek: "Cek_Rem", deskripsi: "Ganti_Kampas_Rem"}
	tabServis[2] = RiwayatServis{platMobil: "B9999XYZ", hari: 1, bulan: 1, tahun: 2026, tipeCek: "Cek_Aki", deskripsi: "Ganti_Aki_GS_Astra"}
	tabServis[3] = RiwayatServis{platMobil: "Z7777AA", hari: 15, bulan: 3, tahun: 2025, tipeCek: "Cek_AC", deskripsi: "Pengisian_Freon_AC"}
	tabServis[4] = RiwayatServis{platMobil: "Z7777AA", hari: 20, bulan: 9, tahun: 2025, tipeCek: "Cek_Roda", deskripsi: "Spooring_dan_Balancing"}
	tabServis[5] = RiwayatServis{platMobil: "F5555BB", hari: 20, bulan: 4, tahun: 2026, tipeCek: "Cek_Mesin", deskripsi: "Tune_up_rutin"}

	for {
		fmt.Println("=== Aplikasi manajemen kendaraan ===")
		fmt.Println("1. Login")
		fmt.Println("2. Registrasi Kendaraan Baru")
		fmt.Println("3. Lihat Seluruh Kendaraan Terdaftar")
		fmt.Println("4. Hapus Kendaraan")
		fmt.Println("5. Keluar Aplikasi")
		fmt.Print("Pilih menu (1-5): ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukan Plat Nomor Kendaraan Anda: ")
			fmt.Scan(&platLogin)

			indeksLogin = binarySearch(tabKendaraan, nKendaraan, platLogin)

			if indeksLogin == -1 {
				fmt.Println("Plat nomor tidak terdaftar.")
			} else {
				keluarMenuSession = false
				for !keluarMenuSession {
					fmt.Println("------------------------------------------")
					fmt.Printf("=== MENU MOBIL(%s) ===\n", platLogin)
					fmt.Println("1. Lihat Profil & Riwayat Servis")
					fmt.Println("2. Catat Pelayanan Servis Baru")
					fmt.Println("3. Edit Riwayat Servis")
					fmt.Println("4. Hapus Riwayat Servis Kendaraan Ini")
					fmt.Println("5. Logout")
					fmt.Print("Pilih menu (1-5): ")

					fmt.Scan(&pilihan)

					switch pilihan {
					case 1:
						var keluarSubMenu bool
						keluarSubMenu = false
						for !keluarSubMenu {
							tampilProfil(indeksLogin, tabKendaraan)
							tampilServis(tabKendaraan[indeksLogin].plat, nServis, tabServis)
							
							fmt.Println("\nOpsi:")
							fmt.Println("1. Urutkan Riwayat Servis (Ascending - Terlama ke Terbaru)")
							fmt.Println("2. Urutkan Riwayat Servis (Descending - Terbaru ke Terlama)")
							fmt.Println("3. Kembali")
							fmt.Print("Pilih menu (1-3): ")
							
							var pilihanUrut int
							fmt.Scan(&pilihanUrut)
							
							if pilihanUrut == 1 {
								selectionSortServisAsc(&tabServis, nServis)
								fmt.Println("\nRiwayat servis berhasil diurutkan secara Ascending!")
							} else if pilihanUrut == 2 {
								selectionSortServisDesc(&tabServis, nServis)
								fmt.Println("\nRiwayat servis berhasil diurutkan secara Descending!")
							} else if pilihanUrut == 3 {
								keluarSubMenu = true
							} else {
								fmt.Println("Opsi tidak valid.")
							}
						}
						
					case 2:
						catatServis(&nServis, &tabServis, platLogin)
						
					case 3:
						fmt.Println("\n--- Pilih Riwayat Servis yang Akan Diedit ---")
						jumlahServis = tampilServis(platLogin, nServis, tabServis)
						if jumlahServis > 0 {
							fmt.Printf("Pilih nomor riwayat servis yang akan diedit (1-%d): ", jumlahServis)
							fmt.Scan(&pilihanEdit)
							if pilihanEdit >= 1 && pilihanEdit <= jumlahServis {
								count = 0
								indeksEdit = -1
								for j = 0; j < nServis && indeksEdit == -1; j++ {
									if tabServis[j].platMobil == platLogin {
										count++
										if count == pilihanEdit {
											indeksEdit = j
										}
									}
								}
								fmt.Println("\n--- Edit Riwayat Servis ---")
								editDataInt("Hari Servis", &tabServis[indeksEdit].hari)
								editDataInt("Bulan Servis", &tabServis[indeksEdit].bulan)
								editDataInt("Tahun Servis", &tabServis[indeksEdit].tahun)
								editData("Tipe Cek", &tabServis[indeksEdit].tipeCek)
								editData("Tindakan/Keterangan", &tabServis[indeksEdit].deskripsi)
								
								fmt.Println("\nData riwayat servis berhasil diperbarui!")
								fmt.Println("Riwayat servis terbaru:")
							} else {
								fmt.Println("Pilihan tidak valid.")
							}
						}
						
					case 4:
						fmt.Println("\n--- Pilih Riwayat Servis yang Akan Dihapus ---")
						jumlahServis = tampilServis(platLogin, nServis, tabServis)
						if jumlahServis > 0 {
							fmt.Printf("Pilih nomor riwayat servis yang akan dihapus (1-%d): ", jumlahServis)
							fmt.Scan(&pilihanHapus)
							if pilihanHapus >= 1 && pilihanHapus <= jumlahServis {
								count = 0
								indeksHapus = -1
								for j = 0; j < nServis && indeksHapus == -1; j++ {
									if tabServis[j].platMobil == platLogin {
										count++
										if count == pilihanHapus {
											indeksHapus = j
										}
									}
								}
								hapusServis(&tabServis, &nServis, indeksHapus)
								fmt.Println("Riwayat servis berhasil dihapus!")
								fmt.Println("\nRiwayat servis terbaru:")
								tampilServis(platLogin, nServis, tabServis)
							} else {
								fmt.Println("Pilihan tidak valid.")
							}
						}
						
					case 5:
						fmt.Println("Keluar sesi akun berhasil.")
						keluarMenuSession = true
					default:
						fmt.Println("Opsi tidak valid.")
					}
				}
			}
			
		case 2:
			isiData(&nKendaraan, &tabKendaraan)
			insertionSortKendaraan(nKendaraan, &tabKendaraan)
			
		case 3:
			tampilData(nKendaraan, tabKendaraan)
			
		case 4:
			fmt.Print("Masukan Plat Nomor Kendaraan yang akan dihapus: ")
			fmt.Scan(&platHapus)
			indeksHapus = binarySearch(tabKendaraan, nKendaraan, platHapus)
			if indeksHapus == -1 {
				fmt.Println("Kendaraan dengan plat tersebut tidak ditemukan.")
			} else {
				fmt.Println("\nDetail Kendaraan yang akan dihapus:")
				tampilProfil(indeksHapus, tabKendaraan)
				tampilServis(platHapus, nServis, tabServis)

				hapusKendaraan(&tabKendaraan, &nKendaraan, indeksHapus)

				i = 0
				for i < nServis {
					if tabServis[i].platMobil == platHapus {
						hapusServis(&tabServis, &nServis, i)
					} else {
						i++
					}
				}
				fmt.Printf("Kendaraan %s dan semua riwayat servisnya berhasil dihapus.\n", platHapus)
			}
			
		case 5:
			return
		default:
			fmt.Println("\nOpsi tidak valid")
		}
	}
}
