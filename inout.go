package main

import "fmt"

func main(){
  var ang1 int;
  var ang2 int;

  fmt.Print("Masukan angka Pertama mu : ")
  fmt.Scan(&ang1)
  fmt.Print("Masuka angka kedua mu : ")
  fmt.Scan(&ang2)

  var ang3 int = ang1 + ang2

  fmt.Println("Hasil penjumlahan mu adalah", ang3)
}
