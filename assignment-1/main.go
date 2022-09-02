package main

import (
	"fmt"
	"os"
	"strconv"
)


type Student struct{
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}
func (s Student)Print(){
	fmt.Printf("Nama     : %s\n", s.Nama)
	fmt.Printf("Alamat   : %s\n", s.Alamat)
	fmt.Printf("Pekejaan : %s\n", s.Pekerjaan)
	fmt.Printf("Alasan   : %s\n", s.Alasan)
}

func main(){
	idx,_ := strconv.ParseInt(os.Args[1], 10, 64)

	students := []Student{ 
		{
			Nama: "Budi",
			Alamat:  "Jalan Keris",
			Pekerjaan: "Programmer",
			Alasan: "Karena suka programming",
		},
		{
			Nama: "Bano",
			Alamat:  "Jalan Kuda",
			Pekerjaan: "Karyawan",
			Alasan: "Karena suka tantangan",
		},
		{
			Nama: "Angga",
			Alamat:  "Jalan Bintar",
			Pekerjaan: "Sales",
			Alasan: "Memiliki ketertarikan kepada golang",
		},
		{
			Nama: "Dimas",
			Alamat:  "Jalan Pabo",
			Pekerjaan: "Gojek",
			Alasan: "Ingin belajar mengenai golang",
		},
		{
			Nama: "Daki",
			Alamat:  "Jalan Emas",
			Pekerjaan: "Admin",
			Alasan: "Penasaran terhadap dunia IT",
		},
	}

	students[idx].Print()
	

}