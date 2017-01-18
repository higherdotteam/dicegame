package main

import "fmt"
import "os"
import "math"

func roll() int {
	f, _ := os.Open("/dev/urandom")
	b := make([]byte, 1)
	f.Read(b)
	f.Close()
	return int(math.Mod(float64(b[0]), 6.0) + 1)
}

func main() {
	fmt.Println("r", roll())
	fmt.Println("r", roll())
	fmt.Println("r", roll())
	fmt.Println("r", roll())
	fmt.Println("r", roll())
}
