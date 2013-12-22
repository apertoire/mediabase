package services

import (
	"apertoire.net/mediabase/bus"
	"apertoire.net/mediabase/helper"
	"apertoire.net/mediabase/message"
	"fmt"
	"github.com/goinggo/tracelog"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Scanner struct {
	Bus    *bus.Bus
	Config *helper.Config

	re           [3]*helper.Rexp
	includedMask string
}

func (self *Scanner) Start() {
	log.Printf("starting scanner service ...")

	log.Printf("compiling regular expressions ...")

	// test:="I am leaving from home in a while"
	// prepositionsRegex := make([]*regexp.Regexp, len(preps))
	// for i := 0; i < len(preps); i++ {
	// prepositionsRegex[i]=regexp.MustCompile(`\b`+preps[i]+`\b`)
	// }

	// for i := 0; i < len(prepositionsRegex); i++ {
	// fmt.Println(prepositionsRegex[i].String())
	// if loc := prepositionsRegex[i].FindStringIndex(test); loc != nil{
	// fmt.Println(test[loc[0]:loc[1]], "found at: ", loc[0])
	// break

	self.re[0] = &helper.Rexp{regexp.MustCompile(`(?i)/volumes/.*?/(?P<Resolution>.*?)/(?P<Name>.*?)\s\((?P<Year>\d\d\d\d)\)/(?:.*/)*bdmv/index.(?P<FileType>bdmv)$`)}
	self.re[1] = &helper.Rexp{regexp.MustCompile(`(?i)/volumes/.*?/(?P<Resolution>.*?)/(?P<Name>.*?)\s\((?P<Year>\d\d\d\d)\)/(?:.*/)*.*\.(?P<FileType>iso|img|nrg|mkv|avi|xvid|ts|mpg|dvr-ms|mdf|wmv)$`)}
	self.re[2] = &helper.Rexp{regexp.MustCompile(`(?i)/volumes/.*?/(?P<Resolution>.*?)/(?P<Name>.*?)\s\((?P<Year>\d\d\d\d)\)/(?:.*/)*(?:video_ts|hv000i01)\.(?P<FileType>ifo)$`)}

	self.includedMask = ".bdmv|.iso|.img|.nrg|.mkv|.avi|.xvid|.ts|.mpg|.dvr-ms|.mdf|.wmv|.ifo"

	go self.react()

	log.Printf("scanner service started")
}

func (self *Scanner) Stop() {
	// nothing right now
	log.Printf("scanner service stopped")
}

func (self *Scanner) react() {
	for {
		select {
		case msg := <-self.Bus.ScanMovies:
			go self.doScanMovies(msg.Reply)
		}
	}
}

func (self *Scanner) visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		log.Printf("from-start err: %s", err)
	}

	// log.Printf("maldito: %s", path)

	if !strings.Contains(self.includedMask, strings.ToLower(filepath.Ext(path))) {
		// log.Printf("[%s] excluding %s", filepath.Ext(path), path)
		return nil
	}

	for i := 0; i < 3; i++ {
		// match := self.re[i].FindStringSubmatch(strings.ToLower(path))
		// if match == nil {
		// 	continue
		// }
		var rmap = self.re[i].Match(path)
		if rmap == nil {
			continue
		}

		movie := &message.Movie{Title: rmap["Name"], File_Title: rmap["Name"], Year: rmap["Year"], Resolution: rmap["Resolution"], FileType: rmap["FileType"], Location: path}
		tracelog.TRACE("mb", "scanner", fmt.Sprintf("FOUND [%s] (%s))", movie.Title, movie.Location))

		self.Bus.MovieFound <- movie

		return nil
	}

	return nil
}

func (self *Scanner) doScanMovies(reply chan string) {
	log.Printf("inside ScanMovies")

	// reply <- "Movie scanning process started ..."

	err := filepath.Walk("/Volumes/hal-films", self.visit)
	if err != nil {
		log.Println("err: %s", err)
	}

	log.Printf("completed scanning hal for movies")

	err = filepath.Walk("/Volumes/wopr-films", self.visit)
	if err != nil {
		log.Println("err: %s", err)
	}

	log.Printf("completed scanning wopr for movies")
}
