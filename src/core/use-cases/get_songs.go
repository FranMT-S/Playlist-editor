package usecases

import (
	"io/fs"
	"strings"

	clc "github.com/FranMT-S/Playlist-editor/src/core/collections"
	appconstants "github.com/FranMT-S/Playlist-editor/src/core/constants"
)

func GetSongs(set clc.HashSet[string]) fs.WalkDirFunc {
	allowed := appconstants.AllowedSong()

	f := func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() {
			parts := strings.Split(info.Name(), ".")
			extension := parts[len(parts)-1]
			extension = strings.ToLower(extension)

			if allowed[extension] {
				set.Add(path)
			}
		}

		return err
	}

	return f
}
