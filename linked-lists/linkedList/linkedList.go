package singlyLinkedList

import "fmt"

type Song struct {
	Name     string
	Artist   string
	Previous *Song
	Next     *Song
}

type Playlist struct {
	Name       string
	head       *Song
	tail       *Song
	NowPlaying *Song
}

func CreatePlayList(name string) *Playlist {
	return &Playlist{
		Name: name,
	}
}

func (p *Playlist) Add(name, artist string) error {
	s := &Song{
		Name:   name,
		Artist: artist,
	}

	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.tail
		currentNode.Next = s
		s.Previous = p.tail
	}
	p.tail = s

	return nil
}

func (p *Playlist) ShowAll() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Playlist is empty")
		return nil
	}

	fmt.Printf("%+v\n", *currentNode)
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

func (p *Playlist) Play() *Song {
	p.NowPlaying = p.head
	return p.NowPlaying
}

func (p *Playlist) Next() *Song {
	p.NowPlaying = p.NowPlaying.Next
	return p.NowPlaying
}

func (p *Playlist) Previous() *Song {
	p.NowPlaying = p.NowPlaying.Previous
	return p.NowPlaying
}
