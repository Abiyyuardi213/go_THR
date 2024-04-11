package main

import (
	"fmt"
	"go_THR/view"
	"os"
	"os/exec"
	"runtime"
)

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
			fmt.Println("Masih loading yaaaa")

		case 4:
			fmt.Println("Masih loading yaaaa")

		case 5:
			view.VReadAll()

		default:
			fmt.Println("Pilihan anda tidak valid!")
		}
	}
}

func main() {
	fmt.Println("halo riva")
	fmt.Println(" ")

	MainProgram()
}