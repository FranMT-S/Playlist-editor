package usecases

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	clc "github.com/FranMT-S/Playlist-editor/src/core/collections"
)

func ReadPlaylists(filePaths []string) (clc.HashSet[string], error) {
	var wg sync.WaitGroup
	setSong := make(clc.HashSet[string])
	chSongs := make(chan string)

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("no se pudo abrir el archivo: %v", err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			processM3u8(file, chSongs)
		}()
	}

	go func() {
		wg.Wait()
		close(chSongs)
	}()

	for song := range chSongs {

		setSong.Add(strings.ToLower(song))
	}

	return setSong, nil
}

func processM3u8(file *os.File, chSongs chan string) {
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// fmt.Printf("\nLeyendo: %s", line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		chSongs <- line
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error leyendo el archivo: %v\n", err)
	}
}
