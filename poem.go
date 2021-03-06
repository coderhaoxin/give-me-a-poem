package main

type Poem struct {
	Content string
	Tags    []string
}

func loadAll() []Poem {
	poems := []Poem{}

	poems = append(poems, chanPoems...)
	poems = append(poems, piliPoems...)
	poems = append(poems, tangPoems...)
	poems = append(poems, songPoems...)

	return poems
}

func loadChan() []Poem {
	return chanPoems
}

func loadPili() []Poem {
	return piliPoems
}

func loadTang() []Poem {
	return tangPoems
}

func loadSong() []Poem {
	return songPoems
}
