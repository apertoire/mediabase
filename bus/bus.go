package bus

import (
	"apertoire.net/mediabase/message"
	"log"
)

type Bus struct {
	ScanMovies   chan *message.ScanMovies
	MovieFound   chan *message.Movie
	GetMovies    chan *message.GetMovies
	SearchMovies chan *message.SearchMovies
	StoreMovie   chan *message.Movie
	CachePicture chan *message.Picture
	UpdateMovie  chan *message.Picture
	Log          chan string
}

func (self *Bus) Start() {
	log.Println("bus starting up ...")

	self.ScanMovies = make(chan *message.ScanMovies)
	self.MovieFound = make(chan *message.Movie)

	self.GetMovies = make(chan *message.GetMovies)
	self.SearchMovies = make(chan *message.SearchMovies)

	self.StoreMovie = make(chan *message.Movie)
	self.CachePicture = make(chan *message.Picture)

	self.UpdateMovie = make(chan *message.Picture)

	self.Log = make(chan string)
}

// type Msg struct {
// 	id int
// }

// channel = make(chan *[]Msg)

// cannot use make(chan *[]Msg, 0) (type chan *[]Msg) as type chan *[]Msg in assignment
