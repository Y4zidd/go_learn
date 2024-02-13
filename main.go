package main

import "fmt"

/// interface / overring

type kendaraan interface {
  suarakendaraan()
}

type hewan interface {
  suarahewan()
}

// membuat tipe abstract

type mobil struct{
  kendaraan
}

type kucing struct {
  hewan
}

// isi struct 

func (mobil) suarakendaraan() {
  fmt.Println("Suara mobil :  Mberrrr ")
}

func (kucing) suarahewan() {
  fmt.Println("suara kucing : meongggg")
}

// Println

func main() {
  mobil{}.suarakendaraan()
  kucing{}.suarahewan()
}
