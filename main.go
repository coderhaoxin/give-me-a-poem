package main

import "github.com/atotto/clipboard"
import "time"
import "fmt"

func main() {
	poems := loadPoems()
	i := time.Now().Nanosecond() % len(poems)
	p := poems[i]

	fmt.Printf("\n%s\n", p)

	_ = clipboard.WriteAll(p)
	fmt.Print("\nCopied!\n")
}
