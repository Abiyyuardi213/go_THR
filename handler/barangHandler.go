package handler

import (
	"fmt"

	"go_THR/controller"
	"go_THR/database"

	// "go_THR/model"
	"go_THR/node"
	"html/template"
	"net/http"
	"strconv"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	//memanggil memberView.html dengan tamplate

	tmpl := template.Must(template.ParseFiles(
		"view/BarangView.html"))
	users := controller.CViewBarang()
	// Menampilkan data ke template HTML
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InsertBarangHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		// Menampilkan form inputan
		tmpl := template.Must(template.ParseFiles(
			"view/InsertBarang.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle form submission
		r.ParseForm()
		name := r.Form.Get("name")
		stockStr := r.Form.Get("stock")
		priceStr := r.Form.Get("price")

		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			http.Error(w, "Invalid value for total", http.StatusBadRequest)
			return
		}

		price, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			http.Error(w, "Invalid value for price", http.StatusBadRequest)
			return
		}

		// Panggil controller untuk insert data
		controller.CInputBarang(name, stock, float32(price))

		// Menampilkan pesan sukses
		// w.Write([]byte("Data berhasil ditambahkan"))
		http.Redirect(w, r, "/view/BarangView.html", http.StatusSeeOther)
	}
}

func UpdateBarangHandler(serialNumber int, w http.ResponseWriter, r *http.Request) *node.BarangLL {
	var tmpLL *node.BarangLL
	tmpLL = &database.DatabaseBarang

	for tmpLL.Next != nil {
		tmpLL = tmpLL.Next
		if tmpLL.DBBarang.SerialNumber == serialNumber {
			// Execute the template with the data and write the result to the response writer
			tmpl := template.Must(template.ParseFiles("view/UpdateBarang.html"))
			if err := tmpl.Execute(w, tmpLL.DBBarang); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return nil
			}
			return tmpLL
		}
	}
	return nil
}
func UpdateBarangHandlerWrapper(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)

	// Extract serial number from URL path or request body
	// For example, you might extract it from URL path like this:
	serialNumberStr := r.URL.Path[len("/update/"):]
	// fmt.Print(serialNumberStr)
	serialNumber, err := strconv.Atoi(serialNumberStr)
	if err != nil {
		// Handle error
		http.Error(w, "Invalid serial number", http.StatusBadRequest)
		return
	}

	// Call UpdateBarangHandler with the extracted serial number
	result := UpdateBarangHandler(serialNumber, w, r)

	if result == nil {
		http.Error(w, "Barang not found", http.StatusNotFound)
		return
	}

	// Process the result if needed
	// For example, you can write the result to the response writer
	// json.NewEncoder(w).Encode(result)
}

func UpdateBarangPost(w http.ResponseWriter, r *http.Request) {
	// Handle form submission
	r.ParseForm()
	serialnumber := r.Form.Get("serialnumber")
	name := r.Form.Get("name")
	stockStr := r.Form.Get("stock")
	priceStr := r.Form.Get("price")

	serialNumberInt, err := strconv.Atoi(serialnumber)
	if err != nil {
		http.Error(w, "Invalid serial number", http.StatusBadRequest)
		return
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		http.Error(w, "Invalid value for total", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		http.Error(w, "Invalid value for price", http.StatusBadRequest)
		return
	}
	fmt.Println(serialnumber)
	// Panggil controller untuk insert data
	controller.CUpdateBarang(serialNumberInt, name, stock, float32(price))

	// Menampilkan pesan sukses
	// w.Write([]byte("Data berhasil ditambahkan"))
	http.Redirect(w, r, "/view/BarangView.html", http.StatusSeeOther)
}

func DeleteBarangHandler(w http.ResponseWriter, r *http.Request) {
	// Handle form submission
	fmt.Println("Delete")
	r.ParseForm()
	serialNumberStr := r.Form.Get("serial_number")
	fmt.Println(serialNumberStr)
	serialNumber, err := strconv.Atoi(serialNumberStr)
	if err != nil {
		http.Error(w, "Invalid value for serial number", http.StatusBadRequest)
		return
	}

	// Memanggil model untuk menghapus data barang
	controller.CDeleteBarang(serialNumber)
	// Menampilkan pesan sukses
	// w.Write([]byte("Data berhasil ditambahkan"))
	http.Redirect(w, r, "/view/BarangView.html", http.StatusSeeOther)
}
