package model

import (
	"go_THR/database"
	"go_THR/node"
)

func InputBarang(name string, stock int, price float32) {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang

	serialNumber := SystemLastId()

	barang := node.DataBarang{
		SerialNumber: serialNumber,
		Name:         name,
		Stock:        stock,
		Price:        price,
	}

	LLBaru := node.BarangLL{
		DBBarang: barang,
	}

	if tmpLL.Next == nil {
		tmpLL.Next = &LLBaru
	} else {
		for tmpLL.Next != nil {
			tmpLL = tmpLL.Next
		}
		tmpLL.Next = &LLBaru
	}
}

func SystemLastId() int {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang

	if tmpLL.Next == nil {
		return 1
	} else {
		maxSerial := 0
		for tmpLL != nil {
			if tmpLL.DBBarang.SerialNumber > maxSerial {
				maxSerial = tmpLL.DBBarang.SerialNumber
			}
			tmpLL = tmpLL.Next
		}
		return maxSerial + 1
	}
}

func MReadAll() []node.DataBarang {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang
	var tabelBarang []node.DataBarang

	for tmpLL.Next != nil {
		tmpLL = tmpLL.Next
		tabelBarang = append(tabelBarang, tmpLL.DBBarang)
	}
	return tabelBarang
}

func FindBySerialNumber(serialNumber int) *node.BarangLL {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang

	for tmpLL.Next != nil {
		tmpLL = tmpLL.Next
		if tmpLL.DBBarang.SerialNumber == serialNumber {
			return tmpLL
		}
	}
	return nil
}

func MUpdateBarang(serialNumber int, name string, stock int, price float32) (*node.DataBarang, bool) {
	tmpLL := FindBySerialNumber(serialNumber)

	if tmpLL != nil {
		barang := &tmpLL.DBBarang
		barang.Name = name
		barang.Stock = stock
		barang.Price = price
		return barang, true
	}
	return nil, false
}

func MDeleteBarang(serialNumber int) (*node.DataBarang, bool) {
	tmpLL := FindBySerialNumber(serialNumber)
	if tmpLL != nil {
		deleteData := tmpLL.DBBarang
		prevLL := &database.DatabaseBarang

		for prevLL.Next != tmpLL {
			prevLL = prevLL.Next
		}
		prevLL.Next = tmpLL.Next
		return &deleteData, true
	}
	return nil, false
}

func MSearchBarang(serialNumber int) (*node.DataBarang, bool) {
	tmpLL := FindBySerialNumber(serialNumber)
	if tmpLL != nil {
		return &tmpLL.DBBarang, true
	}
	return nil, false
}
