package model

import (
	"go_THR/database"
	"go_THR/node"
	"time"
)

func InputBarang(name string, total int, price float32, shiper string) {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang

	serialNumber := SystemLastId()

	barang := node.DataBarang{
		SerialNumber: serialNumber,
		Name: 		name,
		Total: 		total,
		Price: 		price,
		Shiper: 	shiper,
		Create_at: 	time.Now().Format("2006-01-02 15:04:05"),
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
		return 0;
	} else {
		for tmpLL.Next != nil {
			tmpLL = tmpLL.Next
		}
		return tmpLL.DBBarang.SerialNumber + 1
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

func MSearch(serialNumber int) *node.BarangLL {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang

	if tmpLL.Next != nil {
		for tmpLL.Next != nil {
			if tmpLL.Next.DBBarang.SerialNumber == serialNumber {
				return tmpLL.Next
			}
			tmpLL = tmpLL.Next
		}
	} else {
		return nil
	}
	return nil
}

func GetBarangBySerialNumber(serialNumber int) *node.DataBarang {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang

	for tmpLL != nil {
		if tmpLL.DBBarang.SerialNumber == serialNumber {
			return &tmpLL.DBBarang
		}
		tmpLL = tmpLL.Next
	}
	return nil
}

func UpdateBarang(serialNumber int, name string, total int, price float32, shiper string) (*node.DataBarang, bool) {
	barang := GetBarangBySerialNumber(serialNumber)

	if barang != nil {
		barang.Name = name
		barang.Total = total
		barang.Price = price
		barang.Shiper = shiper
		barang.Create_at = time.Now().Format("2006-01-02 15:04:05")
		return barang, true
	}
	return nil, false
}

func DeleteBarang(serialNumber int) (*node.DataBarang, bool) {
	var prevLL *node.BarangLL
	tmpLL := &database.DatabaseBarang

	for tmpLL != nil {
		if tmpLL.DBBarang.SerialNumber == serialNumber {
			deletedData := tmpLL.DBBarang
			if prevLL == nil {
				database.DatabaseBarang = *tmpLL.Next
			} else {
				prevLL.Next = tmpLL.Next
			}
			return &deletedData, true
		}
		prevLL = tmpLL
		tmpLL = tmpLL.Next
	}
	return nil, false
}

