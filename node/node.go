package node

type DataBarang struct {
	SerialNumber int
	Name         string
	Stock        int
	Price        float32
}

type BarangLL struct {
	DBBarang DataBarang
	Next     *BarangLL
}