package main

import "fmt"

func main(){
  var masukan int;
  fmt.Print("masukan angka anda : ")
  fmt.Scan(&masukan)

  switch {
  case masukan == 10:
    fmt.Println("wow goodjob")
  case (masukan < 9) && (masukan > 4):
     fmt.Println("not bad")
  default:
    {
      fmt.Println("are you fu*king idiot?")
    }
  }
}
