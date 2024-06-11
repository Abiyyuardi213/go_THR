package main

import (
	"fmt"
	"go_THR/handler"
	"go_THR/view"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func WebProgram() {
	http.HandleFunc("/", handler.ViewHandler)
	http.HandleFunc("/update/", handler.UpdateBarangHandlerWrapper)
	http.HandleFunc("/updatePost", handler.UpdateBarangPost)

	http.HandleFunc("/insert", handler.InsertBarangHandler)
	http.HandleFunc("/delete", handler.DeleteBarangHandler)

	// Menjalankan web server
	http.ListenAndServe(":8080", nil)
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("Clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func MainProgram() {
	var pilih int

	for {
		view.ShowMenu()
		fmt.Print("-- Input your choice : ")
		fmt.Scan(&pilih)
		ClearScreen()

		switch pilih {
		case 1:
			view.VInputBarang()

		case 2:
			view.VUpdateBarang()

		case 3:
			view.VDeleteBarang()

		case 4:
			view.VSearchBarang()

		case 5:
			view.VReadAll()

		default:
			fmt.Println("Pilihan anda tidak valid!")
		}
	}
}

func main() {
	// MainProgram()
	// model.InputBarang("batu", 100, 12000)
	WebProgram()
}
