package main

import "fmt"

func main() {
	var nama, NIM, kelas string
	fmt.Print("masukkan nama =")
	fmt.Scan(&nama)
	fmt.Print("masukkan NIM =")	
	fmt.Scan(&NIM)
	fmt.Print("masukkan kelas =")
	fmt.Scan(&kelas)

	fmt.Println("Perkenalkan saya adalah ",nama, "salah satu mahasiswa Prodi S1-IF dari kelas",kelas, "dengan NIM",NIM+".")


}
	   