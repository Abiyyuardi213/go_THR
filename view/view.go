package view

import (
	"fmt"
	"go_THR/controller"
	"go_THR/model"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func VClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("Clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ShowMenu() {
	fmt.Println("=============================================")
	fmt.Println("=====        PT. BUKIT ASAM. TBK        =====")
	fmt.Println("=====     MINING INFORMATION SYSTEM     =====")
	fmt.Println("=============================================")
	fmt.Println(" ")

	fmt.Println("---------------------------------------------")
	fmt.Println("-- 1. Input Data                           --")
	fmt.Println("-- 2. Update Data                          --")
	fmt.Println("-- 3. Delete Data                          --")
	fmt.Println("-- 4. Search Data                          --")
	fmt.Println("-- 5. Read All Data                        --")
	fmt.Println("---------------------------------------------")
}

func VInputBarang() {
	var nama string
	var total int
	var harga float32
	var stake string

	fmt.Println("=====================================")
	fmt.Println("=====     INPUT DATA BARANG     =====")
	fmt.Println("=====================================")
	fmt.Println(" ")

	fmt.Println("-------------------------------------")
	fmt.Print("-- Nama    : ")
	fmt.Scan(&nama)

	fmt.Print("-- Jumlah  : ")
	fmt.Scan(&total)

	fmt.Print("-- Harga   : ")
	fmt.Scan(&harga)

	fmt.Print("-- Pemasok : ")
	fmt.Scan(&stake)

	cek := controller.CInputBarang(nama, total, harga, stake)
	VClearScreen()

	fmt.Println(" ")
	if cek == true {
		serialNumber := model.SystemLastId()

		fmt.Println("-- Input barang berhasil --")
		fmt.Println(" ")

		fmt.Println("=============================================")
		fmt.Println("=====----- DETAIL INFORMASI BARANG -----=====")
		fmt.Println("=============================================")
		fmt.Println(" ")

		fmt.Println("---------------------------------------------")
		fmt.Printf("-- Serial Number : %d\n", serialNumber)
		fmt.Printf("-- Nama Barang   : %s\n", nama)
		fmt.Printf("-- Jumlah Barang : %d\n", total)
		fmt.Printf("-- Harga Barang  : %.2f\n", harga)
		fmt.Printf("-- Supplier      : %s\n", stake)
		fmt.Printf("-- Create At     : %s\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("---------------------------------------------")
	} else {
		fmt.Println("-- Input barang gagal --")
	}
	fmt.Println(" ")
}

func VReadAll() {
	barangs := controller.CViewBarang()

	if barangs != nil {
		fmt.Println("=========================================")
		fmt.Println("=====-----     DATA BARANG     -----=====")
		fmt.Println("=========================================")
		fmt.Println(" ")

		fmt.Println("-----------------------------------------")
		for _, barang := range barangs {
			fmt.Println("-- Serial Number : ", barang.SerialNumber)
			fmt.Println("-- Nama Barang   : ", barang.Name)
			fmt.Println("-- Jumlah Barang : ", barang.Total)
			fmt.Println("-- Harga Barang  : ", barang.Price)
			fmt.Println("-- Supplier      : ", barang.Shiper)
			fmt.Println("-- Create At     : ", barang.Create_at)
			fmt.Println("-----------------------------------------")
		}
	} else {
		fmt.Println("-- Data Barang Masih Kosong --")
	}
}

func VUpdateBarang() {
	var serialNumber int

	fmt.Println("================================================")
	fmt.Println("=====-----     UPDATE DATA BARANG     -----=====")
	fmt.Println("================================================")
	fmt.Println(" ")

	fmt.Println("------------------------------------------------")
	fmt.Print("-- Input Serial Number : ")
	fmt.Scan(&serialNumber)

	barang := controller.CGetBarangBySerialNumber(serialNumber)
	VClearScreen()

	if barang == nil {
		fmt.Println("====================================================")
		fmt.Println("=====-----     BARANG TIDAK DITEMUKAN     -----=====")
		fmt.Println("====================================================")
		return
	}

	fmt.Println(" ")
	fmt.Println("================================================")
	fmt.Println("=====-----     DETAIL DATA BARANG     -----=====")
	fmt.Println("================================================")
	fmt.Println(" ")

	fmt.Println("------------------------------------------------")
	fmt.Println("-- Serial Number : ", barang.SerialNumber)
	fmt.Println("-- Nama Barang   : ", barang.Name)
	fmt.Println("-- Jumlah Barang : ", barang.Total)
	fmt.Println("-- Harga Barang  : ", barang.Price)
	fmt.Println("-- Supplier      : ", barang.Shiper)
	fmt.Println("-- Create At     : ", barang.Create_at)
	fmt.Println("------------------------------------------------")
	fmt.Println(" ")

	var name string
	var total int
	var price float32
	var shiper string

	fmt.Println("================================================")
	fmt.Println("=====-----     UPDATE DATA BARANG     -----=====")
	fmt.Println("================================================")
	fmt.Println(" ")

	fmt.Println("------------------------------------------------")
	fmt.Print("-- Nama Barang   : ")
	fmt.Scan(&name)

	fmt.Print("-- Jumlah Barang : ")
	fmt.Scan(&total)

	fmt.Print("-- Harga Baru    : ")
	fmt.Scan(&price)

	fmt.Print("-- Supplier      : ")
	fmt.Scan(&shiper)

	updatedBarang, success := controller.CUpdateBarang(serialNumber, name, total, price, shiper)
	VClearScreen()

	fmt.Println(" ")
	if success {
		fmt.Println("----- Barang Berhasil Di Perbarui -----")
		fmt.Println(" ")

		fmt.Println("=================================================")
		fmt.Println("=====-----     DATA BARANG TERBARU     -----=====")
		fmt.Println("=================================================")
		fmt.Println(" ")

		fmt.Println("-------------------------------------------------")
		fmt.Println("-- Serial Number : ", updatedBarang.SerialNumber)
		fmt.Println("-- Nama Barang   : ", updatedBarang.Name)
		fmt.Println("-- Jumlah Barang : ", updatedBarang.Total)
		fmt.Println("-- Harga Barang  : ", updatedBarang.Price)
		fmt.Println("-- Supplier      : ", updatedBarang.Shiper)
		fmt.Println("-- Create At     : ", updatedBarang.Create_at)
		fmt.Println("-------------------------------------------------")
	} else {
		fmt.Println("----- Gagal Memperbarui Barang -----")
	}
}

func VDeleteBarang() {
	var serialNumber int

	fmt.Println("================================================")
	fmt.Println("=====-----     DELETE DATA BARANG     -----=====")
	fmt.Println("================================================")
	fmt.Println(" ")

	fmt.Println("------------------------------------------------")
	fmt.Print("-- Input Serial Number : ")
	fmt.Scan(&serialNumber)

	barangToDelete, found := controller.CSearchBarang(serialNumber)
	VClearScreen()

	if found {
		fmt.Println("===========================================")
		fmt.Println("=====-----     DETAIL BARANG     -----=====")
		fmt.Println("===========================================")
		fmt.Println(" ")

		fmt.Println("-------------------------------------------")
		fmt.Println("-- Serial Number : ", barangToDelete.SerialNumber)
		fmt.Println("-- Nama Barang   : ", barangToDelete.Name)
		fmt.Println("-- Jumlah Barang : ", barangToDelete.Total)
		fmt.Println("-- Harga Barang  : ", barangToDelete.Price)
		fmt.Println("-- Supplier      : ", barangToDelete.Shiper)
		fmt.Println("-- Create At     : ", barangToDelete.Create_at)
		fmt.Println("-------------------------------------------")
		fmt.Println(" ")
		fmt.Println("-- Apakah anda yakin akan menghapus barang diatas? (y/n)")

		var confirm string
		fmt.Scan(&confirm)
		VClearScreen()

		if confirm == "y" || confirm == "Y" {
			deletedData, success := controller.CDeleteBarang(serialNumber)

			if success {
				fmt.Println(" ")
				fmt.Println("============================================")
				fmt.Println("=====     BARANG BERHASIL DI HAPUS     =====")
				fmt.Println("============================================")
				fmt.Println("-- Serial Number : ", deletedData.SerialNumber)
				fmt.Println("-- Nama Barang   : ", deletedData.Name)
				fmt.Println("-- Jumlah Barang : ", deletedData.Total)
				fmt.Println("-- Harga Barang  : ", deletedData.Price)
				fmt.Println("-- Supplier      : ", deletedData.Shiper)
				fmt.Println("-- Create At     : ", deletedData.Create_at)
				fmt.Println("--------------------------------------------")
			} else {
				fmt.Println("==========================================")
				fmt.Println("=====     GAGAL MENGHAPUS BARANG     =====")
				fmt.Println("==========================================")
			}
		} else {
			fmt.Println(" ")
			fmt.Println("-- Penghapusan Barang di Batalkan --")
		}
	} else {
		fmt.Println("============================================")
		fmt.Println("=====      BARANG TIDAK DITEMUKAN      =====")
		fmt.Println("============================================")
	}
}

func VSearchBarang() {
	var serialNumber int

	fmt.Println("==============================================")
	fmt.Println("=====-----     CARI DATA BARANG     -----=====")
	fmt.Println("==============================================")
	fmt.Println(" ")

	fmt.Println("----------------------------------------------")
	fmt.Print("-- Input Serial Number : ")
	fmt.Scan(&serialNumber)

	barang, found := controller.CReadBarang(serialNumber)
	VClearScreen()

	if found {
		fmt.Println("================================================")
		fmt.Println("=====-----     DETAIL DATA BARANG     -----=====")
		fmt.Println("================================================")
		fmt.Println(" ")

		fmt.Println("------------------------------------------------")
		fmt.Println("-- Serial Number : ", barang.SerialNumber)
		fmt.Println("-- Nama Barang   : ", barang.Name)
		fmt.Println("-- Jumlah Barang : ", barang.Total)
		fmt.Println("-- Harga Barang  : ", barang.Price)
		fmt.Println("-- Supplier      : ", barang.Shiper)
		fmt.Println("-- Create At     : ", barang.Create_at)
		fmt.Println("------------------------------------------------")
	} else {
		fmt.Println("====================================================")
		fmt.Println("=====-----     BARANG TIDAK DITEMUKAN     -----=====")
		fmt.Println("====================================================")
	}
}