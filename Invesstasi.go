package main

import "fmt"

const NMAX int = 1000

type invest struct {
	nama           string
	danaInvestasi  float64
	jenisInvestasi string
	nilaiInvestasi float64
}

type tabInvest [NMAX]invest

//fungsi untuk menghitung persentase keuntungan/kerugian
func keuntungan(inv invest) float64 {
	if inv.danaInvestasi == 0 {
		return 0.0
	}
	return (inv.nilaiInvestasi - inv.danaInvestasi) / inv.danaInvestasi * 100
}

//mengurutkan berdasarkan danaInvestasi
func urutDanaNaik(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].danaInvestasi < A[idx].danaInvestasi {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func urutDanaTurun(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].danaInvestasi > A[idx].danaInvestasi {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

//mengurutkan berdasarkan nama
func urutNamaNaik(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].nama < A[idx].nama {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func urutNamaTurun(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].nama > A[idx].nama {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

//mengurutkan berdasarkan jenisInvestasi
func urutJenisNaik(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].jenisInvestasi < A[idx].jenisInvestasi {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func urutJenisTurun(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].jenisInvestasi > A[idx].jenisInvestasi {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

//mengurutkan berdasarkan nilaiInvestasi
func urutNilaiNaik(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].nilaiInvestasi < A[idx].nilaiInvestasi {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func urutNilaiTurun(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].nilaiInvestasi > A[idx].nilaiInvestasi {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

//mengurutkan berdasarkan persentase keuntungan
func urutPersentaseNaik(A *tabInvest, n int) {
	var i, j int
	var key invest
	var keyPersen float64
	for i = 1; i < n; i++ {
		key = A[i]
		keyPersen = keuntungan(key)
		j = i - 1
		for j >= 0 && keuntungan(A[j]) > keyPersen {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

func urutPersentaseTurun(A *tabInvest, n int) {
	var i, j int
	var key invest
	var keyPersen float64
	for i = 1; i < n; i++ {
		key = A[i]
		keyPersen = keuntungan(key)
		j = i - 1
		for j >= 0 && keuntungan(A[j]) < keyPersen {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
}

//pencarian secara sequential
func sequentialSearchByNama(A tabInvest, n int, cariNama string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if idx == -1 && A[i].nama == cariNama {
			idx = i
		}
	}
	return idx
}

func sequentialSearchByJenis(A tabInvest, n int, cariJenis string) int {
	var i, idx int
	idx = -1
	for i = 0; i < n; i++ {
		if idx == -1 && A[i].jenisInvestasi == cariJenis {
			idx = i
		}
	}
	return idx
}

//pencarian secara binary
func binarySearchNama(A tabInvest, n int, x string) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < A[mid].nama {
			right = mid - 1
		} else if x > A[mid].nama {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func binarySearchJenis(A tabInvest, n int, x string) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < A[mid].jenisInvestasi {
			right = mid - 1
		} else if x > A[mid].jenisInvestasi {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

//tampilkan data investasi
func tampilData(A tabInvest, n int) {
	var i int
	var persen float64
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-3s | %-15s | %-15s | %-15s | %-15s | %-15s |\n", "No", "Nama", "Jenis Investasi", "Dana Investasi", "Nilai Investasi", "Persen Keuntungan")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	for i = 0; i < n; i++ {
		persen = keuntungan(A[i])
		fmt.Printf("| %-3d | %-15s | %-15s | Rp.%-12.2f | Rp.%-12.2f | %-15.2f%% |\n", i+1, A[i].nama, A[i].jenisInvestasi, A[i].danaInvestasi, A[i].nilaiInvestasi, persen)
	}
	fmt.Println("-------------------------------------------------------------------------------------------------")
}

//laporan ringkasan portofolio investasi
func laporanRingkasan(A tabInvest, n int) {
	var i int
	var totalDana, totalNilai, totalProfit, avgPercent float64
	for i = 0; i < n; i++ {
		totalDana += A[i].danaInvestasi
		totalNilai += A[i].nilaiInvestasi
	}
	totalProfit = totalNilai - totalDana
	if totalDana != 0 {
		avgPercent = totalProfit / totalDana * 100
	}
	fmt.Println("==============================================")
	fmt.Println("      Laporan Portofolio Investasi")
	fmt.Println("----------------------------------------------")
	fmt.Printf("Total Dana Investasi : Rp.%.2f\n", totalDana)
	fmt.Printf("Total Nilai Investasi: Rp.%.2f\n", totalNilai)
	fmt.Printf("Total Profit/Loss    : Rp.%.2f\n", totalProfit)
	fmt.Printf("Rata-rata Persentase : %.2f%%\n", avgPercent)
	fmt.Println("==============================================")
}

//fungsi untuk menambahkan data investasi baru
func tambahData(A *tabInvest, n *int) {
	var inv invest
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&inv.nama)
	fmt.Print("Masukkan jenis investasi (saham/obligasi/reksa_dana) : ")
	fmt.Scan(&inv.jenisInvestasi)
	fmt.Print("Masukkan dana investasi : ")
	fmt.Scan(&inv.danaInvestasi)
	fmt.Print("Masukkan nilai investasi saat ini : ")
	fmt.Scan(&inv.nilaiInvestasi)
	A[*n] = inv
	*n = *n + 1
	fmt.Println("Data berhasil ditambahkan.")
}

//fungsi untuk mengubah data investasi
func ubahData(A *tabInvest, n int) {
	var cariNama, input string
	var inputB float64
	var i, idx int
	idx = -1
	fmt.Print("Masukkan nama yang akan diubah : ")
	fmt.Scan(&cariNama)
	for i = 0; i < n; i++ {
		if idx == -1 && A[i].nama == cariNama {
			idx = i
		}
	}
	if idx != -1 {
		fmt.Println("Data ditemukan. Ketik '-' jika tidak ingin mengubah suatu data.")
		fmt.Print("Nama baru : ")
		fmt.Scan(&input)
		if input != "-" {
			A[idx].nama = input
		}
		fmt.Print("Jenis investasi baru : ")
		fmt.Scan(&input)
		if input != "-" {
			A[idx].jenisInvestasi = input
		}
		fmt.Print("Dana investasi baru (-1 jika tidak diubah) : ")
		fmt.Scan(&inputB)
		if inputB != -1 {
			A[idx].danaInvestasi = inputB
		}
		fmt.Print("Nilai investasi baru (-1 jika tidak diubah) : ")
		fmt.Scan(&inputB)
		if inputB != -1 {
			A[idx].nilaiInvestasi = inputB
		}
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

//fungsi untuk menghapus data investasi
func hapusData(A *tabInvest, n *int) {
	var cariNama string
	var i, j, idx int
	idx = -1
	fmt.Print("Masukkan nama yang akan dihapus : ")
	fmt.Scan(&cariNama)
	for i = 0; i < *n; i++ {
		if idx == -1 && A[i].nama == cariNama {
			idx = i
		}
	}
	if idx != -1 {
		for j = idx; j < *n-1; j++ {
			A[j] = A[j+1]
		}
		*n = *n - 1
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

//fungsi pencarian data investasi
func cariData(A tabInvest, n int) {
	var pilihCari int
	var pilihMetode int
	var x string
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Jenis investasi")
	fmt.Print("Pilih (1/2) : ")
	fmt.Scan(&pilihCari)
	switch pilihCari {
	case 1:
		fmt.Print("Masukkan nama : ")
		fmt.Scan(&x)
		fmt.Println("Metode pencarian:")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search (data akan diurutkan terlebih dahulu)")
		fmt.Print("Pilih (1/2) : ")
		fmt.Scan(&pilihMetode)
		idx := -1
		if pilihMetode == 1 {
			idx = sequentialSearchByNama(A, n, x)
		} else {
			urutNamaNaik(&A, n)
			idx = binarySearchNama(A, n, x)
		}
		if idx != -1 {
			fmt.Println("Data ditemukan:")
			fmt.Printf("Nama      : %s\n", A[idx].nama)
			fmt.Printf("Jenis     : %s\n", A[idx].jenisInvestasi)
			fmt.Printf("Dana      : %.2f\n", A[idx].danaInvestasi)
			fmt.Printf("Nilai     : %.2f\n", A[idx].nilaiInvestasi)
			fmt.Printf("Keuntungan: %.2f%%\n", keuntungan(A[idx]))
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	case 2:
		fmt.Print("Masukkan jenis investasi : ")
		fmt.Scan(&x)
		fmt.Println("Metode pencarian:")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search (data akan diurutkan terlebih dahulu)")
		fmt.Print("Pilih (1/2) : ")
		fmt.Scan(&pilihMetode)
		idx := -1
		if pilihMetode == 1 {
			idx = sequentialSearchByJenis(A, n, x)
		} else {
			urutJenisNaik(&A, n)
			idx = binarySearchJenis(A, n, x)
		}
		if idx != -1 {
			fmt.Println("Data ditemukan:")
			fmt.Printf("Nama      : %s\n", A[idx].nama)
			fmt.Printf("Jenis     : %s\n", A[idx].jenisInvestasi)
			fmt.Printf("Dana      : %.2f\n", A[idx].danaInvestasi)
			fmt.Printf("Nilai     : %.2f\n", A[idx].nilaiInvestasi)
			fmt.Printf("Keuntungan: %.2f%%\n", keuntungan(A[idx]))
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func dummyData(A *tabInvest, n *int) {
	var i int
	dataDummy := []invest{
		{"faris", 10000, "saham", 12000},
		{"rian", 15000, "obligasi", 15500},
		{"damar", 20000, "reksa_dana", 18000},
		{"david", 5000, "saham", 7000},
		{"aan", 12000, "obligasi", 12500},
		{"lena", 25000, "reksa_dana", 26000},
		{"alex", 30000, "saham", 35000},
		{"jibran", 18000, "obligasi", 17500},
		{"vio", 22000, "reksa_dana", 23000},
		{"yogi", 40000, "saham", 42000},
	}
	for i = 0; i < len(dataDummy); i++ {
		A[i] = dataDummy[i]
	}
	*n = len(dataDummy)
}

//fungsi main untuk menampilkan menu
func main() {
	var data tabInvest
	var nData int = 0
	var pilihan int = -1
	var pilUrut int
	var urut int
	dummyData(&data, &nData)
	
	for pilihan != 0 {
		fmt.Println("\n==============================================")
		fmt.Println("      Aplikasi Manajemen Investasi Sederhana")
		fmt.Println("==============================================")
		fmt.Println("1. Tampilkan data investasi")
		fmt.Println("2. Tambah data investasi")
		fmt.Println("3. Ubah data investasi")
		fmt.Println("4. Hapus data investasi")
		fmt.Println("5. Cari data investasi")
		fmt.Println("6. Urutkan data investasi")
		fmt.Println("7. Laporan portofolio investasi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("\n---------------- Data Investasi ----------------")
			tampilData(data, nData)
		case 2:
			fmt.Println("\n--------- Tambah Data Investasi ---------")
			tambahData(&data, &nData)
		case 3:
			fmt.Println("\n--------- Ubah Data Investasi ---------")
			ubahData(&data, nData)
		case 4:
			fmt.Println("\n--------- Hapus Data Investasi ---------")
			hapusData(&data, &nData)
		case 5:
			fmt.Println("\n--------- Cari Data Investasi ---------")
			cariData(data, nData)
		case 6:
			fmt.Println("\n--------- Urutkan Data Investasi ---------")
			fmt.Println("Urutkan berdasarkan:")
			fmt.Println("1. Dana investasi")
			fmt.Println("2. Nama")
			fmt.Println("3. Jenis investasi")
			fmt.Println("4. Nilai investasi")
			fmt.Println("5. Persentase keuntungan")
			fmt.Print("Pilih (1-5): ")
			fmt.Scan(&pilUrut)
			switch pilUrut {
			case 1:
				fmt.Println("1. Ascending")
				fmt.Println("2. Descending")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&urut)
				switch urut {
				case 1:
					urutDanaNaik(&data, nData)
				case 2:
					urutDanaTurun(&data, nData)
				default:
					fmt.Println("Pilihan tidak valid.")
				}
				tampilData(data, nData)
			case 2:
				fmt.Println("1. Ascending")
				fmt.Println("2. Descending")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&urut)
				switch urut {
				case 1:
					urutNamaNaik(&data, nData)
				case 2:
					urutNamaTurun(&data, nData)
				default:
					fmt.Println("Pilihan tidak valid.")
				}
				tampilData(data, nData)
			case 3:
				fmt.Println("1. Ascending")
				fmt.Println("2. Descending")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&urut)
				switch urut {
				case 1:
					urutJenisNaik(&data, nData)
				case 2:
					urutJenisTurun(&data, nData)
				default:
					fmt.Println("Pilihan tidak valid.")
				}
				tampilData(data, nData)
			case 4:
				fmt.Println("1. Ascending")
				fmt.Println("2. Descending")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&urut)
				switch urut {
				case 1:
					urutNilaiNaik(&data, nData)
				case 2:
					urutNilaiTurun(&data, nData)
				default:
					fmt.Println("Pilihan tidak valid.")
				}
				tampilData(data, nData)
			case 5:
				fmt.Println("1. Ascending")
				fmt.Println("2. Descending")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&urut)
				switch urut {
				case 1:
					urutPersentaseNaik(&data, nData)
				case 2:
					urutPersentaseTurun(&data, nData)
				default:
					fmt.Println("Pilihan tidak valid.")
				}
				tampilData(data, nData)
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		case 7:
			fmt.Println("\n--------- Laporan Portofolio Investasi ---------")
			laporanRingkasan(data, nData)
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
