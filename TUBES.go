package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type invest struct {
	nama           string
	jenisInvestasi string
	danaInvestasi  float64
	nilaiInvestasi float64
}

type tabInvest [100]invest

// Pencarian secara binary
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

// Fungsi untuk mengecek nama
func checkNama(nama string) bool {
	if nama == "" {
		return false
	}
	for i := 0; i < len(nama); i++ {
		c := nama[i]
		// Hanya huruf (besar/kecil) dan spasi yang diperbolehkan
		if !(c == ' ' || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			return false
		}
	}
	return true
}

func dummyData(A *tabInvest, n *int) {
	dataDummy := []invest{
		{"Faris Surya", "saham", 10000, 12000},
		{"Rian Prasetro", "obligasi", 15000, 15500},
		{"Damar Widiawan", "reksa dana", 20000, 18000},
		{"David Rio", "saham", 5000, 7000},
		{"Farhan Andika", "obligasi", 12000, 12500},
		{"Lena Purtri", "reksa dana", 25000, 26000},
		{"Alex Ganendra", "saham", 30000, 35000},
		{"Jibran Yoseva", "obligasi", 18000, 17500},
		{"Vio Nazwa", "reksa dana", 22000, 23000},
		{"Yogi Yoseva", "saham", 40000, 42000},
	}
	for i, v := range dataDummy {
		A[i] = v
	}
	*n = len(dataDummy)
}

