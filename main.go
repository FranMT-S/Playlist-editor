package main

import (
	"fmt"
	"slices"
	"time"

	hel "github.com/FranMT-S/Playlist-editor/src/core/helpers"
	uc "github.com/FranMT-S/Playlist-editor/src/core/use-cases"
)

func main() {
	var playlistList = []string{}
	// var playlists = uc.GetPlaylist("path")
	// var playlistList = make([]string, 0, len(playlists))
	var songInPc = uc.GetSongs("path")

	// for key := range songInPc {
	// 	playlistList = append(playlistList, key)
	// }

	var songs, err = uc.ReadPlaylists(playlistList)
	// var songs, err = uc.ReadPlaylists(playlistList)

	if err != nil {
		fmt.Println("Sucedio un error")
		return
	}

	fmt.Println("\n<-----Songs----->")
	for k := range songs {
		if _, ok := songInPc[k]; !ok {
			fmt.Printf("Canciones que no se encuentran en la PC %v\n", k)
			continue
		}

		songInPc[k] = true
	}
	fmt.Println("<-----Fin Songs----->")

	fmt.Println("\n<-----Faltantes en los albumes----->")
	var songMisses []string
	for k, v := range songInPc {
		if !v {
			// fmt.Printf("%v --> %v\n", k, v)
			songMisses = append(songMisses, k)
			// fmt.Printf("%v \n", k)
		}
	}

	slices.Sort(songMisses)
	now := time.Now()
	name := "songsMissess_" + now.Format("2006-01-02_15-04-05") + ".txt"
	hel.NewLog(name, songMisses)
	fmt.Println("<------------->")
	fmt.Printf("Lenght SOngs PC:%v", len(songInPc))

}
