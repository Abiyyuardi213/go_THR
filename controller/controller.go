package controller

import (
	"fmt"
	"go_THR/model"
	"go_THR/node"
)

func CInputBarang(name string, stock int, price float32) bool {
	if name != "" && stock != 0 && price != 0 {
		model.InputBarang(name, stock, price)
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

func CFindBySerialNumber(serialNumber int) *node.DataBarang {
	barangLL := model.FindBySerialNumber(serialNumber)
	if barangLL != nil {
		return &barangLL.DBBarang
	}
	return nil
}

func CUpdateBarang(serialNumber int, name string, total int, price float32) (*node.DataBarang, bool) {
	fmt.Println(serialNumber)
	barang, success := model.MUpdateBarang(serialNumber, name, total, price)
	if success {
		return barang, true
	}
	return nil, false
}

func CDeleteBarang(serialNumber int) (*node.DataBarang, bool) {
	deleteData, success := model.MDeleteBarang(serialNumber)
	if success {
		return deleteData, true
	}
	return nil, false
}

func CSearchBarang(serialNumber int) (*node.DataBarang, bool) {
	barang, found := model.MSearchBarang(serialNumber)
	if found {
		return barang, true
	}
	return nil, false
}
