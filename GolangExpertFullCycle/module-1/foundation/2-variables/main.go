package main

const a = "Hello world"

var (
	b bool    = true  // Padrão inferido de bool é sempre false
	c int     = 10    // Padrão 0
	d string  = "Gui" // ""
	e float64 = 1.2   // +0.0000000c+000
)

func main() {
	a := "X"
	a = "fd"
	println(a, b, c, d, e)
}

func x() {

}
