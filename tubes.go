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

func isiData(n int, tabData *arrData) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&tabData[i].plat)
		fmt.Scan(&tabData[i].model)
		fmt.Scan(&tabData[i].tahun)
		fmt.Scan(&tabData[i].servis)
	}
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

func main() {
	var n int
	var cariPlat string
	var tabData arrData

	fmt.Scan(&n)

	isiData(n, &tabData)

	fmt.Scan(&cariPlat)

	cariData(n, cariPlat, tabData)
}