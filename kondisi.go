package main

import (
	"fmt"
	"go/printer"
)

func main(){
  var nilaianda int;

  fmt.Print("masukan nilai ujian anda : ")
  fmt.Scanf("%d", &nilaianda)

// if else  
  if nilaianda == 100 {
    fmt.Printf("A")
  } else if nilaianda >= 80 {
    fmt.Println("B")
  } else if nilaianda >= 60 {
    fmt.Println("C")
  } else if nilaianda >= 40 {
    fmt.Println("D")
  } else {
    fmt.Println("anda tidak lulus")
  }

// variable tempolery
   var point = 1000.0

   if percent := point / 100; percent >= 100 {
     fmt.Printf("%.1f%s perfect!\n", percent, "%")
   } else if percent >= 70 {
     fmt.Printf("%.1f%s good\n", percent, "%")
   } else {
     fmt.Printf("%.1f%s not bad\n", percent, "%")
   }
}


