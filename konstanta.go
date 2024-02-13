package main

import (
	"fmt"
)

func main (){
  //untuk nilai yang tidak bisa diubah
  const firstName string = "Yazid Istiqlal"

  const lastname string = "Adhy Fadhillah"

//multi constanta
const (
    square          = "kotak"
    isToday bool    = true
    numeric uint8   = 1
    floatNum        = 2.23
)

const (
     lele string = "ikan tawar"
     wader
     ikan bool = true
)

// gabungan konstanta
const hewan1, hewan2= "anjing", "kucing"


//const gabungan 

  fmt.Print("",firstName, lastname,"\n")
  fmt.Println("ini adalah", square,"dan ini adalah", isToday, "dan ini adalah nilai numeric",numeric, "\n")
  fmt.Printf("ini adalah desimal %.0f\n",floatNum)

  fmt.Println("lele dan wader adalah ikan",lele, ", apakah dia ikan air tawar?", ikan)
  fmt.Println("dasar kamu", hewan1, "dan", hewan2)
}
