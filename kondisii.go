package main 

import "fmt"

func main(){
  var masukan int;
  fmt.Print("masukan angka : ")
  fmt.Scan(&masukan)

  var masukan2 int;
  fmt.Print("masukan angka kedua : ")
  fmt.Scan(&masukan2)

  var banding int;
  fmt.Print("masukan anda ingin menggunakan operator apa 1. (!=), 2. (>), 3. (<) : ")
  fmt.Scan(&banding)

  if banding == 1 {
      fmt.Println(masukan,"!=", masukan2, "=" ,masukan != masukan2)
  } else if banding == 2 {
    fmt.Println(masukan,">", masukan2, "=", masukan > masukan2)
  } else if banding == 3 {
    fmt.Println(masukan,"<", masukan2, "=", masukan < masukan2)
  } else {
    fmt.Println("silahkan masukan format yang benar")
   }
}

