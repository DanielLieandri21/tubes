package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// struct untuk menyimpan data PDA
type PDA struct {
	Ketua      string
	Anggota    []string
	Prodi      string
	Judul      string
	SumberDana string
	Luaran     string
	Tahun      int
}

// slice untuk menyimpan semua PDA
var pdaList []PDA

// membersihkan layar
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// menambahkan PDA baru
func tambahPDA() {
	var pda PDA

	// input data PDA
	fmt.Println("Masukkan data PDA:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ketua: ")
	pda.Ketua, _ = reader.ReadString('\n')
	pda.Ketua = strings.TrimSpace(pda.Ketua)

	fmt.Print("Anggota (pisahkan dengan koma): ")
	anggotaInput, _ := reader.ReadString('\n')
	pda.Anggota = strings.Split(strings.TrimSpace(anggotaInput), ",")

	fmt.Print("Prodi: ")
	pda.Prodi, _ = reader.ReadString('\n')
	pda.Prodi = strings.TrimSpace(pda.Prodi)

	fmt.Print("Judul: ")
	pda.Judul, _ = reader.ReadString('\n')
	pda.Judul = strings.TrimSpace(pda.Judul)

	fmt.Print("Sumber Dana (internal/eksternal): ")
	pda.SumberDana, _ = reader.ReadString('\n')
	pda.SumberDana = strings.TrimSpace(pda.SumberDana)

	fmt.Print("Luaran (publikasi/produk): ")
	pda.Luaran, _ = reader.ReadString('\n')
	pda.Luaran = strings.TrimSpace(pda.Luaran)

	fmt.Print("Tahun: ")
	fmt.Scanln(&pda.Tahun)

	// tambahkan PDA ke slice
	pdaList = append(pdaList, pda)
	fmt.Println("PDA berhasil ditambahkan!")
}

// mengubah PDA
func ubahPDA() {
	var indeks int
	fmt.Print("Masukkan indeks PDA yang akan diubah: ")
	fmt.Scanln(&indeks)

	// cek validitas indeks
	if indeks < 0 || indeks >= len(pdaList) {
		fmt.Println("Indeks tidak valid")
		return
	}

	var pda PDA

	// input data PDA yang baru
	fmt.Println("Masukkan data PDA yang baru:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ketua: ")
	pda.Ketua, _ = reader.ReadString('\n')
	pda.Ketua = strings.TrimSpace(pda.Ketua)

	fmt.Print("Anggota (pisahkan dengan koma): ")
	anggotaInput, _ := reader.ReadString('\n')
	pda.Anggota = strings.Split(strings.TrimSpace(anggotaInput), ",")

	fmt.Print("Prodi: ")
	pda.Prodi, _ = reader.ReadString('\n')
	pda.Prodi = strings.TrimSpace(pda.Prodi)

	fmt.Print("Judul: ")
	pda.Judul, _ = reader.ReadString('\n')
	pda.Judul = strings.TrimSpace(pda.Judul)

	fmt.Print("Sumber Dana (internal/eksternal): ")
	pda.SumberDana, _ = reader.ReadString('\n')
	pda.SumberDana = strings.TrimSpace(pda.SumberDana)

	fmt.Print("Luaran (publikasi/produk): ")
	pda.Luaran, _ = reader.ReadString('\n')
	pda.Luaran = strings.TrimSpace(pda.Luaran)

	fmt.Print("Tahun: ")
	fmt.Scanln(&pda.Tahun)

	// ubah data PDA pada indeks yang diberikan
	pdaList[indeks] = pda
	fmt.Println("PDA berhasil diubah!")
}

// menghapus PDA
func hapusPDA() {
	var index int
	fmt.Print("Masukkan indeks PDA yang akan dihapus: ")
	fmt.Scanln(&index)

	// cek indeks PDA
	if index < 0 || index >= len(pdaList) {
		fmt.Println("Indeks PDA tidak valid")
		return
	}

	// hapus PDA dari slice
	pdaList = append(pdaList[:index], pdaList[index+1:]...)
	fmt.Println("PDA berhasil dihapus!")
}

// menampilkan PDA berdasarkan tahun
func tampilPDAByTahun() {
	var tahun int
	fmt.Print("Masukkan tahun PDA: ")
	fmt.Scanln(&tahun)

	// duplikat slice pdaList
	pdaListCopy := make([]PDA, len(pdaList))
	copy(pdaListCopy, pdaList)

	// mengurutkan PDA berdasarkan tahun
	sort.Slice(pdaListCopy, func(i, j int) bool {
		return pdaListCopy[i].Tahun < pdaListCopy[j].Tahun
	})

	var kiri, kanan, tengah int
	var found bool
	kiri = 0
	kanan = len(pdaListCopy) - 1
	for kiri <= kanan && !found {
		tengah = (kiri + kanan) / 2
		if pdaListCopy[tengah].Tahun < tahun {
			kiri = tengah + 1
		} else if pdaListCopy[tengah].Tahun > tahun {
			kanan = tengah - 1
		} else {
			found = true
		}
	}
	fmt.Println("Ketua:", pdaListCopy[tengah].Ketua)
	fmt.Println("Anggota:", pdaListCopy[tengah].Anggota)
	fmt.Println("Prodi:", pdaListCopy[tengah].Prodi)
	fmt.Println("Judul:", pdaListCopy[tengah].Judul)
	fmt.Println("Sumber Dana:", pdaListCopy[tengah].SumberDana)
	fmt.Println("Luaran:", pdaListCopy[tengah].Luaran)
	fmt.Println("Tahun:", pdaListCopy[tengah].Tahun)
	fmt.Println()

	// tampilkan PDA berdasarkan tahun
	// for _, pda := range pdaList {
	// 	if pda.Tahun == tahun {
	// 		fmt.Println("Ketua:", pda.Ketua)
	// 		fmt.Println("Anggota:", pda.Anggota)
	// 		fmt.Println("Prodi:", pda.Prodi)
	// 		fmt.Println("Judul:", pda.Judul)
	// 		fmt.Println("Sumber Dana:", pda.SumberDana)
	// 		fmt.Println("Luaran:", pda.Luaran)
	// 		fmt.Println("Tahun:", pda.Tahun)
	// 		fmt.Println()
	// 	}
	// }
}

