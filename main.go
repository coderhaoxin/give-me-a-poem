package main

import "github.com/atotto/clipboard"
import "github.com/spf13/cobra"
import "time"
import "fmt"

func main() {

	version := "v0.2.0"

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

	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Get the version",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(version)
		},
	})

	root.Execute()
}

func getPoem(mode string) {
	var poems []Poem

	switch mode {
	case "all":
		poems = loadAll()
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

func printAndCopy(p Poem) {
	fmt.Printf("\n%s\n", p.Content)
	_ = clipboard.WriteAll(p.Content)
	fmt.Print("\nCopied!\n")
}