// Fungsi untuk membersihkan layar terminal
func clearScreen() {
	if os.Getenv("OS") == "Windows_NT" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// Perin
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Fungsi untuk menampilkan data investasi dalam bentuk tabel
func tampilData(A tabInvest, n int) {
	colNo := 5
	colNama := 17
	colJenis := 17
	colDana := 20
	colNilai := 20
	colPersen := 18

	top := "╔" + strings.Repeat("═", colNo) + "╦" + strings.Repeat("═", colNama) + "╦" +
		strings.Repeat("═", colJenis) + "╦" + strings.Repeat("═", colDana) + "╦" +
		strings.Repeat("═", colNilai) + "╦" + strings.Repeat("═", colPersen) + "╗"

	sep := "╠" + strings.Repeat("═", colNo) + "╬" + strings.Repeat("═", colNama) + "╬" +
		strings.Repeat("═", colJenis) + "╬" + strings.Repeat("═", colDana) + "╬" +
		strings.Repeat("═", colNilai) + "╬" + strings.Repeat("═", colPersen) + "╣"

	bottom := "╚" + strings.Repeat("═", colNo) + "╩" + strings.Repeat("═", colNama) + "╩" +
		strings.Repeat("═", colJenis) + "╩" + strings.Repeat("═", colDana) + "╩" +
		strings.Repeat("═", colNilai) + "╩" + strings.Repeat("═", colPersen) + "╝"

	header := fmt.Sprintf("║ %-3s ║ %-15s ║ %-15s ║ %-18s ║ %-18s ║ %-18s ║",
		"No", "Nama", "Jenis Investasi", "Dana Investasi", "Nilai Investasi", "Persen Keuntungan")

	fmt.Println(top)
	fmt.Println(header)
	fmt.Println(sep)
	for i := 0; i < n; i++ {
		persen := keuntungan(A[i])
		row := fmt.Sprintf("║ %-3d ║ %-15s ║ %-15s ║ Rp.%-15.2f ║ Rp.%-15.2f ║ %-17.2f%% ║",
			i+1, A[i].nama, A[i].jenisInvestasi, A[i].danaInvestasi, A[i].nilaiInvestasi, persen)
		fmt.Println(row)
	}
	fmt.Println(bottom)
}

// Fungsi untuk menghitung keuntungan dalam persen
func keuntungan(inv invest) float64 {
	if inv.danaInvestasi == 0 {
		return 0
	}
	return ((inv.nilaiInvestasi - inv.danaInvestasi) / inv.danaInvestasi) * 100
}

// Fungsi untuk menambahkan data investasi baru
func tambahData(A *tabInvest, n *int, reader *bufio.Reader) {
	var inv invest
	var valid bool

	valid = false
	for valid == false {
		fmt.Print("Masukkan nama (ketik '-' untuk batal): ")
		namaInput, _ := reader.ReadString('\n')
		inv.nama = strings.TrimSpace(namaInput)
		if inv.nama == "-" {
			fmt.Println("Operasi tambah data dibatalkan.")
			valid = true
			return
		} else if checkNama(inv.nama) {
			valid = true
		} else {
			fmt.Println("Nama tidak valid. Hanya huruf dan spasi yang diperbolehkan.")
		}
	}

	valid = false
	for valid == false {
		fmt.Print("Masukkan jenis investasi (saham/obligasi/reksa dana) (ketik '-' untuk batal): ")
		jenisInput, _ := reader.ReadString('\n')
		inv.jenisInvestasi = strings.TrimSpace(jenisInput)
		if inv.jenisInvestasi == "-" {
			fmt.Println("Operasi tambah data dibatalkan.")
			valid = true
			return
		} else if inv.jenisInvestasi == "saham" || inv.jenisInvestasi == "obligasi" || inv.jenisInvestasi == "reksa dana" {
			valid = true
		} else {
			fmt.Println("Jenis investasi tidak valid. Pilih antara saham, obligasi, atau reksa dana.")
		}
	}

	valid = false
	for valid == false {
		fmt.Print("Masukkan dana investasi (ketik -1 untuk batal): ")
		danaStr, _ := reader.ReadString('\n')
		danaStr = strings.TrimSpace(danaStr)
		var dana float64
		_, err := fmt.Sscanf(danaStr, "%f", &dana)
		if err != nil {
			fmt.Println("Input tidak valid. Pastikan memasukkan angka.")
		} else if dana == -1 {
			fmt.Println("Operasi tambah data dibatalkan.")
			valid = true
			return
		} else if dana > 0 {
			inv.danaInvestasi = dana
			valid = true
		} else {
			fmt.Println("Dana investasi harus lebih dari 0.")
		}
	}

	valid = false
	for valid == false {
		fmt.Print("Masukkan nilai investasi (ketik -1 untuk batal): ")
		nilaiStr, _ := reader.ReadString('\n')
		nilaiStr = strings.TrimSpace(nilaiStr)
		var nilai float64
		_, err := fmt.Sscanf(nilaiStr, "%f", &nilai)
		if err != nil {
			fmt.Println("Input tidak valid. Pastikan memasukkan angka.")
		} else if nilai == -1 {
			fmt.Println("Operasi tambah data dibatalkan.")
			valid = true
			return
		} else if nilai >= 0 {
			inv.nilaiInvestasi = nilai
			valid = true
		} else {
			fmt.Println("Nilai investasi tidak boleh negatif.")
		}
	}

	A[*n] = inv
	*n = *n + 1
	fmt.Println("Data berhasil ditambahkan.")
}

// Fungsi untuk mengubah data investasi
func ubahData(A *tabInvest, n int, reader *bufio.Reader) {
	var cariNama, input string
	var inputB float64
	var idx int
	idx = -1

	fmt.Print("Masukkan nama data yang akan diubah (ketik '-' untuk batal): ")
	namaCari, _ := reader.ReadString('\n')
	cariNama = strings.TrimSpace(namaCari)
	if cariNama == "-" {
		fmt.Println("Operasi ubah data dibatalkan.")
	} else {
		for i := 0; i < n; i++ {
			if idx == -1 && A[i].nama == cariNama {
				idx = i
			}
		}
		if idx == -1 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			fmt.Println("Data ditemukan. Jika ingin membatalkan, ketik '-' untuk string atau -1 untuk numerik.")
			fmt.Print("Nama baru: ")
			newName, _ := reader.ReadString('\n')
			input = strings.TrimSpace(newName)
			if input == "-" {
				fmt.Println("Operasi ubah data dibatalkan.")
			} else {
				A[idx].nama = input
				fmt.Print("Jenis investasi baru: ")
				newJenis, _ := reader.ReadString('\n')
				input = strings.TrimSpace(newJenis)
				if input == "-" {
					fmt.Println("Operasi ubah data dibatalkan.")
				} else {
					A[idx].jenisInvestasi = input
					fmt.Print("Dana investasi baru (ketik -1 untuk batal): ")
					danaStr, _ := reader.ReadString('\n')
					danaStr = strings.TrimSpace(danaStr)
					var err error
					_, err = fmt.Sscanf(danaStr, "%f", &inputB)
					if err != nil {
						fmt.Println("Input tidak valid. Operasi ubah data dibatalkan.")
					} else {
						if inputB == -1 {
							fmt.Println("Operasi ubah data dibatalkan.")
						} else {
							A[idx].danaInvestasi = inputB
							fmt.Print("Nilai investasi baru (ketik -1 untuk batal): ")
							nilaiStr, _ := reader.ReadString('\n')
							nilaiStr = strings.TrimSpace(nilaiStr)
							_, err = fmt.Sscanf(nilaiStr, "%f", &inputB)
							if err != nil {
								fmt.Println("Input tidak valid. Operasi ubah data dibatalkan.")
							} else {
								if inputB == -1 {
									fmt.Println("Operasi ubah data dibatalkan.")
								} else {
									A[idx].nilaiInvestasi = inputB
									fmt.Println("Data berhasil diubah.")
								}
							}
						}
					}
				}
			}
		}
	}
}