// menampilkan PDA berdasarkan fakultas/prodi
func tampilPDAByProdi() {
	var prodi string
	fmt.Print("Masukkan fakultas/prodi: ")
	fmt.Scanln(&prodi)

	// tampilkan PDA berdasarkan fakultas/prodi
	for _, pda := range pdaList {
		if strings.ToLower(pda.Prodi) == strings.ToLower(prodi) {
			fmt.Println("Ketua:", pda.Ketua)
			fmt.Println("Anggota:", pda.Anggota)
			fmt.Println("Prodi:", pda.Prodi)
			fmt.Println("Judul:", pda.Judul)
			fmt.Println("Sumber Dana:", pda.SumberDana)
			fmt.Println("Luaran:", pda.Luaran)
			fmt.Println("Tahun:", pda.Tahun)
			fmt.Println()
		}
	}
}

// mengurutkan PDA berdasarkan tahun
func urutkanPDAByTahun() {
	// duplikat slice pdaList
	pdaListCopy := make([]PDA, len(pdaList))
	copy(pdaListCopy, pdaList)

	// mengurutkan PDA berdasarkan tahun
	sort.Slice(pdaListCopy, func(i, j int) bool {
		return pdaListCopy[i].Tahun < pdaListCopy[j].Tahun
	})

	// tampilkan PDA yang sudah diurutkan
	for _, pda := range pdaListCopy {
		fmt.Println("Ketua:", pda.Ketua)
		fmt.Println("Anggota:", pda.Anggota)
		fmt.Println("Prodi:", pda.Prodi)
		fmt.Println("Judul:", pda.Judul)
		fmt.Println("Sumber Dana:", pda.SumberDana)
		fmt.Println("Luaran:", pda.Luaran)
		fmt.Println("Tahun:", pda.Tahun)
		fmt.Println()
	}
}

// menampilkan menu
func tampilkanMenu() {

	fmt.Println("=== APLIKASI TRI-DARMA PERGURUAN TINGGI ===")
	fmt.Println("1. Tambah PDA")
	fmt.Println("2. Ubah PDA")
	fmt.Println("3. Hapus PDA")
	fmt.Println("4. Tampilkan PDA Berdasarkan Tahun")
	fmt.Println("5. Tampilkan PDA Berdasarkan Fakultas/Prodi")
	fmt.Println("6. Urutkan PDA Berdasarkan Tahun")
	fmt.Println("7. Urutkan PDA Berdasarkan Jumlah PDA Pertahun")
	fmt.Println("0. Keluar")
	fmt.Println()
	fmt.Print("Pilih menu: ")
}

func urutkanPDA() {
	// menghitung jumlah PDA pertahun
	jumlahPDA := make(map[int]int)
	for _, pda := range pdaList {
		jumlahPDA[pda.Tahun]++
	}

	// membuat slice yang berisi tahun dan jumlah PDA
	type TahunJumlah struct {
		Tahun  int
		Jumlah int
	}
	var tahunJumlahList []TahunJumlah
	for tahun, jumlah := range jumlahPDA {
		tahunJumlahList = append(tahunJumlahList, TahunJumlah{Tahun: tahun, Jumlah: jumlah})
	}

	// mengurutkan slice berdasarkan jumlah PDA secara menurun
	sort.Slice(tahunJumlahList, func(i, j int) bool {
		return tahunJumlahList[i].Jumlah > tahunJumlahList[j].Jumlah
	})

	// menampilkan hasil pengurutan
	fmt.Println("PDA diurutkan berdasarkan jumlah PDA pertahun:")
	for _, tahunJumlah := range tahunJumlahList {
		fmt.Printf("Tahun %d: %d PDA\n", tahunJumlah.Tahun, tahunJumlah.Jumlah)
	}
}

func main() {

	// menjalankan aplikasi CLI
	for {
		tampilkanMenu()

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {

		case 1:
			clearScreen()
			tambahPDA()
		case 2:
			clearScreen()
			ubahPDA()
		case 3:
			clearScreen()
			hapusPDA()
		case 4:
			clearScreen()
			tampilPDAByTahun()
		case 5:
			clearScreen()
			tampilPDAByProdi()
		case 6:
			clearScreen()
			urutkanPDAByTahun()
		case 7:
			clearScreen()
			urutkanPDA()
		case 0:
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			os.Exit(0)
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid")
		}
	}
}
