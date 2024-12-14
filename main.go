package main

import (
	"fmt"
	"path/filepath"

	clc "github.com/FranMT-S/Playlist-editor/src/core/collections"
	uc "github.com/FranMT-S/Playlist-editor/src/core/use-cases"
)

func main() {

	hashSet := make(clc.HashSet[string])
	filepath.WalkDir("D:/Red Sadness/Musica", uc.GetSongs(hashSet))
	fmt.Println(hashSet)
}
