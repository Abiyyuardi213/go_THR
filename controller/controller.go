package controller

import (
	"go_THR/database"
	"go_THR/model"
	"go_THR/node"
)

func CInputBarang(name string, total int, price float32, shiper string) bool {
	if name != "" && total != 0 && price != 0 && shiper != "" {
		model.InputBarang(name, total, price, shiper)
		return true
	}
	return false
}

func CViewBarang() []node.DataBarang {
	barang := model.MReadAll()

	if barang == nil {
		return nil
	}
	return barang
}

func CGetBarangBySerialNumber(serialNumber int) *node.DataBarang {
	barang := model.GetBarangBySerialNumber(serialNumber)
	return barang
}

func CUpdateBarang(serialNumber int, name string, total int, price float32, shiper string) (*node.DataBarang, bool) {
	barang, success := model.UpdateBarang(serialNumber, name, total, price, shiper)
	return barang, success
}

func CDeleteBarang(serialNumber int) (*node.DataBarang, bool) {
	deletedData, success := model.DeleteBarang(serialNumber)
	return deletedData, success
}

func CSearchBarang(serialNumber int) (*node.DataBarang, bool) {
	tmpLL := database.DatabaseBarang

	for tmpLL.Next != nil {
		tmpLL = *tmpLL.Next
		if tmpLL.DBBarang.SerialNumber == serialNumber {
			return &tmpLL.DBBarang, true
		}
	}
	return nil, false
}

func CReadBarang(serialNumber int) (*node.DataBarang, bool) {
	result := model.MSearch(serialNumber)

	if result != nil {
		return &result.DBBarang, true
	}

	return nil, false
}