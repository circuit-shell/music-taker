package service

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func extractAudio(src, dst string) {
	fmt.Printf("Extracting audio from %s to %s", src, dst)
	// Create a new command.
	cmd := exec.Command(
		"ffmpeg",
		"-i", src,
		"-vn",      // No video output
		"-ac", "2", // Set audio channels to 2
		"-ab", "192k", // Set audio bitrate
		"-ar", "48000", // Set audio sampling rate
		"-f", "mp3",
		dst)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("ffmpeg output:\n%s\n", out)
		fmt.Printf("Error: %v", err)
	}
	// save the output to a file
	fmt.Printf("Extracted audio to %s", dst)

	// Remove the video file.
	if err := os.Remove(src); err != nil {
		fmt.Printf("Error removing video file: %v", err)
	}
}

func enumeratePlaylistVideos(playlist *youtube.Playlist) {
	header := fmt.Sprintf("Playlist %s by %s", playlist.Title, playlist.Author)
	println(header)
	println(strings.Repeat("=", len(header)) + "\n")

	for k, v := range playlist.Videos {
		fmt.Printf("(%d) %s - '%s'\n", k+1, v.Author, v.Title)
	}
}

func downloadPlaylist(playlistID string) {
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		panic(err)
	}

	enumeratePlaylistVideos(playlist)

	for _, entry := range playlist.Videos {
		title := entry.Title
		video, err := client.VideoFromPlaylistEntry(entry)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)

		audioStream, _, err := client.GetStream(video, &video.Formats.WithAudioChannels()[0])
		if err != nil {
			panic(err)
		}

		file, err := os.Create(title + ".mp4")
		if err != nil {
			panic(err)
		}

		defer file.Close()
		_, err = io.Copy(file, audioStream)
		if err != nil {
			panic(err)
		}

		println("Downloaded video to " + title + ".mp4")
		extractAudio(title+".mp4", title+".mp3")
	}
}

func main() {
	playlistID := "PL7nDh4yU5JdeF1A-T7KtuKW3ptJzkI3fi"

	downloadPlaylist(playlistID)
}
