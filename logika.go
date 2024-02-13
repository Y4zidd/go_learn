package main

import "fmt"

func main(){
  var kiri = false
  var kanan = true

  var kiridankanan = kiri && kanan
  fmt.Printf("kiri && kanan \t(%t) \n", kiridankanan)
}
