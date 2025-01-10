package appconstants

const (
	M3U8 = "m3u8"
	M3U  = "m3u"
	WPL  = "wpl"
	MP3  = "mp3"
	FLAC = "flac"
	M4A  = "mp4"
)

var allowedSong = map[string]bool{
	MP3:  true,
	FLAC: true,
	M4A:  true,
}

var allowedPlaylist = map[string]bool{
	M3U8: true,
	M3U:  true,
	WPL:  true,
}

func AllowedSong() map[string]bool {
	return allowedSong
}

func AllowedPlaylist() map[string]bool {
	return allowedPlaylist
}
