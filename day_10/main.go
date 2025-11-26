package day10

import (
	"fmt"
)

type urlShortener struct {
	store map[string]string
}

func newURLShortener() *urlShortener {
	return &urlShortener{
		store: make(map[string]string),
	}
}

func (u *urlShortener) add(shortURL, originalURL string) {
	u.store[shortURL] = originalURL
}

func (u *urlShortener) get(shortURL string) (string, bool) {
	val, ok := u.store[shortURL]
	return val, ok
}

func (u *urlShortener) remove(shortURL string) {
	delete(u.store, shortURL)
}

func (u *urlShortener) display() {
	fmt.Println("\nAll URL mappings:")
	for short, original := range u.store {
		fmt.Printf("%s → %s\n", short, original)
	}
}

func URLShortenerDemo() {
	shortener := newURLShortener()

	shortener.add("abc123", "https://google.com")
	shortener.add("yt9", "https://youtube.com")
	shortener.add("gh12", "https://github.com")

	shortener.display()

	fmt.Println("\nRetrieving yt9:")
	if original, ok := shortener.get("yt9"); ok {
		fmt.Println("Original URL:", original)
	} else {
		fmt.Println("Short URL not found")
	}

	fmt.Println("\nRemoving abc123")
	shortener.remove("abc123")

	shortener.display()
}
