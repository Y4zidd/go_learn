package main

import "fmt"

// override 

type Model interface {
  Modelan()
}

type Nomor interface{
  Nomoran() int
}

// abstract

type Mobil struct {
  Model string
  Nomor int
}

type Motor struct {
  Model string
  Nomor int
}

type Keduanya struct {
  Mobil Mobil
  Motor Motor
}

// definisi dengan return

func (k Keduanya) Modelan() string{
  return k.Mobil.Model + " Dan " + k.Motor.Model
}

func (k Keduanya) Nomoran() int{
  return k.Mobil.Nomor + k.Motor.Nomor
}

// isi nilai nya
func main() {
  k:= Keduanya {
    Mobil: Mobil{Model: "Toyota", Nomor: 453},
    Motor: Motor{Model: "Honda", Nomor: 687},
  }

  fmt.Println(`model dan juga nomor nya 
  1. Model nya ada : `, k.Modelan(), "\n",
` 2. Nomor nya adalah : `, k.Nomoran())
}
