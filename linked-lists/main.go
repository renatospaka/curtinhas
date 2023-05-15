package main

import (
	"fmt"

	"github.com/renatospaka/linked-lists/linkedList"
)

func main() {
	playList := singlyLinkedList.CreatePlayList("MyPlaylist")
	fmt.Println("Created playlist")
	fmt.Println()

	fmt.Println("Adding songs to the playlist...")
	playList.Add("Ophelia", "The Lumineers")
	playList.Add("Shape of you", "Ed Sheeran")
	playList.Add("Stubborn Love", "The Lumineers")
	playList.Add("Feels", "Calvin Harris")

	fmt.Printf("Showing all songs in playlist...\n")
	playList.ShowAll()
	
	playList.Play()
	fmt.Printf("Now playing: %s by %s\n", playList.NowPlaying.Name, playList.NowPlaying.Artist)
	
	fmt.Printf("Changing next song...\n")
	playList.Next()
	fmt.Printf("Now playing: %s by %s\n", playList.NowPlaying.Name, playList.NowPlaying.Artist)
	
	fmt.Printf("Changing next song...\n")
	playList.Next()
	fmt.Printf("Now playing: %s by %s\n", playList.NowPlaying.Name, playList.NowPlaying.Artist)
}
