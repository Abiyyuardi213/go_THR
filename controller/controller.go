package controller

import (
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