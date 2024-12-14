package appconstants

var allowedSong = map[string]bool{
	"mp3":  true,
	"flac": true,
}

var allowedPlaylist = map[string]bool{
	"m3u8": true,
	"m3u":  true,
	"wpl":  true,
}

func AllowedSong() map[string]bool {
	return allowedSong
}

func AllowedPlaylist() map[string]bool {
	return allowedPlaylist
}