// Fungsi untuk menghapus data investasi
func hapusData(A *tabInvest, n *int, reader *bufio.Reader) {
	var cariNama string
	var i, j, idx int
	idx = -1

	fmt.Print("Masukkan nama data yang akan dihapus (ketik '-' untuk batal): ")
	namaInput, _ := reader.ReadString('\n')
	cariNama = strings.TrimSpace(namaInput)

	if cariNama == "-" {
		fmt.Println("Operasi hapus data dibatalkan.")
	} else {
		for i = 0; i < *n; i++ {
			if idx == -1 && A[i].nama == cariNama {
				idx = i
			}
		}

		if idx == -1 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			for j = idx; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n = *n - 1
			fmt.Println("Data berhasil dihapus.")
		}
	}
}

// Fungsi untuk mencari data investasi
func cariData(A tabInvest, n int, reader *bufio.Reader) {
	var pilihCari int
	var x string

	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama")
	fmt.Println("2. Jenis investasi")
	fmt.Print("Pilih (1/2): ")
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	_, err := fmt.Sscanf(inputStr, "%d", &pilihCari)

	if err != nil {
		fmt.Println("Input tidak valid!")
	} else {
		if pilihCari == 1 {
			fmt.Print("Masukkan nama (ketik '-' untuk batal): ")
			namaInput, _ := reader.ReadString('\n')
			x = strings.TrimSpace(namaInput)
			if x == "-" {
				fmt.Println("Operasi pencarian dibatalkan.")
			} else {
				idx := -1
				urutNamaNaik(&A, n)
				idx = binarySearchNama(A, n, x)
				if idx != -1 {
					fmt.Println("Data ditemukan:")
					fmt.Printf("Nama      : %s\n", A[idx].nama)
					fmt.Printf("Jenis     : %s\n", A[idx].jenisInvestasi)
					fmt.Printf("Dana      : Rp.%.2f\n", A[idx].danaInvestasi)
					fmt.Printf("Nilai     : Rp.%.2f\n", A[idx].nilaiInvestasi)
					fmt.Printf("Keuntungan: %.2f%%\n", keuntungan(A[idx]))
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			}
		} else if pilihCari == 2 {
			fmt.Print("Masukkan jenis investasi (ketik '-' untuk batal): ")
			jenisInput, _ := reader.ReadString('\n')
			x = strings.TrimSpace(jenisInput)
			if x == "-" {
				fmt.Println("Operasi pencarian dibatalkan.")
			} else {
				idx := -1
				urutJenisNaik(&A, n)
				idx = binarySearchJenis(A, n, x)
				if idx != -1 {
					fmt.Println("Data ditemukan:")
					fmt.Printf("Nama      : %s\n", A[idx].nama)
					fmt.Printf("Jenis     : %s\n", A[idx].jenisInvestasi)
					fmt.Printf("Dana      : Rp.%.2f\n", A[idx].danaInvestasi)
					fmt.Printf("Nilai     : Rp.%.2f\n", A[idx].nilaiInvestasi)
					fmt.Printf("Keuntungan: %.2f%%\n", keuntungan(A[idx]))
				} else {
					fmt.Println("Data tidak ditemukan.")
				}
			}
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi pengurutan berdasarkan danaInvestasi (ascending)
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

// Fungsi pengurutan berdasarkan danaInvestasi (descending)
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

// Fungsi pengurutan berdasarkan nama (ascending)
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

// Fungsi pengurutan berdasarkan nama (descending)
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

// Fungsi pengurutan berdasarkan jenis investasi (ascending)
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

// Fungsi pengurutan berdasarkan jenis investasi (descending)
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

// Fungsi pengurutan berdasarkan nilai investasi (ascending)
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

// Fungsi pengurutan berdasarkan nilai investasi (descending)
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

// Fungsi pengurutan berdasarkan persen keuntungan (ascending) dengan selection sort
func urutPersentaseNaik(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if keuntungan(A[i]) < keuntungan(A[idx]) {
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

// Fungsi pengurutan berdasarkan persen keuntungan (descending) dengan selection sort
func urutPersentaseTurun(A *tabInvest, n int) {
	var i, idx, pass int
	var temp invest
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if keuntungan(A[i]) > keuntungan(A[idx]) {
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

// Sub-menu untuk operasi Urutkan Data Investasi tanpa menggunakan continue/break
func menuUrutData(A *tabInvest, n int, reader *bufio.Reader) {
	menuActive := true
	for menuActive {
		clearScreen()
		fmt.Println("--------- Menu Urutkan Data Investasi ---------")
		tampilData(*A, n)
		fmt.Println("\nPilihan:")
		fmt.Println("1. Urutkan berdasarkan Dana Investasi")
		fmt.Println("2. Urutkan berdasarkan Nama")
		fmt.Println("3. Urutkan berdasarkan Jenis Investasi")
		fmt.Println("4. Urutkan berdasarkan Nilai Investasi")
		fmt.Println("5. Urutkan berdasarkan Persen Keuntungan")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var pilih int
		_, err := fmt.Sscanf(input, "%d", &pilih)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {
			if pilih == 0 {
				menuActive = false
			} else {
				// Variabel untuk pilihan urutan (1: ascending, 2: descending)
				var urut int
				fmt.Println("1. Ascending")
				fmt.Println("2. Descending")
				fmt.Print("Pilih: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				_, err := fmt.Sscanf(input, "%d", &urut)
				if err != nil {
					fmt.Println("Input tidak valid!")
				} else {
					if pilih == 1 { // Urutkan berdasarkan danaInvestasi
						if urut == 1 {
							urutDanaNaik(A, n)
						} else if urut == 2 {
							urutDanaTurun(A, n)
						} else {
							fmt.Println("Pilihan tidak valid!")
						}
					} else if pilih == 2 { // Berdasarkan nama
						if urut == 1 {
							urutNamaNaik(A, n)
						} else if urut == 2 {
							urutNamaTurun(A, n)
						} else {
							fmt.Println("Pilihan tidak valid!")
						}
					} else if pilih == 3 { // Berdasarkan jenisInvestasi
						if urut == 1 {
							urutJenisNaik(A, n)
						} else if urut == 2 {
							urutJenisTurun(A, n)
						} else {
							fmt.Println("Pilihan tidak valid!")
						}
					} else if pilih == 4 { // Berdasarkan nilaiInvestasi
						if urut == 1 {
							urutNilaiNaik(A, n)
						} else if urut == 2 {
							urutNilaiTurun(A, n)
						} else {
							fmt.Println("Pilihan tidak valid!")
						}
					} else if pilih == 5 { // Berdasarkan persen keuntungan
						if urut == 1 {
							urutPersentaseNaik(A, n)
						} else if urut == 2 {
							urutPersentaseTurun(A, n)
						} else {
							fmt.Println("Pilihan tidak valid!")
						}
					} else {
						fmt.Println("Pilihan tidak valid!")
					}
				}
			}
		}
		fmt.Println("\nData setelah pengurutan:")
		tampilData(*A, n)
		fmt.Println("Tekan Enter untuk melanjutkan...")
		_, _ = reader.ReadString('\n')
	}
}

// Fungsi laporanRingkasan untuk menampilkan laporan portofolio investasi
func laporanRingkasan(A tabInvest, n int) {
	var totalDana, totalNilai, totalProfit, avgPercent float64
	for i := 0; i < n; i++ {
		totalDana += A[i].danaInvestasi
		totalNilai += A[i].nilaiInvestasi
	}
	totalProfit = totalNilai - totalDana
	if totalDana != 0 {
		avgPercent = (totalProfit / totalDana) * 100
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

// Fungsi untuk menampilkan menu utama
func showMenu() {
	clearScreen()
	borderTop := "╔════════════════════════════════════════════════╗"
	borderMid := "╠════════════════════════════════════════════════╣"
	borderBot := "╚════════════════════════════════════════════════╝"

	fmt.Println(borderTop)
	fmt.Println("║          Aplikasi Manajemen Investasi          ║")
	fmt.Println(borderMid)
	fmt.Println("║ 1. Tampilkan data investasi                    ║")
	fmt.Println("║ 2. Tambah data investasi                       ║")
	fmt.Println("║ 3. Ubah data investasi                         ║")
	fmt.Println("║ 4. Hapus data investasi                        ║")
	fmt.Println("║ 5. Cari data investasi                         ║")
	fmt.Println("║ 6. Urutkan data investasi                      ║")
	fmt.Println("║ 7. Laporan portofolio investasi                ║")
	fmt.Println("║ 0. Keluar                                      ║")
	fmt.Println(borderBot)
	fmt.Print("Pilih menu: ")
}

// Menu untuk Tambah Data
func menuTambahData(A *tabInvest, n *int, reader *bufio.Reader) {
	menuActive := true
	for menuActive {
		clearScreen()
		fmt.Println("--------- Menu Tambah Data Investasi ---------")
		tampilData(*A, *n)
		fmt.Println("\nPilihan:")
		fmt.Println("1. Tambah Data Investasi")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var pilih int
		_, err := fmt.Sscanf(input, "%d", &pilih)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {
			if pilih == 0 {
				menuActive = false
			} else if pilih == 1 {
				tambahData(A, n, reader)
				fmt.Println("\nData setelah penambahan:")
				tampilData(*A, *n)
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		}
		fmt.Println("Tekan Enter untuk melanjutkan...")
		_, _ = reader.ReadString('\n')
	}
}

// Menu untuk Ubah Data
func menuUbahData(A *tabInvest, n int, reader *bufio.Reader) {
	menuActive := true
	for menuActive {
		clearScreen()
		fmt.Println("--------- Menu Ubah Data Investasi ---------")
		tampilData(*A, n)
		fmt.Println("\nPilihan:")
		fmt.Println("1. Ubah Data Investasi")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var pilih int
		_, err := fmt.Sscanf(input, "%d", &pilih)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {
			if pilih == 0 {
				menuActive = false
			} else if pilih == 1 {
				ubahData(A, n, reader)
				fmt.Println("\nData setelah perubahan:")
				tampilData(*A, n)
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		}
		fmt.Println("Tekan Enter untuk melanjutkan...")
		_, _ = reader.ReadString('\n')
	}
}

// Menu untuk Hapus Data
func menuHapusData(A *tabInvest, n *int, reader *bufio.Reader) {
	menuActive := true
	for menuActive {
		clearScreen()
		fmt.Println("--------- Menu Hapus Data Investasi ---------")
		tampilData(*A, *n)
		fmt.Println("\nPilihan:")
		fmt.Println("1. Hapus Data Investasi")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var pilih int
		_, err := fmt.Sscanf(input, "%d", &pilih)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {
			if pilih == 0 {
				menuActive = false
			} else if pilih == 1 {
				hapusData(A, n, reader)
				fmt.Println("\nData setelah penghapusan:")
				tampilData(*A, *n)
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		}
		fmt.Println("Tekan Enter untuk melanjutkan...")
		_, _ = reader.ReadString('\n')
	}
}

// Fungsi utama
func main() {
	var pilihan int
	var data tabInvest
	var nData int = 0
	reader := bufio.NewReader(os.Stdin)
	dummyData(&data, &nData)

	mainActive := true
	for mainActive {
		showMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		_, err := fmt.Sscanf(input, "%d", &pilihan)
		if err != nil {
			fmt.Println("Input tidak valid!")
		} else {
			if pilihan == 1 {
				fmt.Println("\n---------------- Data Investasi ----------------")
				tampilData(data, nData)
			} else if pilihan == 2 {
				menuTambahData(&data, &nData, reader)
			} else if pilihan == 3 {
				menuUbahData(&data, nData, reader)
			} else if pilihan == 4 {
				menuHapusData(&data, &nData, reader)
			} else if pilihan == 5 {
				fmt.Println("\n--------- Cari Data Investasi ---------")
				cariData(data, nData, reader)
			} else if pilihan == 6 {
				menuUrutData(&data, nData, reader)
			} else if pilihan == 7 {
				fmt.Println("\n--------- Laporan Portofolio Investasi ---------")
				laporanRingkasan(data, nData)
			} else if pilihan == 0 {
				fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
				mainActive = false
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		}
		fmt.Println("Tekan Enter untuk kembali ke menu...")
		_, _ = reader.ReadString('\n')
	}
}