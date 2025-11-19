package day8

import (
	"fmt"
	"time"
)

type book struct {
	id      string
	name    string
	author  string
	genre   string
	ratings int
}

type library struct {
	books         map[string]book
	indexByName   map[string][]string
	indexByAuthor map[string][]string
	indexByGenre  map[string][]string
	indexByRating map[int][]string
}

func newLibrary() *library {
	return &library{
		books:         make(map[string]book),
		indexByName:   make(map[string][]string),
		indexByAuthor: make(map[string][]string),
		indexByGenre:  make(map[string][]string),
		indexByRating: make(map[int][]string),
	}
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func appendIndex[T comparable](m map[T][]string, key T, id string) {
	m[key] = append(m[key], id)
}

func removeID(list []string, id string) []string {
	for i, v := range list {
		if v == id {
			list[i] = list[len(list)-1]
			return list[:len(list)-1]
		}
	}
	return list
}

func (l *library) insertBook(b book) {
	if b.id == "" {
		b.id = generateID()
	}

	l.books[b.id] = b
	appendIndex(l.indexByName, b.name, b.id)
	appendIndex(l.indexByAuthor, b.author, b.id)
	appendIndex(l.indexByGenre, b.genre, b.id)
	appendIndex(l.indexByRating, b.ratings, b.id)
}

func (l *library) removeBookByName(name string) {
	ids, ok := l.indexByName[name]
	if !ok {
		fmt.Println("No book found:", name)
		return
	}

	for _, id := range ids {
		b := l.books[id]

		l.indexByAuthor[b.author] = removeID(l.indexByAuthor[b.author], id)
		l.indexByGenre[b.genre] = removeID(l.indexByGenre[b.genre], id)
		l.indexByRating[b.ratings] = removeID(l.indexByRating[b.ratings], id)

		delete(l.books, id)
	}
	delete(l.indexByName, name)

	fmt.Printf("Removed all books named: %s\n\n", name)
}

func (l *library) search(name, author, genre string, rating int) {
	var ids []string

	switch {
	case name != "":
		ids = l.indexByName[name]
	case author != "":
		ids = l.indexByAuthor[author]
	case genre != "":
		ids = l.indexByGenre[genre]
	case rating != -1:
		ids = l.indexByRating[rating]
	default:
		for _, b := range l.books {
			fmt.Println(b)
		}
		return
	}

	for _, id := range ids {
		fmt.Println(l.books[id])
	}

	fmt.Println("")
}

func Library() {
	library := newLibrary()

	library.insertBook(book{name: "Dune", author: "Frank Herbert", genre: "SciFi", ratings: 5})
	library.insertBook(book{name: "Dune", author: "Unknown Edition", genre: "SciFi", ratings: 3})
	library.insertBook(book{name: "The Hobbit", author: "Tolkien", genre: "Fantasy", ratings: 5})
	library.insertBook(book{name: "It", author: "Stephen King", genre: "Horror", ratings: 4})
	library.insertBook(book{name: "Rich Dad Poor Dad", author: "Robert Kiyosaki", genre: "Finance", ratings: 4})
	library.insertBook(book{name: "Ikigai", author: "Robert Kiyosaki", genre: "Self Help", ratings: 4})

	fmt.Println("Search by genre = SciFi")
	library.search("", "", "SciFi", -1)

	fmt.Println("Search by name = Dune")
	library.search("Dune", "", "", -1)

	fmt.Println("Search by author = Robert Kiyosaki")
	library.search("", "Robert Kiyosaki", "", -1)

	library.removeBookByName("Dune")

	fmt.Println("Searching all books")
	library.search("", "", "", -1)
}
