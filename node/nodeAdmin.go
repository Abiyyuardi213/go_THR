package node

type DataAdmin struct {
	id 			int
	name 		string
	email 		string
	phone 		int
	username 	string
	password 	string
}

type AdminLL struct {
	DBAdmin DataAdmin
	Next* AdminLL
}