package main

const a = "Hello world"

type ID int
type Currency float64
type Name string
type Total int

var (
	b bool     = true  // Padrão inferido de bool é sempre false
	c Total    = 10    // Padrão 0
	d Name     = "Gui" // ""
	e Currency = 1.2   // +0.0000000c+000
	f ID       = 1
)

func main() {
	a := "X"
	a = "fd"
	println(a, b, c, d, e)
}

func x() {

}
