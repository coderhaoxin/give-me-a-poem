package main

import "math/rand"
import "fmt"

func main() {
	poems := loadPoems()
	i := rand.Int() % len(poems)
	fmt.Println(poems[i])
}
