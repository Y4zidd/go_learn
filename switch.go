package main

import "fmt"

func main(){
  var operator string;
  fmt.Print(`masukan operator yang ingin digunakan
  1. tambah
  2. kurang
  3. bagi 
  4. kali
silahka pilih angak 1-4
silahkan masukan pilihan anda : `)
  fmt.Scan(&operator)

  var ang1, ang2 int
  fmt.Print("masukan angka pertama : ")
  fmt.Scan(&ang1)
   
  var tangkap1 int = ang1

  fmt.Print("masukan angka kedua : ")
  fmt.Scan(&ang2)

  var tangkap2 int = ang2

  switch operator{
  case "1":
    fmt.Println(tangkap1,"+", tangkap2, "=" ,ang1 + ang2)
  case "2":
    fmt.Println(tangkap1, "-", tangkap2, "=", ang1 - ang2)
  case "3":
    fmt.Println(tangkap1, ":", tangkap2, "=", ang1 / ang2)
  case "4":
    fmt.Println(tangkap1, "x", tangkap2, "=", ang1 * ang2)
  default:
    {
      fmt.Println("need input number :(")
    }
  }

}
