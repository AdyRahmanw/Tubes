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
	tanggal   string
	tipeCek   string
	deskripsi string
}

type arrKendaraan [NMAX]Kendaraan
type arrServis [NMAX]RiwayatServis

func ubahData(label string, data *string) {
	fmt.Printf("Masukkan %s Baru (saat ini: %s): ", label, *data)
	fmt.Scan(data)
}

func isiData(nKendaraan *int, tabKendaraan *arrKendaraan) {
	var plat string
	var i int

	fmt.Print("Masukan Plat Nomor Kendaraan (tanpa spasi, contoh: D1234AB): ")
	fmt.Scan(&plat)

	for i = 0; i < *nKendaraan; i++ {
		if tabKendaraan[i].plat == plat {
			fmt.Printf("\nKendaraan dengan plat '%s' sudah terdaftar sebagai %s!\n", plat, tabKendaraan[i].model)
			return
		}
	}

	tabKendaraan[*nKendaraan].plat = plat

	fmt.Print("Masukan Merk Mobil (tanpa spasi, contoh: Toyota_Avanza): ")
	fmt.Scan(&tabKendaraan[*nKendaraan].model)

	fmt.Print("Masukan Tahun Produksi Mobil: ")
	fmt.Scan(&tabKendaraan[*nKendaraan].tahun)

	*nKendaraan = *nKendaraan + 1
	fmt.Println("\nAkun kendaraan berhasil didaftarkan!")
}

func catatServis(nServis *int, tabServis *arrServis, platMobil string) {
	fmt.Println("------------------------------------------")
	fmt.Printf("--- Catat Servis Baru untuk %s ---\n", platMobil)

	tabServis[*nServis].platMobil = platMobil

	fmt.Print("Masukan tanggal servis (format YYYY-MM-DD, contoh: 2026-05-28): ")
	fmt.Scan(&tabServis[*nServis].tanggal)

	fmt.Print("Apa yang di cek: ")
	fmt.Scan(&tabServis[*nServis].tipeCek)

	fmt.Print("Tindakan: ")
	fmt.Scan(&tabServis[*nServis].deskripsi)

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
			fmt.Printf(" %d. Tanggal : %s\n", noServis, tabServis[j].tanggal)
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
	var temp Kendaraan

	for i = 1; i < nKendaraan; i++ {
		j = i
		for j > 0 && tabKendaraan[j-1].plat > tabKendaraan[j].plat {
			temp = tabKendaraan[j]
			tabKendaraan[j] = tabKendaraan[j-1]
			tabKendaraan[j-1] = temp
			j = j - 1
		}
	}
}

func selectionSortServis(nServis int, tabServis *arrServis, urutMenaik bool) {
	var i, j, indeksTarget int
	var temp RiwayatServis

	for i = 0; i < nServis-1; i++ {
		indeksTarget = i
		for j = i + 1; j < nServis; j++ {
			if urutMenaik {
				if tabServis[j].tanggal < tabServis[indeksTarget].tanggal {
					indeksTarget = j
				}
			} else {
				if tabServis[j].tanggal > tabServis[indeksTarget].tanggal {
					indeksTarget = j
				}
			}
		}
		temp = tabServis[indeksTarget]
		tabServis[indeksTarget] = tabServis[i]
		tabServis[i] = temp
	}
}

