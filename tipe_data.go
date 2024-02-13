package main

import "fmt"

func main () {
// non desimal
var positivenumber int64 = 1234
var negativnumber  = -23234

fmt.Printf ("Bilangan Positive : %d\n", positivenumber)
fmt.Println ("bilangan Negtif : ", negativnumber)//jika menggunakan println tidak usah di deklarasikan %d
  
//desimal number
var desimalnumber = 2.34

  fmt.Println("ini adalah hasil desimal number : ", desimalnumber)
  fmt.Printf("ini adalah hasil desimal number : %.4f\n", desimalnumber)

//boolean
exits := true

fmt.Printf("apakah anda ingin keluar : %t \n",exits)

//string
// tanda "``" digunakan jika string anda lebih dari 1 line
var nama string = `Yazid Istiqlal Adhy Fadhillah.
saya kenal dengan anda.
anda adalah manusia.`

fmt.Printf (nama)
}
