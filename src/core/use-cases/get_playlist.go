package usecases

import (
	"io/fs"
	"path/filepath"
	"strings"

	clc "github.com/FranMT-S/Playlist-editor/src/core/collections"
	appconstants "github.com/FranMT-S/Playlist-editor/src/core/constants"
)

func GetPlaylist(path string) clc.HashSet[string] {
	allowed := appconstants.AllowedPlaylist()
	set := make(clc.HashSet[string])

	f := func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() {
			parts := strings.Split(info.Name(), ".")
			extension := parts[len(parts)-1]
			extension = strings.ToLower(extension)

			if allowed[extension] {
				set.Add(strings.ToLower(path))
			}
		}

		return err
	}

	filepath.WalkDir(path, f)

	return set
}

// func GetPlaylist(set clc.HashSet[string]) fs.WalkDirFunc {
// 	allowed := appconstants.AllowedPlaylist()

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
