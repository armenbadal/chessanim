package main

func main() {
	var s0 string = "a1"
	var b0 []byte = []byte(s0)
	println(b0)
	println(string(rune(b0[0])))
	println(string(rune(b0[1])))
}
