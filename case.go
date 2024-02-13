package main

import "fmt"

func main() {
  var kondisi int;
  fmt.Print("Masukan nomor pilihan anda : ")
  fmt.Scan(&kondisi)

// switch ganda
  switch kondisi {
  case 1,2,3,4,5 :
    fmt.Println("yah anda kurang beruntung")
  case 6,7,8,9,10:
    fmt.Println("selamat anda beruntung")
  default:
    {
      fmt.Println("goblok disuruh pilih 1 - 10, malah lebih dari 10 bloon")
    }
  }
}
