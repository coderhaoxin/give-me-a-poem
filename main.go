package main

import "github.com/atotto/clipboard"
import "github.com/spf13/cobra"
import "time"
import "fmt"

func main() {

	root := &cobra.Command{
		Use:   "give-me-a-poem",
		Short: "Get a poem.",
		Run: func(c *cobra.Command, args []string) {
			getPoem("all")
		},
	}

	root.AddCommand(&cobra.Command{
		Use:   "p",
		Short: "Get a pili poem",
		Run: func(c *cobra.Command, args []string) {
			getPoem("pili")
		},
	})

	root.AddCommand(&cobra.Command{
		Use:   "t",
		Short: "Get a tang poem",
		Run: func(c *cobra.Command, args []string) {
			getPoem("tang")
		},
	})

	root.AddCommand(&cobra.Command{
		Use:   "s",
		Short: "Get a song poem",
		Run: func(c *cobra.Command, args []string) {
			getPoem("song")
		},
	})

	root.Execute()
}

func getPoem(mode string) {
	var poems []string

	switch mode {
	case "all":
		poems = loadPili()
	case "pili":
		poems = loadPili()
	case "tang":
		poems = loadTang()
	case "song":
		poems = loadSong()
	}

	i := time.Now().Nanosecond() % len(poems)
	poem := poems[i]

	printAndCopy(poem)
}

func printAndCopy(poem string) {
	fmt.Printf("\n%s\n", poem)
	_ = clipboard.WriteAll(poem)
	fmt.Print("\nCopied!\n")
}
