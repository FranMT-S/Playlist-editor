package usecases

import (
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	appconstants "github.com/FranMT-S/Playlist-editor/src/core/constants"
	hel "github.com/FranMT-S/Playlist-editor/src/core/helpers"
)

func GetSongs(path string) map[string]bool {
	allowed := appconstants.AllowedSong()
	set := make(map[string]bool)
	var songs []string
	f := func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() {
			parts := strings.Split(info.Name(), ".")
			extension := parts[len(parts)-1]
			extension = strings.ToLower(extension)

			if allowed[extension] {
				path = strings.ToLower(path)
				songs = append(songs, path)
				set[path] = false
			}
		}

		return err
	}

	filepath.WalkDir(path, f)
	now := time.Now()
	name := "songs_" + now.Format("2006-01-02_15-04-05") + ".txt"
	hel.NewLog(name, songs)
	return set
}

// func GetSongs(set clc.HashSet[string]) fs.WalkDirFunc {
// 	allowed := appconstants.AllowedSong()

// 	f := func(path string, info fs.DirEntry, err error) error {
// 		if !info.IsDir() {
// 			parts := strings.Split(info.Name(), ".")
// 			extension := parts[len(parts)-1]
// 			extension = strings.ToLower(extension)

// 			if allowed[extension] {
// 				set.Add(path)
// 			}
// 		}

// 		return err
// 	}

// 	return f
// }
