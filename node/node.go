package node

type DataBarang struct {
	SerialNumber	int
	Name			string
	Total			int
	Price			float32
	Shiper			string
	Create_at		string
}

type BarangLL struct {
	DBBarang DataBarang
	next* BarangLL
}