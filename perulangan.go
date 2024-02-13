package main

import "fmt"

func main() {
    var ang int;
    fmt.Print("Pilihlah angka yang ingin diualang : ")
    fmt.Scan(&ang)

    for i := 0; i < ang; i++ {
    fmt.Println("angka", i)
  }
}
