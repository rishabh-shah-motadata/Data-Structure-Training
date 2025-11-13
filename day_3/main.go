package day3

import (
	"fmt"
	"math/rand"
)

type node struct {
	song string
	next *node
}

type playlist struct {
	head *node
}

func (p *playlist) addAtBeginning(song string) {
	newNode := &node{song: song, next: p.head}
	p.head = newNode
}

func (p *playlist) addAtEnd(song string) {
	newNode := &node{song: song}
	if p.head == nil {
		p.head = newNode
		return
	}
	curr := p.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (p *playlist) deleteSong(song string) {
	if p.head == nil {
		fmt.Println("Playlist is empty!!")
		return
	}

	if p.head.song == song {
		p.head = p.head.next
		return
	}

	curr := p.head
	for curr.next != nil && curr.next.song != song {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func (p *playlist) display() {
	if p.head == nil {
		fmt.Println("Playlist is empty!!")
		return
	}

	curr := p.head
	fmt.Println("Playlist:")
	for curr != nil {
		fmt.Println("-", curr.song)
		curr = curr.next
	}
	fmt.Println()
}

func (p *playlist) merge(other *playlist) {
	if p.head == nil {
		p.head = other.head
		return
	}

	curr := p.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = other.head
	other.head = nil
}

func (p *playlist) shuffle() {
	if p.head == nil || p.head.next == nil {
		return
	}

	songs := []string{}
	curr := p.head
	for curr != nil {
		songs = append(songs, curr.song)
		curr = curr.next
	}

	rand.Shuffle(len(songs), func(i, j int) {
		songs[i], songs[j] = songs[j], songs[i]
	})

	p.head = nil
	for _, song := range songs {
		p.addAtEnd(song)
	}

	songs = []string{}
}

func MusicPlayer() {
	playlist1 := &playlist{}
	playlist1.addAtEnd("Song B")
	playlist1.addAtBeginning("Song A")
	playlist1.addAtEnd("Song B-1")
	playlist1.addAtEnd("Song C")

	playlist2 := &playlist{}
	playlist2.addAtEnd("Song Z")
	playlist2.addAtBeginning("Song Y")
	playlist2.addAtBeginning("Song X")

	fmt.Println("Original Playlist 1:")
	playlist1.display()

	playlist1.deleteSong("Song B-1")
	fmt.Println("After Deleting 'Song B', Playlist 1:")
	playlist1.display()

	fmt.Println("Original Playlist 2:")
	playlist2.display()

	playlist1.merge(playlist2)
	fmt.Println("Merged Playlist:")
	playlist1.display()

	playlist1.shuffle()
	fmt.Println("Shuffled Playlist:")
	playlist1.display()
}
