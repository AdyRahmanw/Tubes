package main

import "fmt"

const NMAX = 100

type Kendaraan struct {
	plat   string
	model  string
	tahun  string
	servis string
}

type arrData [NMAX]Kendaraan

func isiData(n *int, tabData *arrData) {
	if *n >= 100 {
		fmt.Println("Kapasitas memori penuh")
	}else{
		fmt.Printf("masukan plat nomor (tanpa spasi): ")
		fmt.Scan(&tabData[*n].plat)
		fmt.Printf("masukan Merk Mobil (tanpa spasi): ")
		fmt.Scan(&tabData[*n].model)
		fmt.Printf("masukan Tahun Produksi Mobil: ")
		fmt.Scan(&tabData[*n].tahun)
		fmt.Printf("Masukan kapan terakhir kali service (tanpa spasi): ")
		fmt.Scan(&tabData[*n].servis)
	}
	*n++ 
}

func cariData(n int, cariPlat string, tabData arrData) {
	var i int

	for i = 0; i < n; i++ {
		if tabData[i].plat == cariPlat {
			fmt.Printf("plat nomor : %s\n", tabData[i].plat)
			fmt.Printf("model kendaraan: %s\n", tabData[i].model)
			fmt.Printf("tahun kendaraan: %s\n", tabData[i].tahun)
			fmt.Printf("Riwayat servis: %s\n", tabData[i].servis)
		}
	}
}

func tampilData(n int, tabData arrData) {
	var i int

	if n == 0 {
		fmt.Printf("\nBelum ada data kendaraan yang terdaftar.\n")
	}else{
		fmt.Printf("\n--- Daftar Seluruh Kendaraan ---\n")
		for i = 0; i < n; i++ {
			fmt.Printf("\nData ke-%d\n", i+1)
			fmt.Printf("Plat Nomor      : %s\n", tabData[i].plat)
			fmt.Printf("Model Kendaraan : %s\n", tabData[i].model)
			fmt.Printf("Tahun Produksi  : %s\n", tabData[i].tahun)
			fmt.Printf("Terakhir Servis : %s\n", tabData[i].servis)
		}
	}
	fmt.Println("--------------------------------")
}

func main() {
	var targetInput int
	var jumlahSekarang int = 5 
	var cariPlat string
	var tabData arrData

	tabData[0] = Kendaraan{"D_1234_AB", "Toyota_Avanza", "2018", "12-Mei-2026"}
	tabData[1] = Kendaraan{"B_9999_XYZ", "Honda_Civic", "2020", "01-Januari-2026"}
	tabData[2] = Kendaraan{"Z_7777_AA", "Suzuki_Ertiga", "2019", "15-Maret-2025"}
	tabData[3] = Kendaraan{"F_5555_BB", "Mitsubishi_Pajero", "2021", "20-April-2026"}
	tabData[4] = Kendaraan{"DK_1111_CC", "Daihatsu_Xenia", "2022", "10-Februari-2026"}


	fmt.Printf("Berapa data yang akan di masukan: ")
	fmt.Scan(&targetInput)

	var i int
	for i = 0; i < targetInput; i++ {
		fmt.Printf("\n--- Input Data ke-%d ---\n", i+1)
		isiData(&jumlahSekarang, &tabData)
	}

	fmt.Printf("\nMasukan plat nomor yang dicari: ")
	fmt.Scan(&cariPlat)

	cariData(jumlahSekarang, cariPlat, tabData)
}