package view

import (
	"fmt"
	"go_THR/controller"
	"go_THR/model"
	"time"
)

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
			fmt.Println("-- Serial Number : ", barang.SerialNumber + 1)
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