func binarySearch(tabKendaraan arrKendaraan, nKendaraan int, plat string) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = nKendaraan - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if tabKendaraan[tengah].plat == plat {
			return tengah
		} else if tabKendaraan[tengah].plat < plat {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func hapusKendaraan(tabKendaraan *arrKendaraan, nKendaraan *int, indeks int) {
	var i int

	if indeks < 0 || indeks >= *nKendaraan {
		return
	}

	for i = indeks; i < *nKendaraan-1; i++ {
		tabKendaraan[i] = tabKendaraan[i+1]
	}

	tabKendaraan[*nKendaraan-1] = Kendaraan{}
	*nKendaraan = *nKendaraan - 1
}

func hapusServis(tabServis *arrServis, nServis *int, indeks int) {
	var i int

	if indeks < 0 || indeks >= *nServis {
		return
	}

	for i = indeks; i < *nServis-1; i++ {
		tabServis[i] = tabServis[i+1]
	}

	tabServis[*nServis-1] = RiwayatServis{}
	*nServis = *nServis - 1
}

func cariIndeksServisKe(tabServis arrServis, nServis int, plat string, nomorKe int) int {
	var j, hitung, indeks int

	hitung = 0
	indeks = -1

	for j = 0; j < nServis && indeks == -1; j++ {
		if tabServis[j].platMobil == plat {
			hitung++
			if hitung == nomorKe {
				indeks = j
			}
		}
	}

	return indeks
}

func hapusServisPerPlat(tabServis *arrServis, nServis *int, plat string) {
	var i int

	i = 0
	for i < *nServis {
		if tabServis[i].platMobil == plat {
			hapusServis(tabServis, nServis, i)
		} else {
			i++
		}
	}
}

func main() {
	var pilihanUtama, pilihanSub, nKendaraan, nServis, indeksLogin, indeksHapus, pilihanUrut int
	var pilihanEdit, pilihanHapus, jumlahServis, indeksEdit int
	var tabKendaraan arrKendaraan
	var tabServis arrServis
	var platLogin, platHapus string
	var keluarMenuSession, keluarSub bool

	nKendaraan = 5
	nServis = 6

	tabKendaraan[0] = Kendaraan{plat: "D1234AB", model: "Toyota_Avanza", tahun: "2018"}
	tabKendaraan[1] = Kendaraan{plat: "B9999XYZ", model: "Honda_Civic", tahun: "2020"}
	tabKendaraan[2] = Kendaraan{plat: "Z7777AA", model: "Suzuki_Ertiga", tahun: "2019"}
	tabKendaraan[3] = Kendaraan{plat: "F5555BB", model: "Mitsubishi_Pajero", tahun: "2021"}
	tabKendaraan[4] = Kendaraan{plat: "DK1111CC", model: "Daihatsu_Xenia", tahun: "2022"}

	insertionSortKendaraan(nKendaraan, &tabKendaraan)

	tabServis[0] = RiwayatServis{platMobil: "D1234AB", tanggal: "2025-05-12", tipeCek: "Cek_Mesin", deskripsi: "Ganti_Oli_dan_Filter"}
	tabServis[1] = RiwayatServis{platMobil: "D1234AB", tanggal: "2025-10-10", tipeCek: "Cek_Rem", deskripsi: "Ganti_Kampas_Rem"}
	tabServis[2] = RiwayatServis{platMobil: "B9999XYZ", tanggal: "2026-01-01", tipeCek: "Cek_Aki", deskripsi: "Ganti_Aki_GS_Astra"}
	tabServis[3] = RiwayatServis{platMobil: "Z7777AA", tanggal: "2025-03-15", tipeCek: "Cek_AC", deskripsi: "Pengisian_Freon_AC"}
	tabServis[4] = RiwayatServis{platMobil: "Z7777AA", tanggal: "2025-09-20", tipeCek: "Cek_Roda", deskripsi: "Spooring_dan_Balancing"}
	tabServis[5] = RiwayatServis{platMobil: "F5555BB", tanggal: "2026-04-20", tipeCek: "Cek_Mesin", deskripsi: "Tune_up_rutin"}

	for {
		fmt.Println("=== Aplikasi manajemen kendaraan ===")
		fmt.Println("1. Login")
		fmt.Println("2. Registrasi Kendaraan Baru")
		fmt.Println("3. Lihat Seluruh Kendaraan Terdaftar")
		fmt.Println("4. Hapus Kendaraan")
		fmt.Println("5. Keluar Aplikasi")
		fmt.Print("Pilih menu (1-5): ")
		fmt.Scan(&pilihanUtama)

		switch pilihanUtama {
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
					fmt.Scan(&pilihanSub)

					switch pilihanSub {
					case 1:
						keluarSub = false
						for !keluarSub {
							tampilProfil(indeksLogin, tabKendaraan)
							tampilServis(tabKendaraan[indeksLogin].plat, nServis, tabServis)

							fmt.Println("\nPilih Urutan Riwayat Servis:")
							fmt.Println("1. Terlama ke Terbaru (Ascending)")
							fmt.Println("2. Terbaru ke Terlama (Descending)")
							fmt.Println("3. Kembali ke Menu Mobil")
							fmt.Print("Pilih (1-3): ")
							fmt.Scan(&pilihanUrut)

							if pilihanUrut == 1 || pilihanUrut == 2 {
								selectionSortServis(nServis, &tabServis, pilihanUrut == 1)
							} else if pilihanUrut == 3 {
								keluarSub = true
							} else {
								fmt.Println("Pilihan tidak valid.")
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
								indeksEdit = cariIndeksServisKe(tabServis, nServis, platLogin, pilihanEdit)
								fmt.Println("\n--- Edit Riwayat Servis ---")
								ubahData("Tanggal Servis", &tabServis[indeksEdit].tanggal)
								ubahData("Tipe Cek", &tabServis[indeksEdit].tipeCek)
								ubahData("Tindakan/Keterangan", &tabServis[indeksEdit].deskripsi)

								fmt.Println("\nData riwayat servis berhasil diperbarui!")
								fmt.Println("Riwayat servis terbaru:")
								tampilServis(platLogin, nServis, tabServis)
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
								indeksHapus = cariIndeksServisKe(tabServis, nServis, platLogin, pilihanHapus)
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
				hapusServisPerPlat(&tabServis, &nServis, platHapus)
				fmt.Printf("Kendaraan %s dan semua riwayat servisnya berhasil dihapus.\n", platHapus)
			}

		case 5:
			return

		default:
			fmt.Println("\nOpsi tidak valid")
		}
	}
}
