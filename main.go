package main

import "time"
import "fmt"

func main() {
	poems := loadPoems()
	i := time.Now().Nanosecond() % len(poems)
	fmt.Println(poems[i])
}
